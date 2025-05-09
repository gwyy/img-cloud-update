package request

type SaveAliyunSecret struct {
	Region          string `json:"region" binding:"required"`
	Bucket          string `json:"bucket" binding:"required"`
	AccessKeyId     string `json:"access_key_id" binding:"required"`
	AccessKeySecret string `json:"access_key_secret" binding:"required"`
}
