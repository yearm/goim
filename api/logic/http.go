package logic

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goim/conf"
	"goim/internal/logic/server"
	"goim/pkg/net/http/middleware"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	srv *server.Server
)

func Init(c *conf.Config) {
	srv = server.New(c)

	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery(), middleware.CORS())
	e = initRouter(e)

	srv := http.Server{
		Addr:    conf.Conf.LogicAddrs.Addr,
		Handler: e,
	}
	go func() {
		log.Printf("Listening and serving HTTP on %s\n", conf.Conf.LogicAddrs.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("logic.Init() HTTP server error(%v)", err))
		}
	}()

	// 开启pprof
	go func() {
		log.Printf("Listening and serving pprof HTTP on %s\n", conf.Conf.LogicAddrs.PProfAddr)
		err := http.ListenAndServe(c.LogicAddrs.PProfAddr, nil)
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("logic.Init() pprof HTTP error(%v)", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT)
	s := <-quit
	log.Printf("logic http server get a signal %s", s.String())
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("logic http server shutdown:", err)
	}
	log.Println("logic http server exiting")
}

func initRouter(e *gin.Engine) *gin.Engine {
	g := e.Group("/v1")
	{
		g.POST("/send", sendMsg)
	}

	return e
}
