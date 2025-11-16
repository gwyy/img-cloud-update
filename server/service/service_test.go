package service

import (
	"fmt"
	"strings"
	"testing"
)

// func TestUrlMode(t *testing.T) {
// 	objectName := "test.png"
// 	path := "test/"
// 	fileName := "testdfdfd"
// 	timestampStr := fmt.Sprintf("%d", time.Now().Unix())
// 	fileNameExt := ""

// 	// h.Filename 是用户上传的文件名，如 "test.png"
// 	ext := filepath.Ext(fileName)
// 	fmt.Println(ext)
// 	if ext == "" {
// 		fileNameExt = fileName
// 	} else {
// 		fileNameExt = timestampStr + ext
// 	}
// 	fileNameExt = strings.ToLower(fileNameExt)

// 	objectName = path + time.Now().Format("200601") + "/" + fileNameExt
// 	fmt.Println(objectName)
// 	fmt.Println("aaaaa")

// }

func TestHost(t *testing.T) {
	host := "https://img1.liangtian.me/"
	if !strings.HasSuffix(host, "/") {
		host += "/"
	}
	fmt.Println(host)
}
