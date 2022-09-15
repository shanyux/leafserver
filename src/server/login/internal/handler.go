/*
 * @Author: Yuxiang Shan
 * @Mail: Yuxiang.Shan@shopee.com
 * @Date: 2022-09-12 19:32:06
 * @FilePath: /leafserver/src/server/login/internal/handler.go
 */
package internal

import (
	"reflect"
)

func handleMsg(m interface{}, h interface{}) {
	skeleton.RegisterChanRPC(reflect.TypeOf(m), h)
}

func init() {

}
