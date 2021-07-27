package main

import (
	"fmt"
	"github.com/siconghe/blog/models"
	"github.com/siconghe/blog/pkg/setting"
	"github.com/siconghe/blog/routers"
	"net/http"
)

func main()  {
	models.InitDB()
	router := routers.InitRouter()
	server := &http.Server{
		Addr:              fmt.Sprintf("0.0.0.0:%d", setting.HTTPPort),
		Handler:           router,
		TLSConfig:         nil,
		ReadTimeout:       setting.ReadTimeout,
		ReadHeaderTimeout: 0,
		WriteTimeout:      setting.WriteTimeout,
		IdleTimeout:       0,
		MaxHeaderBytes:    1 << 20, //1<<20也就是1*2^20=1MB
		TLSNextProto:      nil,
		ConnState:         nil,
		ErrorLog:          nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}
	server.ListenAndServe()

}
