package tools

import (
	"bytes"
	"encoding/binary"
	"hash/crc32"
	"sync/atomic"
)

const (
	base32EncodeChars string = "abcdefghgklmnpqrstuvwxyz23456789"
)

var currentSeq uint32 = 0

func GenerateMsgId(time int64, channelType byte, targetId string) string {
	seq := getSeq()
	time = time << 12
	time = time | int64(seq)

	time = time << 4
	time = time | (int64(channelType) & 0xf)

	targetHashCode := crc32.ChecksumIEEE([]byte(targetId))
	targetHashCode = targetHashCode & 0x3fffff
	time = time << 6
	time = time | (int64(targetHashCode >> 16))
	low := targetHashCode << 16

	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, time)
	binary.Write(buf, binary.BigEndian, low)

	bs := buf.Bytes()
	b1 := bs[0]
	b2 := bs[1]
	b3 := bs[2]
	b4 := bs[3]
	b5 := bs[4]
	b6 := bs[5]
	b7 := bs[6]
	b8 := bs[7]
	b9 := bs[8]
	b10 := bs[9]

	retBs := []byte{}

	retBs = append(retBs, base32EncodeChars[b1>>3])
	retBs = append(retBs, base32EncodeChars[((b1&0x7)<<2)|(b2>>6)])
	retBs = append(retBs, base32EncodeChars[(b2&0x3e)>>1])
	retBs = append(retBs, base32EncodeChars[((b2&0x1)<<4)|(b3>>4)])
	//retBs = append(retBs, '-')
	retBs = append(retBs, base32EncodeChars[((b3&0xf)<<1)|(b4>>7)])
	retBs = append(retBs, base32EncodeChars[(b4&0x7c)>>2])
	retBs = append(retBs, base32EncodeChars[((b4&0x3)<<3)|(b5>>5)])
	retBs = append(retBs, base32EncodeChars[b5&0x1f])
	//retBs = append(retBs, '-')
	retBs = append(retBs, base32EncodeChars[b6>>3])
	retBs = append(retBs, base32EncodeChars[((b6&0x7)<<2)|(b7>>6)])
	retBs = append(retBs, base32EncodeChars[(b7&0x3e)>>1])
	retBs = append(retBs, base32EncodeChars[((b7&0x1)<<4)|(b8>>4)])
	//retBs = append(retBs, '-')
	retBs = append(retBs, base32EncodeChars[((b8&0xf)<<1)|(b9>>7)])
	retBs = append(retBs, base32EncodeChars[(b9&0x7c)>>2])
	retBs = append(retBs, base32EncodeChars[((b9&0x3)<<3)|(b10>>5)])
	retBs = append(retBs, base32EncodeChars[b10&0x1f])

	return string(retBs)
}

func getSeq() uint32 {
	seq := atomic.AddUint32(&currentSeq, 1)
	seq = seq & 0xfff
	return seq
}
