package opensubtitles

import (
	"net/url"
	"strconv"
	"strings"
)

// SubtitleQueryParameters represents the allowed search parameters
type SubtitleQueryParameters struct {
	AITranslated      string   // ai_translated / exclude, include (default: exclude)
	EpisodeNumber     int      // episode_number
	ForeignPartsOnly  string   // foreign_parts_only / include, only (default: include)
	HearingImpaired   string   // hearing_impaired / include, exclude, only. (default: include)
	ID                string   // id
	ImdbID            string   // imdb_id without tt
	Languages         []string // languages: coma separated (en,fr)
	MachineTranslated string   // machine_translated
	MovieHash         string   // moviehash
	MovieHashMatch    string   // moviehash_match
	OrderBy           string   // order_by
	OrderDirection    string   // order_direction
	Page              int      // page
	ParentFeatureID   int      // parent_feature_id
	ParentImdbID      int      // parent_imdb_id
	ParentTmdbID      int      // parent_tmdb_id
	Query             string   // query string
	SeasonNumber      int      // season_number
	TmdbID            string   // tmdb_id
	TrustedSources    string   // trusted_sources / include, only (default: include)
	Type              string   // type / movie, episode or all, (default: all)
	UserID            string   // user_id
	Year              int      // year
}

// Encode encodes the query parameters
func (q *SubtitleQueryParameters) Encode() string {
	v := url.Values{}

	for key, value := range map[string]string{
		"ai_translated":      q.AITranslated,
		"foreign_parts_only": q.ForeignPartsOnly,
		"hearing_impaired":   q.HearingImpaired,
		"id":                 q.ID,
		"imdb_id":            q.ImdbID,
		"machine_translated": q.MachineTranslated,
		"moviehash":          q.MovieHash,
		"moviehash_match":    q.MovieHashMatch,
		"order_by":           q.OrderBy,
		"order_direction":    q.OrderDirection,
		"query":              q.Query,
		"tmdb_id":            q.TmdbID,
		"trusted_sources":    q.TrustedSources,
		"type":               q.Type,
		"user_id":            q.UserID,
		"languages":          strings.Join(q.Languages, ","),
	} {
		if value != "" {
			v.Add(key, value)
		}
	}

	for key, value := range map[string]int{
		"episode_number":    q.EpisodeNumber,
		"page":              q.Page,
		"parent_feature_id": q.ParentFeatureID,
		"parent_imdb_id":    q.ParentImdbID,
		"parent_tmdb_id":    q.ParentTmdbID,
		"season_number":     q.SeasonNumber,
		"year":              q.Year,
	} {
		if value != 0 {
			v.Add(key, strconv.Itoa(value))
		}
	}

	return v.Encode()
}
