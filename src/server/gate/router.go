/*
 * @Author: Yuxiang Shan
 * @Mail: Yuxiang.Shan@shopee.com
 * @Date: 2022-09-12 19:32:06
 * @FilePath: /leafserver/src/server/gate/router.go
 */
package gate

import (
	"github.com/shanyux/leafserver/src/proto/c2s"
	"github.com/shanyux/leafserver/src/server/game"
	"github.com/shanyux/leafserver/src/server/msg"
)

func init() {
	// 这里指定消息 Hello 路由到 game 模块
	// 模块间使用 ChanRPC 通讯，消息路由也不例外
	msg.Processor.SetRouter(&c2s.Hello{}, game.ChanRPC)
}
