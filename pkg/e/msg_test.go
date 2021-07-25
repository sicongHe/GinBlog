package e_test

import (
	"fmt"
	"github.com/siconghe/blog/pkg/e"
	"testing"
)

var(
	want = map[int]string {
		e.SUCCESS : "ok",
		e.ERROR : "fail",
		e.INVALID_PARAMS : "请求参数错误",
		e.ERROR_EXIST_TAG : "已存在该标签名称",
		e.ERROR_NOT_EXIST_TAG : "该标签不存在",
		e.ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
		e.ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
		e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
		e.ERROR_AUTH_TOKEN : "Token生成失败",
		e.ERROR_AUTH : "Token错误",
	}
)

func TestGetMsg(t *testing.T) {
	for key,value := range want{
		t.Run(fmt.Sprintf("错误信息测试：%s",value), func(t *testing.T) {
			checkMsg(key,value,t)
		})
	}
}

func checkMsg(key int,value string,t *testing.T){
	got,ok := e.MsgFlags[key]
	if ok {
		if got != value {
			t.Errorf("want %s, got %s",value,got)
		}
	} else {
		t.Errorf("错误码未定义")
	}


}
