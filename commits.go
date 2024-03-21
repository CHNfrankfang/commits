package main

import (
	"flag"
	"fmt"
	"log"

	"commits/internal/config"
	"commits/internal/handler"
	"commits/internal/model"
	"commits/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

type JobConfig struct {
	// service.ServiceConf
	config.Config
}

var configFile = flag.String("f", "etc/commits-api.toml", "the config file")

func main() {
	flag.Parse()

	var c JobConfig
	conf.MustLoad(*configFile, &c)
	err := c.Conf.SetUp()
	if err != nil {
		log.Fatal(err)
	}

	// var lc logc.LogConf = logx.LogConf{
	// 	Encoding:   "plain",
	// 	TimeFormat: "2006-01-02T15:04:05Z07:00",
	// }
	// logc.MustSetup(lc)

	server := rest.MustNewServer(c.Conf)
	defer server.Stop()

	model.SQLX(c.Mysql.DSN)
	ctx := svc.NewServiceContext(c.Config)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Conf.Host, c.Conf.Port)
	server.Start()
}
