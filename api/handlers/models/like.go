package models

// Stuatus ...
type Status struct {
	Liked bool `json:"liked"`
}

// PostLikeOwners ...
type PostLikeOwners struct {
	UserId string `json:"user_id"`
	PostId string `json:"post_id"`
}

// CommentLikeOwners
type CommentLikeOwners struct {
	UserId    string `json:"user_id"`
	CommentId string `json:"comment_id"`
}
