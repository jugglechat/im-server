package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/jugglechat/im-server/commons/gmicro"
	"github.com/jugglechat/im-server/commons/gmicro/actorsystem"
	"github.com/jugglechat/im-server/commons/gmicro/utils"
	timewheel "github.com/rfyiamcool/go-timewheel"
	"google.golang.org/protobuf/proto"
)

type Inter interface {
	StringX() string
}

type Father struct {
	Fat string
}

type Stu struct {
	Name string
	Age  int
	Father
}

type MyActor struct {
	actorsystem.UntypedActor
}

func (act *MyActor) OnReceive(input proto.Message) {
	fmt.Println("process has been executed.")
	fmt.Println("type:", reflect.TypeOf(input))
	stu := input.(*utils.Student)
	fmt.Println(stu.Name)
	time.Sleep(3 * time.Second)
}
func (act *MyActor) CreateInputObj() proto.Message {
	return &utils.Student{}
}

func main() {
	main_a()
}

func main_a() {
	actorSystem := actorsystem.NewActorSystemNoRpc("MyActorSystem")

	actorSystem.RegisterActor("m1", func() actorsystem.IUntypedActor { return &MyActor{} }, 3)

	for i := 0; i < 10; i++ {
		actor := actorSystem.LocalActorOf("m1")
		actor.Tell(&utils.Student{
			Name: "name2",
			Age:  1,
		}, actorsystem.NoSender)
	}

	time.Sleep(500 * time.Second)
}

func TimewheelTest() {
	tw, err := timewheel.NewTimeWheel(1*time.Second, 360)
	if err != nil {
		panic(err)
	}

	tw.Start()

	tw.Add(5*time.Second, func() {
		fmt.Println("aabbcc")
	})
	tw.Add(5*time.Second, func() {
		fmt.Println("ddeeff")
	})
	task := tw.Add(5*time.Second, func() {
		fmt.Println("gghhii")
	})

	time.Sleep(1 * time.Second)
	tw.Remove(task)

	time.Sleep(10 * time.Second)
}

func ActorSystemTest() {
	actorSystem := actorsystem.NewActorSystemNoRpc("MyActorSystem")
	// actorSystem.RegisterActorProcessor("m1", func() proto.Message {
	// 	return &utils.Student{}
	// }, MyProcessor, 10)

	actor := actorSystem.LocalActorOf("m1")
	actor.Tell(&utils.Student{
		Name: "name2",
		Age:  1,
	}, actorsystem.NoSender)
	time.Sleep(5 * time.Second)
}

func MyProcessor(sender actorsystem.ActorRef, input proto.Message) {
	fmt.Println("process has been executed.")
	fmt.Println("type:", reflect.TypeOf(input))
	stu := input.(*utils.Student)
	fmt.Println(stu.Name)

	sender.Tell(stu, actorsystem.NoSender)
}

func ClusterTest() {
	cluster := gmicro.NewCluster("cluster1", "node1", "127.0.0.1", 8888, []string{"127.0.0.1:2181"})
	// cluster.RegisterActorProcessor("m1", func() proto.Message {
	// 	return &utils.Student{}
	// }, MyProcessor, 10)

	stu := &utils.Student{
		Name: "name1",
		Age:  1,
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Begin to Tell")
	cluster.UnicastRouteWithNoSender("m1", stu.Name, stu)

	time.Sleep(10 * time.Second)
}
