package entity

type Vote struct {
	PostId   string `json:"postId" bson:"postId"`
	UpVote   int    `json:"upvote" bson:"upVote"`
	DownVote int    `json:"downvote" bson:"downVote"`
}
