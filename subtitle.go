package opensubtitles

import "time"

// SubtitleData holds the subtitle
type SubtitleData struct {
	ID       string   `json:"id"`
	Type     string   `json:"type"`
	Subtitle Subtitle `json:"attributes"`
}

// Subtitle represents a subtile response
type Subtitle struct {
	ID               string         `json:"subtitle_id"`
	Language         string         `json:"language"`
	DownloadCount    int            `json:"download_count"`
	NewDownloadCount int            `json:"new_download_count"`
	HearingImpared   bool           `json:"hearing_impared"`
	HD               bool           `json:"hd"`
	FPS              float64        `json:"fps"`
	Votes            int            `json:"votes"`
	Points           int            `json:"points"`
	Rating           int            `json:"rating"`
	FromTrusted      bool           `json:"from_trusted"`
	ForeignPartsOnly bool           `json:"foreign_parts_only"`
	AutoTranslation  bool           `json:"auto_translation"`
	AITranslated     bool           `json:"ai_translated"`
	UploadDate       time.Time      `json:"upload_date"`
	Release          string         `json:"release"`
	Comments         string         `json:"comments"`
	LegacySubtitleID int            `json:"legacy_subtitle_id"`
	URL              string         `json:"url"`
	FeatureDetails   FeatureDetails `json:"feature_details"`
	Uploader         Uploader       `json:"uploader"`
	RelatedLinks     RelatedLinks   `json:"related_links"`
	Files            []File         `json:"files"`
	MovieHashMatch   bool           `json:"movie_hash_match"`

	// Format and MachineTranslated are not listed because they are not
	// required and their type it not defined in the documentation...
}

// Uploader represents an uploader
type Uploader struct {
	ID   int    `json:"uploader_id"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}

// FeatureDetails holds the details of the video this subtitle refers to
type FeatureDetails struct {
	ID        int    `json:"feature_id"`
	Type      string `json:"feature_type"`
	Year      int    `json:"year"`
	Title     string `json:"title"`
	MovieName string `json:"movie_name"`
	ImdbID    int    `json:"imdb_id"`
	TmdbID    int    `json:"tmdb_id"`
}

// File represents a file
type File struct {
	ID       int    `json:"file_id"`
	CDNumber int    `json:"cd_number"`
	FileName string `json:"file_name"`
}

// RelatedLinks holds some link
type RelatedLinks struct {
	Label  string `json:"label"`
	URL    string `json:"url"`
	ImgURL string `json:"img_url"`
}
