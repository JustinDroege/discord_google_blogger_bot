package models

type Posts struct {
	Kind          string `json:"kind"`
	NextPageToken string `json:"nextPageToken"`
	Items         []Post `json:"items"`
	Etag          string `json:"etag"`
}

type Post struct {
	Kind      string   `json:"kind"`
	Id        string   `json:"id"`
	Published string   `json:"published"`
	Updated   string   `json:"updated"`
	Url       string   `json:"url"`
	SelfLink  string   `json:"selfLink"`
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Author    Author   `json:"author"`
	Replies   Replies  `json:"replies"`
	Labels    []string `json:"labels"`
	Location  Location `json:"location"`
	Etag      string   `json:"etag"`
}

type Author struct {
	Id          string `json:"id"`
	DisplayName string `json:"displayName"`
	Image       Image  `json:"image"`
	URL         string `json:"url"`
}

type Image struct {
	Image string `json:"image"`
}

type Replies struct {
	SelfLink   string `json:"selfLink"`
	TotalItems string `json:"totalItems"`
}

type Location struct {
	Name      string  `json:"name"`
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
	Span      string  `json:"span"`
}
