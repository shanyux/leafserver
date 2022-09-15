/*
 * @Author: Yuxiang Shan
 * @Mail: Yuxiang.Shan@shopee.com
 * @Date: 2022-09-12 19:52:36
 * @FilePath: /leafserver/src/client/main.go
 */
package main

import (
	"encoding/binary"
	"net"

	"github.com/golang/protobuf/proto"
	"github.com/shanyux/leafserver/src/proto/c2s"
	"github.com/shanyux/leafserver/src/server/msg"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:3563")
	if err != nil {
		panic(err)
	}

	// Hello 消息（JSON 格式）
	// 对应游戏服务器 Hello 消息结构体
	// data := []byte(`{
	// 	"Hello": {
	// 		"Name": "leaf"
	// 	}
	// }`)
	hello := &c2s.Hello{
		Name: proto.String("12dfsgfs3"),
	}
	m := packageMsg(hello)

	// log.Debug("hello %v", data)

	// // len + data
	// m := make([]byte, 2+len(data))

	// // 默认使用大端序
	// binary.BigEndian.PutUint16(m, uint16(len(data)))

	// copy(m[2:], data)

	// 发送消息
	conn.Write(m)
}

func packageMsg(message proto.Message) []byte {
	//
	data, _ := msg.Processor.Marshal(message)
	//...此处可先进行数据加密
	// len +id + data
	m := make([]byte, 4+len(data[1]))

	// 默认使用大端序
	binary.BigEndian.PutUint16(m, uint16(2+len(data[1])))

	//两个字节+数据
	copy(m[2:], data[0])
	copy(m[4:], data[1])

	return m

}
