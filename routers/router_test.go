package routers_test

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/siconghe/blog/pkg/util"
	"github.com/siconghe/blog/routers"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestInitRouter(t *testing.T) {
	t.Run("router test => GET: /tom", func(t *testing.T) {
		request := newGetRequestForTest("/tom",t)
		response := httptest.NewRecorder()
		r := routers.InitRouter()
		r.ServeHTTP(response,request)
		want := gin.H{"tom": "喵喵"}
		var got gin.H
		err := json.NewDecoder(response.Body).Decode(&got)
		util.AssertErrShouldBeNil(err,t)
		if !reflect.DeepEqual(want, got) {
			t.Errorf("got:%v,want:%v",got,want)
		}
	})
}

func newGetRequestForTest(url string,t *testing.T) *http.Request {
	ret,err := http.NewRequest(http.MethodGet,"/tom", nil)
	if err != nil {
		t.Errorf("获取http请求失败")
	}
	return ret
}


