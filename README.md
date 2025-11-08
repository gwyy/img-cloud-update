# 图片云更新  IMG-CLOUD-UPDATE

## 项目介绍

图片云更新是一款前端使用 vue3+element-plus，后端使用 go+gin+bbolt 的图片云更新项目。使用该项目可以让你更快的把本地图片传到远程图片云仓库。

### 已支持：
- 阿里云oss
- 自动修改名称
- 列表展示
- 安装页面



## 依赖组件：
1. 前端：
   1. vue3
   2. element-plus
   3. pinia
   4. axios
   5. vue-router
2. 后端：
   1. go
   2. gin
   3. bbolt

## 安装
1. docker直接安装：
```bash
docker run -d --name img-cloud-update -p 8888:8888 -v ./data:/app/data -v ./logs:/app/logs img-cloud-update
```


参考项目：
- https://github.com/etcd-io/bbolt  二进制文本数据库
- https://help.aliyun.com/zh/oss/developer-reference/quick-start-for-oss-go-sdk-v2?spm=a2c4g.11186623.0.0.4ac75a05vzGr7u  oss go v2 sdk