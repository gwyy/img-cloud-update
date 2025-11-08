package response

type FileDownload struct {
	Name string `json:"name" form:"name" gorm:"column:name;comment:文件名"` // 文件名
	Url  string `json:"url" form:"url" gorm:"column:url;comment:文件地址"`   // 文件地址
	Key  string `json:"key" form:"key" gorm:"column:key;comment:编号"`     // 编号
}
