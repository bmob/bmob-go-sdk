package bmob

type BaseUserInfo struct {
	//Must have
	UserName string `json:"username"`
	Password string `json:"password"`
	//Optional
	Email             string `json:"email"`
	MobilePhoneNumber string `json:"mobilePhoneNumber"`
}
