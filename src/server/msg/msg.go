/*
 * @Author: Yuxiang Shan
 * @Mail: Yuxiang.Shan@shopee.com
 * @Date: 2022-09-12 19:32:06
 * @FilePath: /leafserver/src/server/msg/msg.go
 */
package msg

import (
	"github.com/name5566/leaf/network/protobuf"
	cs_msg "github.com/shanyux/leafserver/src/proto/cs_msg.pb.go"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = protobuf.NewProcessor()

func init() {
	// 这里我们注册了一个 JSON 消息 Hello
	Processor.Register(&cs_msg.Hello{})
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
// type Hello struct {
// 	Name string
// }
