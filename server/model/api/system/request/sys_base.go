package request

type SaveAliyunSecret struct {
	Region          string `json:"region" binding:"required"`
	Host            string `json:"host" binding:"required"`
	Bucket          string `json:"bucket" binding:"required"`
	AccessKeyId     string `json:"access_key_id" binding:"required"`
	AccessKeySecret string `json:"access_key_secret" binding:"required"`
	DefaultPath     string `json:"default_path" `
	SysPassword     string `json:"sys_password" binding:"required"`
	UrlMode         string `json:"url_mode" binding:"required"`
	CopyImgFormat   string `json:"copy_img_format" binding:"required"`
}

type Login struct {
	Password string `json:"password" binding:"required"`
}
