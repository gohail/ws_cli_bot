package config

import "fmt"

type AppConf struct {
	Host string
	Port int
}

func (c AppConf) GetHostString() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}
