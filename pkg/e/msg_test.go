package e

import (
	"fmt"
	"testing"
)

var(
	want = map[int]string {
		SUCCESS : "ok",
		ERROR : "fail",
		INVALID_PARAMS : "请求参数错误",
		ERROR_EXIST_TAG : "已存在该标签名称",
		ERROR_NOT_EXIST_TAG : "该标签不存在",
		ERROR_NOT_EXIST_ARTICLE : "该文章不存在",
		ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
		ERROR_AUTH_CHECK_TOKEN_TIMEOUT : "Token已超时",
		ERROR_AUTH_TOKEN : "Token生成失败",
		ERROR_AUTH : "Token错误",
	}
)

func TestMsg(t *testing.T) {
	for key,value := range want{
		t.Run(fmt.Sprintf("错误信息测试：%s",value), func(t *testing.T) {
			checkMsg(key,value,t)
		})
	}
}

func checkMsg(key int,value string,t *testing.T){
	got,ok := MsgFlags[key]
	if ok {
		if got != value {
			t.Errorf("want %s, got %s",value,got)
		}
	} else {
		t.Errorf("错误码未定义")
	}


}
