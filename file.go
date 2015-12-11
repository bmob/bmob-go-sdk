package bmob

//http://file.bmob.cn/M00/00/01/wKgBP1N3FAWRJXsSAAAB_rYZATs52.html
type FileResponse struct {
	Filename string `json:"filename"`
	Group    string `json:"group"`
	URL      string `json:"url"`
}

/*
crc32：文件片段的crc32校验码。
ctx：服务端不透明的内部参数，需要在下一次上传时提交。
file： 如果文件没有完整上传，返回false, 否则返回文件名、分组和文件地址。
host：下一个片的上传URL（注意：不固定）。
offset：下一个片的偏移位置。
*/

type SliceResponse struct {
	Crc32  string `json:"crc32"`
	Ctx    string `json:"ctx"`
	File   string `json:"file"`
	Host   string `json:"host"`
	Offset int    `json:"offset"`
}
