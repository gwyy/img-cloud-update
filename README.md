# 图片云上传  IMG-CLOUD-UPDATE

## 项目介绍

图片云上传是一款前端使用 vue3+element-plus，后端使用 go+gin+bbolt 的图片上传项目。使用该项目可以让你更快的把本地图片传到远程图片云仓库。当前版本支持阿里云oss。不用担心你的access_key_id会泄露，本系统保存到本地二进制数据库，不做任何远程传输。

本项目也适合初学者学习参考，内有大量的注释，对于学习 golang 和 VUE 都有比较大的帮助。

### 已支持：
- [x] 阿里云oss
- [x] 自定义路径
- [x] 自定义/自动修改名称
- [x] 自动添加markdown格式
- [x] 支持设置密码/默认安装页面
### 未支持
- [ ] 显示多层文件列表
- [ ] aws/腾讯云/七牛等更多远程仓库

### 截图
![](https://imgur.com/1kA2Bav)

![](https://imgur.com/gIbloDS)


## 依赖组件：
**1. 前端：**
   1. vue3
   2. element-plus
   3. pinia
   4. axios
   5. vue-router

**2. 后端：**
   1. go
   2. gin
   3. bbolt

## 安装
1. docker 编译安装

直接执行 make 即可，会先编译 web 端，再编译 server 端。再打包容器镜像。最后 docker 运行镜像。





## 参考项目：
- https://github.com/etcd-io/bbolt  二进制文本数据库
- https://help.aliyun.com/zh/oss/developer-reference/quick-start-for-oss-go-sdk-v2?spm=a2c4g.11186623.0.0.4ac75a05vzGr7u  阿里云 oss go v2 sdk