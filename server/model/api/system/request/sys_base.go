package request

type SaveAliyunSecret struct {
	AccessKeyId     string `json:"access_key_id" binding:"required"`
	AccessKeySecret string `json:"access_key_secret" binding:"required"`
}
