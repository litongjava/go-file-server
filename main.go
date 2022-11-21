package main

import (
  "go-file-server/config"
  "go-file-server/controller"
  "go-file-server/log"
  "net/http"
  "os"
  "strconv"
)

func main() {
  //port := "10406"
  port := strconv.Itoa(config.CONFIG.App.Port)
  for i := 1; i < len(os.Args); i += 2 {
    param := os.Args[i]
    if param == "--port" {
      port = os.Args[i+1]
    }
  }
  log.Info("start listen on", port)
  controller.RegisterRoutes()
  err := http.ListenAndServe(":"+port, nil)
  if err != nil {
    log.Error(err.Error())
  }
}
