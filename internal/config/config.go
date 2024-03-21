package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	Conf  rest.RestConf
	Mysql struct {
		DSN string
	}
}
