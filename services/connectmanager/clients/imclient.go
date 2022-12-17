package clients

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-netty/go-netty"
	"github.com/jugglechat/im-server/commons/tools"
	"github.com/jugglechat/im-server/services/connectmanager/server/codec"
	"github.com/jugglechat/im-server/services/connectmanager/server/utils"
)

type ConnectState uint8

const (
	State_Disconnect ConnectState = 0
	State_Connecting ConnectState = 1
	State_Connected  ConnectState = 2
)

type ImClient struct {
	Address         string
	Token           string
	Appkey          string
	Platform        string
	DeviceId        string
	DeviceCompany   string
	DeviceModel     string
	DeviceOsVersion string
	PushToken       string

	ServerPubListener  func(topic, target string, time int64, data []byte)
	DisconnectCallback func(code ClientErrorCode, disMsg *codec.DisconnectMsgBody)

	UserId string

	channel         netty.Channel
	state           ConnectState
	accssorCache    sync.Map
	myIndex         uint16
	connAckAccessor *tools.DataAccessor
	pongAccessor    *tools.DataAccessor
}

func NewImClient(address, appkey, token string) *ImClient {
	return &ImClient{
		Address:         address,
		Appkey:          appkey,
		Token:           token,
		accssorCache:    sync.Map{},
		connAckAccessor: tools.NewDataAccessor(),
		pongAccessor:    tools.NewDataAccessorWithSize(100),
	}
}

func (client *ImClient) Connect(network, ispNum string, callback func(code ClientErrorCode, connAck *codec.ConnectAckMsgBody)) {
	if client.state == State_Disconnect {
		// setup child pipeline initializer.
		childInitializer := func(channel netty.Channel) {
		}
		// setup client pipeline initializer.
		clientInitializer := func(channel netty.Channel) {
			channel.Pipeline().
				AddLast(ImClientCodecHandler{}).
				AddLast(ImClientMessageHandler{Client: client})
		}
		// new bootstrap
		var bootstrap = netty.NewBootstrap(netty.WithChildInitializer(childInitializer), netty.WithClientInitializer(clientInitializer))
		var err error
		client.channel, err = bootstrap.Connect(client.Address, nil)
		if err == nil {
			connectMsg := codec.NewConnectMessage(&codec.ConnectMsgBody{
				ProtoId:         codec.ProtoId,
				SdkVersion:      "1.0.1",
				Appkey:          client.Appkey,
				Token:           client.Token,
				Platform:        client.Platform,
				DeviceId:        client.DeviceId,
				DeviceCompany:   client.DeviceCompany,
				DeviceModel:     client.DeviceModel,
				DeviceOsVersion: client.DeviceOsVersion,
				PushToken:       client.PushToken,

				NetworkId: network,
				IspNum:    ispNum,
			})
			client.channel.Write(connectMsg)
			connAckObj, err := client.connAckAccessor.GetWithTimeout(10 * time.Second)
			if err == nil {
				connAck := connAckObj.(*codec.ConnectAckMessage)
				clientCode := Trans2ClientErrorCoce(connAck.MsgBody.Code)
				if connAck.MsgBody.Code == utils.ConnectAckState_Access { //链接成功
					client.UserId = connAck.MsgBody.UserId
					client.state = State_Connected
					callback(clientCode, connAck.MsgBody)
				} else {
					callback(clientCode, nil)
				}
			} else { //超时
				callback(ClientErrorCode_ConnectTimeout, nil)
			}
		} else {
			callback(ClientErrorCode_SocketFailed, nil)
		}
	} else {
		callback(ClientErrorCode_ConnectExisted, nil)
	}
}

func (client *ImClient) Reconnect(network, ispNum string, callback func(code ClientErrorCode, connAck *codec.ConnectAckMsgBody)) {
	if client.state == State_Connected {
		if client.channel != nil {
			client.channel.Close(fmt.Errorf("reconnect"))
		}
		client.Connect(network, ispNum, callback)
	} else {
		callback(ClientErrorCode_ConnectExisted, nil)
	}
}
func (client *ImClient) Disconnect() {
	if client.channel != nil {
		disMsg := codec.NewDisconnectMessage(&codec.DisconnectMsgBody{
			Code: utils.DisconnectState_Quit,
		})
		client.channel.Write(disMsg)
		time.AfterFunc(5*time.Second, func() {
			client.channel.Close(fmt.Errorf("disconnect"))
		})
		client.channel = nil
	}
	client.state = State_Disconnect
}
func (client *ImClient) Logout() {
	if client.channel != nil {
		disMsg := codec.NewDisconnectMessage(&codec.DisconnectMsgBody{
			Code: utils.DisconnectState_Logout,
		})
		client.channel.Write(disMsg)
		client.channel.Close(fmt.Errorf("logout"))
		client.channel = nil
	}
	client.state = State_Disconnect
}

func (client *ImClient) OnConnectAck(msg *codec.ConnectAckMessage) {
	client.connAckAccessor.Put(msg)
}

func (client *ImClient) OnPublishAck(msg *codec.UserPublishAckMessage) {
	dataAccessor, ok := client.accssorCache.LoadAndDelete(msg.MsgBody.Index)
	if ok {
		dataAccessor.(*tools.DataAccessor).Put(msg)
	}
}
func (client *ImClient) OnQueryAck(msg *codec.QueryAckMessage) {
	dataAccessor, ok := client.accssorCache.LoadAndDelete(msg.MsgBody.Index)
	if ok {
		dataAccessor.(*tools.DataAccessor).Put(msg)
	}
}
func (client *ImClient) OnDisconnect(msg *codec.DisconnectMessage) {
	if client.DisconnectCallback != nil {
		client.DisconnectCallback(Trans2ClientErrorCoce(msg.MsgBody.Code), msg.MsgBody)
	}
}
func (client *ImClient) OnPong(msg *codec.PongMessage) {
	client.pongAccessor.Put(msg)
}
func (client *ImClient) OnPublish(msg *codec.ServerPublishMessage) {
	if client.ServerPubListener != nil {
		client.ServerPubListener(msg.MsgBody.Topic, msg.MsgBody.TargetId, msg.MsgBody.Timestamp, msg.MsgBody.Data)
		if msg.GetQoS() == codec.QoS_NeedAck {
			ackMsg := codec.NewServerPublishAckMessage(&codec.PublishAckMsgBody{
				Index: msg.MsgBody.Index,
			})
			client.channel.Write(ackMsg)
		}
	}
}

func (client *ImClient) OnInactive(ex netty.Exception) {
	if client.DisconnectCallback != nil {
		client.DisconnectCallback(ClientErrorCode_Unknown, nil)
	}
}
func (client *ImClient) OnException(ex netty.Exception) {
	if client.DisconnectCallback != nil {
		client.DisconnectCallback(ClientErrorCode_Unknown, nil)
	}
}

func (client *ImClient) Write(msg codec.IMessage) {
	client.channel.Write(msg)
}

func (client *ImClient) Publish(method, targetId string, data []byte, callback func(code ClientErrorCode, pubAck *codec.PublishAckMsgBody)) {
	if client.state == State_Connected {
		index := int32(client.getMyIndex())
		protoMsg := codec.NewUserPublishMessage(&codec.PublishMsgBody{
			Index:    index,
			Topic:    method,
			TargetId: targetId,
			Data:     data,
		})
		dataAccessor := tools.NewDataAccessor()
		client.accssorCache.Store(index, dataAccessor)
		client.Write(protoMsg)
		obj, err := dataAccessor.GetWithTimeout(10 * time.Second)
		if err == nil {
			pubAck := obj.(*codec.UserPublishAckMessage)
			if pubAck.MsgBody != nil {
				callback(Trans2ClientErrorCoce(pubAck.MsgBody.Code), pubAck.MsgBody)
			}
		} else { //超时
			callback(ClientErrorCode_SendTimeout, nil)
		}
	}
}

func (client *ImClient) Ping(callback func(code ClientErrorCode)) {
	if client.state == State_Connected {
		pingMsg := codec.NewPingMessage()
		client.Write(pingMsg)
		_, err := client.pongAccessor.GetWithTimeout(15 * time.Second)
		if err == nil {
			callback(ClientErrorCode_Success)
		} else {
			callback(ClientErrorCode_PingTimeout)
		}
	}
}

func (client *ImClient) Query(method, targetId string, data []byte, callback func(code ClientErrorCode, qryAck *codec.QueryAckMsgBody)) {
	if client.state == State_Connected {
		index := int32(client.getMyIndex())
		protoMsg := codec.NewQueryMessage(&codec.QueryMsgBody{
			Index:    index,
			Topic:    method,
			TargetId: targetId,
			Data:     data,
		})
		dataAccessor := tools.NewDataAccessor()
		client.accssorCache.Store(index, dataAccessor)
		client.Write(protoMsg)
		obj, err := dataAccessor.GetWithTimeout(10 * time.Second)
		if err == nil {
			queryAck := obj.(*codec.QueryAckMessage)
			if queryAck.MsgBody != nil {
				callback(Trans2ClientErrorCoce(queryAck.MsgBody.Code), queryAck.MsgBody)
			}
		} else { //timeout
			callback(ClientErrorCode_QueryTimeout, nil)
		}
	}
}

func (client *ImClient) getMyIndex() uint16 {
	client.myIndex = client.myIndex + 1
	return client.myIndex
}
