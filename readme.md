## go-file-server 
go-file-server 是一个文件服务器,包含了上传和下载功能

## Install
```
mkdir /opt/package/go-file-server -p
cd /opt/package/go-file-server
wget https://gitee.com/ppnt/go-file-server/releases/download/v1.0/go-file-server-linux-amd64-1.0.0.zip
unzip go-file-server-linux-amd64-1.0.0.zip -d /opt/
cd /opt/go-file-server-linux-amd64-1.0.0/
```
```
mkdir config
vi config/config.yml
```

```
app:
  port: 10406
  filePath: /data/upload
```
```
chmod +x go-file-server
./go-file-server
```
## upload file
```
curl --location --request POST 'http://127.0.0.1:10406/file/upload' \
--form 'file=@"/root/code/project/go-file-server/readme.md"'
```
repsponse
```
2023-07-17/35a3a36d-0aa4-4de0-822f-09d50ffef056
```

Directory structure
```
/data/upload/
└── 2023-07-17
└── 35a3a36d-0aa4-4de0-822f-09d50ffef056
└── readme.md
```

## wownload file
http://127.0.0.1:10406/file/download/2023-07-17/35a3a36d-0aa4-4de0-822f-09d50ffef056
