package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/wa1kman999/cowman/global"
	"github.com/wa1kman999/cowman/initialize"
	httpServer "github.com/wa1kman999/cowman/internal/http"
	"golang.org/x/sync/errgroup"
)

var (
	g errgroup.Group
)

func main() {
	g.Go(func() error {
		if err := httpServer.Serve(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}
		return nil
	})

	g.Go(func() error {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
		for {
			si := <-c
			switch si {
			case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
				return shutdown()
			case syscall.SIGHUP:
			default:
				return nil
			}
		}
	})

	if err := g.Wait(); err != nil {
		logrus.Printf("服务运行失败: %s", err)
		panic(err)
	}

}

// shutdown 关闭服务
func shutdown() error {
	// 关闭数据库
	db, _ := global.CMMysql.DB()
	_ = db.Close()
	// 关闭http服务
	if err := httpServer.Shutdown(); err != nil {
		return err
	}
	return nil
}

func init() {
	// 配置文件初始化
	if err := initialize.ConfigInit(); err != nil {
		panic(err)
	}

	// 初始化mysql连接
	global.CMMysql = initialize.GormMysql()
	if global.CMMysql != nil {
		// 迁移表
		initialize.RegisterTables(global.CMMysql)
	}
}
