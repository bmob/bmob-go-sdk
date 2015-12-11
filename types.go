package bmob

type Op struct {
	Op      string        `json:"__op"`
	Objects []interface{} `json:"objects"`
}

type RestConfig struct {
	AppID   string
	RestKey string
}

type BaseReq struct {
	Method string `json:"method"`
	Path   string `json:"path"`
	Token  string `json:"token"`
}

//for RAW HTTP
type RestRequest struct {
	BaseReq
	Type string
	Body []byte
}

//for BATCH request
type BatchRequest struct {
	BaseReq
	Body interface{} `json:"body"`
}

/*
 * RESTful response
 */
type RestResponse struct {
	/*
	 * Error related
	 */
	Code  int    `json:"code"`
	Error string `json:"error"`

	Message string `json:"msg"`

	/*
	 * Bmob data related
	 */
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	ObjectId  string `json:"objectId"`

	/*
	 * Bmob query related
	 */
	Count   string        `json:"count"`
	Results []interface{} `json:"results"`

	/*
	 * Bmob user related
	 */
	SessionToken string `json:"sessionToken"`
}

/*
 * Bmob-defined data type info
 */
type DataType struct {
	Type string `json:"__type"`
	//Date
	ISODate string `json:"iso,emitempty"`
	//File
	Group    string `json:"group,emitempty"`
	FileName string `json:"filename,emitempty"`
	URL      string `json:"url,emitempty"`
	//Pointer
	ClassName string `json:"className,emitempty"` //Also for Relation
	ObjectId  string `json:"objectId,emitempty"`
	//GeoLocation
	Latitude   string `json:"latitude,emitempty"`
	Longtitude string `json:"longitude,emitempty"`
}
