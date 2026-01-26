########### 第一阶段：构建前端 (web)
FROM node:20-alpine AS build-web
WORKDIR /build/web
# 复制 web 目录
COPY web/ ./
# 删除旧的 dist（如果存在），安装依赖并构建
RUN if [ -d "dist" ]; then rm -rf dist; else echo "OK!"; fi \
    && npm install \
    && npm run build

########### 第二阶段 编译golang 项目
FROM golang:1.24-alpine AS build-server
WORKDIR /build/server
# 复制 server 目录
COPY server/ ./
# 设置 Go 环境，tidy mod 并构建   删除旧的 server（如果存在），但由于这是新构建，通常不需要，但为了匹配 Makefile
RUN if [ -f "server" ]; then echo "Built successfully"; else echo "Build failed"; fi \
    && apk add --no-cache git \
    && go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go mod tidy \
    && GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go 
   
########### 第三阶段 打包项目
FROM alpine:3.19
WORKDIR /opt/img-cloud-update
ENV LANG=en_US.utf8
# 安装运行时依赖
RUN apk add --no-cache nginx musl-locales bash
# 复制 entrypoint 和其他配置文件（假设 deploy/ 和 server/ 在项目根目录）
COPY deploy/entrypoint.sh /opt/img-cloud-update/entrypoint.sh
COPY server/config.yaml /opt/img-cloud-update/server/config.yaml
COPY server/data /opt/img-cloud-update/server/data/
COPY deploy/nginx.conf /etc/nginx/http.d/nginx.conf
# 从构建阶段复制产物
COPY --from=build-web /build/web/dist/ /opt/img-cloud-update/web/
COPY --from=build-server /build/server/server /opt/img-cloud-update/server/server
# 设置权限和 locale
RUN echo "LANG=en_US.utf8" > /etc/locale.conf \
    && chmod +x ./entrypoint.sh
EXPOSE 80
ENTRYPOINT ["/bin/bash", "./entrypoint.sh"]