package config

import (
  "fmt"
  "go-file-server/log"
  "gopkg.in/yaml.v2"
  "io/ioutil"
)

var CONFIG *Config

func init() {
  yamlFile, err := ioutil.ReadFile("./config/config.yml")
  if err != nil {
    log.Error(err.Error())
  }

  err = yaml.Unmarshal(yamlFile, &CONFIG)
  if err != nil {
    fmt.Println("error", err.Error())
  }
}
