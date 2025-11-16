package initialize

import (
	"context"
	"log"
	"testing"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// 定义全局变量
var (
	region          string = ""
	bucketName      string = ""
	accessKeyID     string = ""
	accessKeySecret string = ""
)

func TestAliyunOssInit(t *testing.T) {
	//AddObject("myblog/imgs/test/doc.txt", "/Users/johnny/wwwroot/person/img-cloud-update/docs/doc.txt")
	//DeleteObject("myblog/imgs/test/doc.txt")

}

func TDeleteObject(objectName string) {
	client := TNewClient()
	// 创建删除对象的请求
	request := &oss.DeleteObjectRequest{
		Bucket: oss.Ptr(bucketName), // 存储空间名称
		Key:    oss.Ptr(objectName), // 对象名称
	}

	// 执行删除对象的操作并处理结果
	result, err := client.DeleteObject(context.TODO(), request)
	if err != nil {
		log.Fatalf("failed to delete object %v", err)
	}

	// 打印删除对象的结果
	log.Printf("delete object result:%#v\n", result)
}

func TAddObject(objectName string, localFile string) {
	client := TNewClient()
	// 创建上传对象的请求
	putRequest := &oss.PutObjectRequest{
		Bucket:       oss.Ptr(bucketName),      // 存储空间名称
		Key:          oss.Ptr(objectName),      // 对象名称
		StorageClass: oss.StorageClassStandard, // 指定对象的存储类型为标准存储
		Acl:          oss.ObjectACLPrivate,     // 指定对象的访问权限为私有访问
	}

	// 执行上传对象的请求
	result, err := client.PutObjectFromFile(context.TODO(), putRequest, localFile)
	if err != nil {
		log.Fatalf("failed to put object from file %v", err)
	}

	// 打印上传对象的结果
	log.Printf("put object from file result:%#v\n", result)
}

func TNewClient() *oss.Client {

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig().
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, accessKeySecret)).
		WithRegion(region)

	// 创建OSS客户端
	client := oss.NewClient(cfg)
	return client
}
