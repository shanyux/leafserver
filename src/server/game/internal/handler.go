/*
 * @Author: Yuxiang Shan
 * @Mail: Yuxiang.Shan@shopee.com
 * @Date: 2022-09-12 19:32:06
 * @FilePath: /leafserver/src/server/game/internal/handler.go
 */
package internal

import (
	"reflect"

	"github.com/golang/protobuf/proto"
	"github.com/name5566/leaf/gate"
	"github.com/name5566/leaf/log"
	"github.com/shanyux/leafserver/src/proto/c2s"
)

func init() {
	// 向当前模块（game 模块）注册 Hello 消息的消息处理函数 handleHello
	handler(&c2s.Hello{}, handleHello)
}

func handler(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func handleHello(args []interface{}) {
	// 收到的 Hello 消息
	m := args[0].(*c2s.Hello)
	// 消息的发送者
	a := args[1].(gate.Agent)

	// 输出收到的消息的内容
	log.Debug("hello %s", m.GetName())

	// 给发送者回应一个 Hello 消息
	a.WriteMsg(&c2s.Hello{
		Name: proto.String("client"),
	})
}
