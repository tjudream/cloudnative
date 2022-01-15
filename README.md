# cloudnative
云原生训练营第一期
# 使用 Dockerfile 制作 Docker 镜像
```shell
go build httpserver/main.go
docker build --tag tjudream/httpserver:v1 .
docker run -d -p 80:8080 --name httpserverv1 tjudream/httpserver:v1
docker logs -f httpserverv1
docker login docker.io
docker push tjudream/httpserver:v1
```
# 使用 Docker 镜像
```shell
docker pull tjudream/httpserver:v1
docker run -d -p 80:8080 --name httpserverv1 tjudream/httpserver:v1
docker logs -f httpserverv1
```
