package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"server/config"
	"server/global"

	"gopkg.in/yaml.v2"
)
// 加载配置文件
func InitConfig() {
	const ConfigFile = "settings.yaml"
	c := &config.Config{}
	b, err := ioutil.ReadFile(ConfigFile)
	if err != nil {
	 panic(fmt.Errorf("get yamlConf error: %s",err))
	}
	err = yaml.Unmarshal(b,c)
	if err != nil {
		log.Fatalf("config Init Unmarshal: %v",err)
	}
	log.Println("config yamlFaie load Init success.")
	global.Config = c
}
