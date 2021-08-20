package opensubtitles

// DownloadQuery represents the query to download a subtitle
type DownloadQuery struct {
	FileID       int    `json:"file_id"` // mandatory
	SubFormat    string `json:"sub_format"`
	FileName     string `json:"file_name"`
	StripHTML    bool   `json:"strip_html"`
	CleanupLinks bool   `json:"cleanup_links"`
	RemoveAdds   bool   `json:"remove_adds"`
	InFPS        int    `json:"in_fps"`
	OutFPS       int    `json:"out_fps"`
	Timeshift    int    `json:"timeshift"`
}

// DownloadResponse represents the response when requesting for a file to
// download
type DownloadResponse struct {
	Link      string `json:"link"`
	Fname     string `json:"fname"`
	Requests  int    `json:"requests"`
	Allowed   int    `json:"allowed"`
	Remaining int    `json:"remaining"`
	Message   string `json:"message"`
}
