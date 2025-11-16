package bbolt_manager

import (
	"time"

	"github.com/gwyy/img-cloud-update/server/global"
	"go.etcd.io/bbolt"
)

// OpenDB 打开数据库
func OpenDB(path string) (*bbolt.DB, error) {
	return bbolt.Open(path, 0600, &bbolt.Options{Timeout: 1 * time.Second})
}

// CreateTable 创建表
func CreateTable(db *bbolt.DB, tableName string) error {
	return db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(tableName))
		return err
	})
}

// Save 保存数据
func Save(tableName string, key string, value string) error {

	return global.BboltDB.Update(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		return bucket.Put([]byte(key), []byte(value))
	})
}

// Get 获取数据
func Get(tableName string, key string) (string, error) {
	var value string
	err := global.BboltDB.View(func(tx *bbolt.Tx) error {
		bucket := tx.Bucket([]byte(tableName))
		value = string(bucket.Get([]byte(key)))
		return nil
	})
	return value, err
}
