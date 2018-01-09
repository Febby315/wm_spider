package updownload

type (
	// Param represents the structure of our resource
	Param struct {
		ImgName string `json:"img_name"`
		ImgURL  string `json:"img_url"`
		ImgSrc  string `json:"img_src"`
	}
)
