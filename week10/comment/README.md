# Kratos Project Template

## Install Kratos
```
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```
项目启动过程

1、启动本地mysql、redis、kafka
2、在comment目录执行kratos run    启动comment http服务和grpc服务
3、在user目录执行kratos run       启动user http服务和grpc服务
4、执行go run internal/job/commentSubcribe.go 订阅kafka

测试
浏览器访问  http://127.0.0.1:8000/comment/add?uid=2&parent_id=7&content=666&relation_id=1&relation_type=1 进行添加评论
     访问  http://127.0.0.1:8000/comment/list?page=1&relation_id=1&relation_type=1 获取评论列表
