package bmob

type Installation struct {
	Badge           int      `json:"badge"`
	Channels        []string `json:"channels"`
	TimeZone        string   `json:"timeZone"`
	DeviceType      string   `json:"deviceType"`
	InstallationId  string   `json:"installationId"`
	DeviceToken     string   `json:"deviceToken"`
	NotificationURI string   `json:"notificationUri"`
}
