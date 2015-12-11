package bmob

type ACL struct {
	Read  bool `json:"read"`
	Write bool `json:"write"`
}

type Role struct {
	Name  string      `json:"name"`
	Acl   interface{} `json:"ACL"`
	Roles interface{} `json:"roles"`
	Users interface{} `json:"users"`
}
