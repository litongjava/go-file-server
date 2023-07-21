package controller

import (
  "fmt"
  "github.com/google/uuid"
  "go-file-server/config"
  "io"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "regexp"
  "time"
)

func registerFileRouter() {
  http.HandleFunc("/file/upload", handleUpload)
  http.HandleFunc("/file/download/", handleDownload)
}

//上传文件
func handleUpload(writer http.ResponseWriter, request *http.Request) {
  file, header, err := request.FormFile("file")
  if err != nil {
    return
  }

  //上传目录
  timeNow := time.Now()
  dateString := timeNow.Format("2006-01-02")
  uuidString := uuid.New().String()

  savePath := config.CONFIG.App.FilePath + "/" + dateString + "/" + uuidString
  isExists := IsExist(savePath)
  if !isExists {
    log.Println("create path", savePath)
    os.MkdirAll(savePath, os.ModePerm)
  }
  //读取文件名
  log.Println("filename", header.Filename)

  //上传目录 uploadPath+/+日期+uuid+filename
  filename := savePath + "/" + header.Filename
  //创建文件
  openFile, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0666)
  if err != nil {
    log.Println(err.Error())
    return
  }
  defer file.Close()
  defer openFile.Close()
  io.Copy(openFile, file)
  fmt.Fprintln(writer, dateString+"/"+uuidString)
}

//下载文件
func handleDownload(writer http.ResponseWriter, request *http.Request) {
  pattern, _ := regexp.Compile(`/file/download/(.+)`)
  matches := pattern.FindStringSubmatch(request.URL.Path)
  if len(matches) > 0 {
    subDir := matches[1]
    log.Println("subDir:", subDir)
    savePath := config.CONFIG.App.FilePath + "/" + subDir
    fileInfoList, err := ioutil.ReadDir(savePath)
    if err != nil {
      log.Println("读取文件列表失败", err.Error())
      fmt.Fprint(writer, "读取文件列表失败", err.Error())
      return
    }
    if len(fileInfoList) > 0 {
      filename := fileInfoList[0].Name()
      filePath := savePath + "/" + filename
      bytes, err := os.ReadFile(filePath)
      if err != nil {
        log.Println("读取文件失败", err)
        return
      }
      writer.Header().Add("Content-Type", "application/octet-stream")
      writer.Header().Add("Content-Disposition", "attachment; filename= "+filename)
      writer.Write(bytes)
    } else {
      log.Println("没有读取到文件")
      fmt.Fprint(writer, "没有读取到文件")
    }

  } else {
    fmt.Fprint(writer, "请求输入正确的路径")
  }
}

func IsExist(path string) bool {
  _, err := os.Stat(path)
  return err == nil || os.IsExist(err)
}
