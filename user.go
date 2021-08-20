package opensubtitles

// User represents the user
type User struct {
	AllowedDownloads   int    `json:"allowed_downloads"`
	Level              string `json:"level"`
	ID                 int    `json:"user_id"`
	ExtInstalled       bool   `json:"ext_installed"`
	VIP                bool   `json:"vip"`
	DownloadCount      int    `json:"download_count"`
	RemainingDownloads int    `json:"remaining_downloads"`
}

// UserLogin represents the data returned during the login process
type UserLogin struct {
	User   User   `json:"user"`
	Token  string `json:"token"`
	Status int    `json:"status"`
}
