SHELL = /bin/bash

#镜像仓库命名空间
IMAGE_NAME          = img-cloud-update

#镜像版本
TAGS_OPT           ?= latest

# 默认构建
build : build-image
	docker run -d -p 8080:80 --name ${IMAGE_NAME} ${IMAGE_NAME}:${TAGS_OPT}

# 镜像打包
build-image:
	docker build --platform linux/amd64,linux/arm64 -t ${IMAGE_NAME}:${TAGS_OPT} -f deploy/Dockerfile .

# 本地编译前端
build-web-local:
	@cd web/ && if [ -d "dist" ];then rm -rf dist; else echo "OK!"; fi \
	&& pnpm install --prod && pnpm install && pnpm run build

# 本地编译服务端
build-server-local:
	@cd server/ && if [ -f "server" ];then rm -rf server; else echo "OK!"; fi \
	&& go env -w GO111MODULE=on && go env -w GOPROXY=https://goproxy.cn,direct \
	&& go env -w CGO_ENABLED=0 && go env  && go mod tidy \
	&& GOOS=linux GOARCH=amd64 go build -o server ./cmd/main.go 

# 本地编译镜像
build-image-local:
	docker build -t ${IMAGE_NAME}:${TAGS_OPT} -f deploy/Dockerfile_Local .