package ws_conn

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"goim/conf"
	"goim/internal/ws_conn/service"
	"goim/pkg/net/http/middleware"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Init(c *conf.Config) {
	service.New(c)

	gin.SetMode(gin.ReleaseMode)
	e := gin.New()
	e.Use(gin.Recovery(), middleware.CORS())
	e = initRouter(e)

	srv := http.Server{
		Addr:    conf.Conf.WsHttp.Addr,
		Handler: e,
	}
	go func() {
		log.Printf("Listening and serving HTTP on %s\n", conf.Conf.WsHttp.Addr)
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("ws_conn.Init() error(%v)", err))
		}
	}()

	// 开启pprof
	go func() {
		log.Printf("Listening and serving pprof HTTP on %s\n", conf.Conf.WsHttp.PProfAddr)
		err := http.ListenAndServe(c.WsHttp.PProfAddr, nil)
		if err != nil && err != http.ErrServerClosed {
			panic(fmt.Errorf("ws_conn.Init() pprof HTTP error(%v)", err))
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGSTOP, syscall.SIGINT)
	s := <-quit
	log.Printf("ws_conn http server get a signal %s", s.String())
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("ws_conn http server shutdown:", err)
	}
	log.Println("ws_conn http server exiting")
}

func initRouter(e *gin.Engine) *gin.Engine {
	g := e.Group("/v1")
	{
		g.GET("/ws", connWs)
	}

	return e
}
