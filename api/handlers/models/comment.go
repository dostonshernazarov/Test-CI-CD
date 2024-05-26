package models

// Comment ...
type Comment struct {
	Id        string `json:"id"`
	Content   string `json:"content"`
	CreatedAt string `json:"createdat"`
	UpdatedAt string `json:"updatedat"`
	DeletedAt string `json:"deletedat"`
	PostId    string `json:"post_id"`
	OwnerId   string `json:"owner_id"`
}

// Comments ...
type Comments struct {
	Id         string  `json:"id"`
	Content    string  `json:"content"`
	CreatedAt  string  `json:"createdat"`
	UpdatedAt  string  `json:"updatedat"`
	DeletedAt  string  `json:"deletedat"`
	PostId     string  `json:"post_id"`
	OwnerId    string  `json:"owner_id"`
	Likes      int64   `json:"likes"`
	LikeOwners []*User `json:"like_owners"`
}
