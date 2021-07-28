package routers_test

import (
	"errors"
	"fmt"
	"github.com/siconghe/blog/pkg/util"
	"github.com/siconghe/blog/routers"
	"net/http"
	"net/http/httptest"
	"testing"
)




func TestTags(t *testing.T) {
	t.Run("testing:api/v1 GetTags", func(t *testing.T) {
		want := "{\"code\":200,\"data\":{\"lists\":null,\"total\":0},\"msg\":\"ok\"}"
		got,err := assertGotFromGetRouter("/api/v1/tags",t)
		util.AssertErrShouldBeNil(err,t)
		if want != got {
			t.Errorf("got:%v,want:%v",got,want)
		}
	})


}

func assertGotFromGetRouter(path string, t *testing.T) (got string,err error){
	util.MockDB()
	request := newGetRequestForTest(path,t)
	response := httptest.NewRecorder()
	r := routers.InitRouter()
	r.ServeHTTP(response,request)
	got = response.Body.String()
	if response.Code!= 200 {
		err = errors.New(fmt.Sprintf("响应状态码非200: %v",err))
	}
	return
}

func newGetRequestForTest(url string,t *testing.T) *http.Request {
	ret,err := http.NewRequest(http.MethodGet,url, nil)
	if err != nil {
		t.Errorf("获取http请求失败")
	}
	return ret
}


