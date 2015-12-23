package bmob

//https://api.bmob.cn/1/images/thumbnail
type ImageReq struct {
	Image   string `json:"image"`
	OutType int    `json:"outType"`
}
type ThumbnailReq struct {
	ImageReq
	Mode    string `json:"mode"`
	Quality string `json:"quality"`
	Width   string `json:"width"`
}

//https://api.bmob.cn/1/images/watermark
type WatermarkReq struct {
	ImageReq
	Watermark string `json:"watermark"`
	Dissolve  string `json:"dissolve"`
	DistanceX string `json:"distanceX"`
	DistanceY string `json:"distanceY"`
	Gravity   string `json:"gravity"`
}

type ImageResponse struct {
	//outType = 0
	FileName string `json:"filename"`
	Group    string `json:"group"`
	URL      string `json:"url"`
	//outType = 1
	File string `json:"file"`
}
