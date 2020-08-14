# Go Build Option
------
## 静态编译，在scratch镜像中使用很有运行必须是静态编译的
go build --ldflags '-linkmode external -extldflags "-static"'
