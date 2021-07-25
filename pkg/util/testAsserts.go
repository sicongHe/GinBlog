package util

import "testing"

func AssertErrShouldBeNil(err error,t *testing.T) {
	if err != nil {
		t.Errorf("Err应该为空！错误信息:%v",err)
	}
}
