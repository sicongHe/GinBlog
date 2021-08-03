package main

import (
	"context"
	"fmt"
	_ "github.com/siconghe/blog/docs"
	"github.com/siconghe/blog/models"
	"github.com/siconghe/blog/pkg/logging"
	"github.com/siconghe/blog/pkg/setting"
	"github.com/siconghe/blog/routers"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)
//...

// @title Gin Blog
// @version 1.0
func main()  {
	setting.Setup()
	models.Setup()
	logging.Setup()
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<- quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Server exiting")

}
