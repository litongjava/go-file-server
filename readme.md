## 文件服务器
接口地址  
https://www.apifox.cn/apidoc/shared-07b05f71-0cee-4886-b40c-3013c14fdcff

使用curl上传
```shell
curl --location --request POST 'http://127.0.0.1:10406/file/upload' \
--form 'file=@"F:\\document\\API文档\\python\\Python3.6 中文文档.pdf"'
```
下载
```shell
http://127.0.0.1:10406/file/download/2022-11-21/1e4fa8b0-f8e4-4578-a681-58136999a5a0
```