package main

import (
  "flag"
  "go-file-server/config"
  "go-file-server/controller"
  "log"
  "net/http"
  "strconv"
)

func init() {
  //设置Flats为 日期 时间 微秒 文件名:行号
  log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
}
func main() {
  var configFilePath string
  flag.StringVar(&configFilePath, "c", "config.yml", "Configuration file path")
  flag.Parse()

  // 读取配置
  config.ReadFile(configFilePath)

  port := strconv.Itoa(config.CONFIG.App.Port)
  log.Println("start listen on", port)
  controller.RegisterRoutes()
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    log.Fatalln(err.Error())
  }
}
