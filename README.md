# CX-Panel-Server
Server of CX-Panel

基于beego开发的服务器网站运维面板服务端

## 运行
```bash
bee run -gendoc=true -downdoc=true
```

这种运行在dev模式下会自动生成路由、自动更新API文档

默认监听`8888`端口

本地访问：`http://127.0.0.1:8888/`

访问根目录`/`为前端主页，api均在`/api/`下

## API文档
```bash
http://127.0.0.1:8888/swagger/
```
## 当前开发进度

仅搭建框架，能够接收RESTful请求。

## 接下来开发目标

- 管理员（Admin）的控制器和模型
- 使用SQLite3作为数据库