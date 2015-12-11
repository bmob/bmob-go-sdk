package bmob

//https://api.bmob.cn/1/images/thumbnail
type ThumbnailReq struct {
	Image string `json:"image"`

	Mode    string `json:"mode"`
	Quality string `json:"quality"`
	Width   string `json:"width"`

	OutType int `json:"outType"`
}

//https://api.bmob.cn/1/images/watermark
type WatermarkReq struct {
	Image     string `json:"image"`
	Watermark string `json:"watermark"`
	Dissolve  string `json:"dissolve"`
	DistanceX string `json:"distanceX"`
	DistanceY string `json:"distanceY"`
	Gravity   string `json:"gravity"`
	OutType   int    `json:"outType"`
}

type ImageResponse struct {
	//outType = 0
	FileName string `json:"filename"`
	Group    string `json:"group"`
	URL      string `json:"url"`
	//outType = 1
	File string `json:"file"`
}
