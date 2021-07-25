package util_test

import (
	"github.com/siconghe/blog/pkg/setting"
	"github.com/siconghe/blog/pkg/util"
	"testing"
)

type MockQueryable struct {
	Keys map[string]string
}

func (mq *MockQueryable)Query(key string) string{
	ret,ok := mq.Keys[key]
	if ok {
		return ret
	}
	return "error"
}

func TestPagination(t *testing.T) {
	mq := &MockQueryable{
		map[string]string{"page":"11"},
	}
	want := (11 - 1 ) * setting.PageSize
	got := util.GetPage(mq)
	if want != got {
		t.Errorf("Pagination测试失败，got: %v, want: %v",got,want)
	}
}
