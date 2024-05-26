package models

// Post ...
type Post struct {
	Id       string `json:"id"`
	Title    string `json:"title"`
	ImageUrl string `json:"image_url"`
	OwnerId  string `json:"owner_id"`
}

// Posts ...
type Posts struct {
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	ImageUrl   string  `json:"image_url"`
	Likes      int64   `json:"likes"`
	LikeOwners []*User `json:"like_owners"`
}