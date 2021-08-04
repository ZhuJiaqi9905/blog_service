package main

import (
	"blog-service/global"
	"blog-service/internal/routers"
	"blog-service/pkg/setting"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"reflect"
	"time"
)

func main(){
	gin.SetMode(global.ServerSetting.RunMode)
	router := routers.NewRouter()
	s := &http.Server{
		Addr: ":" + global.ServerSetting.HttpPort,
		Handler: router,
		ReadTimeout: global.ServerSetting.ReadTimeout,
		WriteTimeout: global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()

}
func init(){
	err := setupSetting()
	if err != nil{
		log.Fatalf("init.setupSetting err: %v", err)
	}
}
func setupSetting() error{
	setting, err := setting.NewSetting()
	if err != nil{
		return err
	}
	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("App", &global.AppSetting)
	if err != nil{
		return err
	}
	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil{
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	return nil
}
func test()  {
	type S struct {
		F string `species:"gopher" color:"blue"`
	}
	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)

	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
}
