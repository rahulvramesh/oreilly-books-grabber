package models

import "time"

//LoginResponse -
type LoginResponse struct {
	IDToken      string `json:"id_token"`
	RefreshToken string `json:"refresh_token"`
	LoggedIn     bool   `json:"logged_in"`
	RedirectURI  string `json:"redirect_uri"`
	UUID         string `json:"uuid"`
}

//LoginPayload -
type LoginPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//BookContext -
type BookContext struct {
	BookID    string `json:"book_id"`
	Publisher struct {
		Slug string `json:"slug"`
		Logo string `json:"logo"`
		Name string `json:"name"`
	} `json:"publisher"`
	TitleSafe      string  `json:"title_safe"`
	DetailURL      string  `json:"detail_url"`
	DetailPosition float64 `json:"detail_position"`
	ThumbnailTag   string  `json:"thumbnail_tag"`
	PubDate        string  `json:"pub_date"`
	Title          string  `json:"title"`
	Items          []struct {
		ID              string   `json:"id"`
		NaturalKey      []string `json:"natural_key"`
		Filename        string   `json:"filename"`
		Href            string   `json:"href"`
		Order           int      `json:"order"`
		Fragment        string   `json:"fragment"`
		MediaType       string   `json:"media_type"`
		FullPath        string   `json:"full_path"`
		Depth           int      `json:"depth"`
		URL             string   `json:"url"`
		Label           string   `json:"label"`
		MinutesRequired float64  `json:"minutes_required"`
	} `json:"items"`
	Authors string `json:"authors"`
}

//BookChapter -
type BookChapter struct {
	Archive         string   `json:"archive"`
	Content         string   `json:"content"`
	URL             string   `json:"url"`
	NaturalKey      []string `json:"natural_key"`
	FullPath        string   `json:"full_path"`
	MinutesRequired float64  `json:"minutes_required"`
	NextChapter     struct {
		URL    string `json:"url"`
		Title  string `json:"title"`
		WebURL string `json:"web_url"`
	} `json:"next_chapter"`
	PreviousChapter struct {
		URL    string `json:"url"`
		Title  string `json:"title"`
		WebURL string `json:"web_url"`
	} `json:"previous_chapter"`
	Stylesheets []struct {
		FullPath    string `json:"full_path"`
		URL         string `json:"url"`
		OriginalURL string `json:"original_url"`
	} `json:"stylesheets"`
	Images               []string      `json:"images"`
	AssetBaseURL         string        `json:"asset_base_url"`
	WebURL               string        `json:"web_url"`
	LastPosition         interface{}   `json:"last_position"`
	Videoclips           []interface{} `json:"videoclips"`
	PublisherScripts     string        `json:"publisher_scripts"`
	PublisherScriptFiles []interface{} `json:"publisher_script_files"`
	AllowScripts         bool          `json:"allow_scripts"`
	Videoclip            interface{}   `json:"videoclip"`
	AcademicExcluded     bool          `json:"academic_excluded"`
	Subjects             []interface{} `json:"subjects"`
	Authors              []struct {
		Name string `json:"name"`
	} `json:"authors"`
	Cover            string        `json:"cover"`
	BookTitle        string        `json:"book_title"`
	Updated          time.Time     `json:"updated"`
	SiteStyles       []string      `json:"site_styles"`
	CreatedTime      time.Time     `json:"created_time"`
	LastModifiedTime time.Time     `json:"last_modified_time"`
	Filename         string        `json:"filename"`
	Path             string        `json:"path"`
	EpubProperties   []interface{} `json:"epub_properties"`
	HeadExtra        interface{}   `json:"head_extra"`
	HasVideo         bool          `json:"has_video"`
	Title            string        `json:"title"`
	Description      string        `json:"description"`
	VirtualPages     int           `json:"virtual_pages"`
}
