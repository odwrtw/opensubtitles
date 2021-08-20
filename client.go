package opensubtitles

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/golang-jwt/jwt"
)

const (
	defaultUserAgent = "osdb-go"
	defaultEndpoint  = "https://api.opensubtitles.com"
)

// Client represents a client to connect to opensubtitles
type Client struct {
	UserAgent string
	Endpoint  string
	APIKey    string
	Username  string
	Password  string
	Token     *jwt.Token
	User      *User
}

// NewClient returns a new client
func NewClient(apiKey, username, password string) *Client {
	return &Client{
		Endpoint: defaultEndpoint,
		APIKey:   apiKey,
		Username: username,
		Password: password,
	}
}

func (c *Client) request(method, url string, body io.Reader, respData interface{}, auth bool) error {
	url = c.Endpoint + url

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Api-Key", c.APIKey)

	if auth {
		// No token or token expired
		if c.Token == nil || c.Token.Claims.Valid() != nil {
			if _, err := c.Login(); err != nil {
				return err
			}
		}

		req.Header.Add("Authorization", "Bearer "+c.Token.Raw)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"opensubtitles: invalid status code %s (%d)",
			resp.Status, resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(respData)
}

func (c *Client) get(url string, resp interface{}, auth bool) error {
	return c.request("GET", url, nil, resp, auth)
}

func (c *Client) post(url string, data, resp interface{}, auth bool) error {
	body := &bytes.Buffer{}
	if err := json.NewEncoder(body).Encode(data); err != nil {
		return err
	}

	return c.request("POST", url, body, resp, auth)
}

// Search searches subtitles using a raw query
func (c *Client) Search(q SubtitleQueryParameters) ([]*SubtitleData, error) {
	resp := struct {
		TotalPages int             `json:"total_pages"`
		TotalCount int             `json:"total_count"`
		Page       int             `json:"page"`
		Data       []*SubtitleData `json:"data"`
	}{}

	if err := c.get("/api/v1/subtitles?"+q.Encode(), &resp, false); err != nil {
		return nil, err
	}

	return resp.Data, nil
}

// SearchByFile searches subtitles by file
func (c *Client) SearchByFile(path string, langs []string) ([]*SubtitleData, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hash, err := Hash(file)
	if err != nil {
		return nil, err
	}

	q := SubtitleQueryParameters{
		Query:          filepath.Base(path),
		MovieHash:      HashString(hash),
		MovieHashMatch: "only",
		Languages:      langs,
	}

	return c.Search(q)
}

// Login logs the client in
func (c *Client) Login() (*UserLogin, error) {
	credentials := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{
		Username: c.Username,
		Password: c.Password,
	}

	login := &UserLogin{}
	if err := c.post("/api/v1/login", &credentials, &login, false); err != nil {
		return nil, err
	}

	parser := &jwt.Parser{}
	claims := jwt.StandardClaims{}
	token, _, err := parser.ParseUnverified(login.Token, &claims)
	if err != nil {
		return nil, err
	}

	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	c.Token = token
	return login, nil
}

// UserInfo returns the user informations
func (c *Client) UserInfo() (*User, error) {
	out := &User{}
	data := struct {
		Data *User `json:"data"`
	}{Data: out}
	return out, c.get("/api/v1/infos/user", &data, true)
}

// DownloadSearch searches for a subtitle to download
func (c *Client) DownloadSearch(q DownloadQuery) (*DownloadResponse, error) {
	out := &DownloadResponse{}
	return out, c.post("/api/v1/download", &q, &out, true)
}

// Download downloads a file
func (c *Client) Download(fileID int, w io.Writer) error {
	q := DownloadQuery{FileID: fileID, SubFormat: "srt"}
	_, err := c.DownloadSearch(q)
	if err != nil {
		return err
	}

	// TODO implement the download when their backend works...

	return nil
}
