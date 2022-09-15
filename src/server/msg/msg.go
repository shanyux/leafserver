/*
 * @Author: Yuxiang Shan
 * @Mail: Yuxiang.Shan@shopee.com
 * @Date: 2022-09-12 19:32:06
 * @FilePath: /leafserver/src/server/msg/msg.go
 */
package msg

import (
	"encoding/json"
	"io/ioutil"
	"reflect"

	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/network/protobuf"
	"github.com/shanyux/leafserver/src/proto/c2s"
)

// 使用默认的 JSON 消息处理器（默认还提供了 protobuf 消息处理器）
var Processor = protobuf.NewProcessor()

type MsgIdInfo struct {
	PbId    int64  `json:"pb_id"`
	MsgName string `json:"msg_name"`
}

func init() {
	msgids := make([]*MsgIdInfo, 0, 100)
	msgids = append(msgids, genMsgID(&c2s.Hello{}))

	b, err := json.Marshal(msgids)
	if err != nil {
		log.Error("error:%s", err.Error())
	}

	//生成json文件
	err = ioutil.WriteFile("leafserver.json", b, 0777)
	if err != nil {
		log.Error("error:%s", err.Error())
	}
}

func genMsgID(msg *c2s.Hello) *MsgIdInfo {
	id := Processor.Register(msg)

	// typeOfA := reflect.TypeOf(msg)
	return &MsgIdInfo{
		PbId:    int64(id),
		MsgName: getStructName(msg),
	}
}

func getStructName(myvar interface{}) string {
	if t := reflect.TypeOf(myvar); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

// 一个结构体定义了一个 JSON 消息的格式
// 消息名为 Hello
// type Hello struct {
// 	Name string
// }
