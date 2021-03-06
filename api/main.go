package main

import (
	"api/pkg/setting"
	"api/routers"
	"fmt"
	"net/http"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://github.com/junun/admin-go

// @contact.name Junun
// @contact.url https://github.com/junun/admin-go
// @contact.email junun717@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        r,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
