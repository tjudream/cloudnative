# cloudnative
云原生训练营第一期
# 使用 Dockerfile 制作 Docker 镜像
```shell
docker build --tag tjudream/httpserver:v1 .
docker run -d -p 80:8080 --name httpserver tjudream/httpserver:v1
docker logs -f httpserver
docker login docker.io
docker push tjudream/httpserver:v1
```
# 使用 Docker 镜像
```shell
docker pull tjudream/httpserver:v1
docker run -d -p 80:8080 --name httpserver tjudream/httpserver:v1
docker logs -f httpserver
```
# 使用 nsenter 进入容器查看 IP 配置
```shell
PID=$(docker inspect --format "{{ .State.Pid }}" httpserver)
nsenter -t $PID -n ip a
```
