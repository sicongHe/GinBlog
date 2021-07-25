package util

import (
	"github.com/siconghe/blog/pkg/setting"
	"github.com/unknwon/com"
)

type Queryable interface {
	Query(key string) string
}

func GetPage(c Queryable) int{
	ret :=0
	page,_ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		ret = (page - 1) * setting.PageSize
	}
	return ret
}
