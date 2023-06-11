package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})
	router.Run(":8080")

}

// package main

// import (
// 	"employee-app/config"
// 	"employee-app/router"
// 	"fmt"
// )

//	func main() {
//		conf, err := config.Init()
//		if err != nil {
//			panic(fmt.Sprintf("Unable to load the config details: %v ", err.Error()))
//		}
//		server := router.Server{Gin: conf.Gin}
//		server.Start(conf)
//	}
// package main

// func main() {

// 	print("int func initialising")
// 	s1, err := Init()
// 	print("int func initialized")

// 	if err != nil {
// 		fmt.Println("server creation failed")
// 	}
// 	print("starting the server")
// 	s1.Start()
// }
// func print(str string) {
// 	fmt.Println(str)
// }

// type Config struct {
// 	port     int
// 	hostname string
// 	mode     string
// }

// type GinConfig struct {
// 	engine *gin.Engine
// 	Config Config
// }
// type Server struct {
// 	ginserver *GinConfig
// }

// func (server *Server) Start() error {
// 	print("server start function")
// 	gin.SetMode(server.ginserver.Config.mode)
// 	server.SetUpRoutes()
// 	server.Run()
// 	return nil
// }
// func (server *Server) SetUpRoutes() {
// 	print("set up routes function")
// 	routergroups := server.ginserver.engine.Group("emp")
// 	routergroups.GET("/", func(ctx *gin.Context) {
// 		ctx.JSON(http.StatusOK, gin.H{"msg": "hi"})
// 	})

// }
// func (server *Server) Run() {
// 	print("run function")
// 	err := server.ginserver.engine.Run(fmt.Sprintf("%s:%d", server.ginserver.Config.hostname, server.ginserver.Config.port))
// 	if err != nil {
// 		fmt.Println("server did not start fine")
// 	}
// }

// func Init() (*Server, error) {
// 	conf := Config{}
// 	print("conf object created")
// 	conf.port = 8080
// 	conf.hostname = "localhost"
// 	conf.mode = "debug"
// 	print("conf object values set")
// 	server := Server{}
// 	ginconf := GinConfig{}
// 	print("server object initialized")
// 	trial := gin.New()
// 	fmt.Sprintln("hi ", trial)
// 	print("initialised trial gin engine")
// 	ginconf.engine = trial
// 	print("engine initialized")
// 	ginconf.Config = conf
// 	server.ginserver = &ginconf
// 	print("conf initialized")
// 	print("server object values set")

// 	return &server, nil
// }
