package controller

import (
	"net/http"
)

type voteController struct{}

type IVoteController interface {
	GetUpVote(response http.ResponseWriter, request *http.Request)
	AddUpVote(response http.ResponseWriter, request *http.Request)
	GetDownVote(response http.ResponseWriter, request *http.Request)
	AddDownVote(response http.ResponseWriter, request *http.Request)
	//need to put methods for removing upvote and downvote
}

func ObjIVoteController() IVoteController {
	return &voteController{}
}

func (*voteController) GetUpVote(response http.ResponseWriter, request *http.Request) {

}
func (*voteController) AddUpVote(response http.ResponseWriter, request *http.Request) {

}
func (*voteController) GetDownVote(response http.ResponseWriter, request *http.Request) {

}
func (*voteController) AddDownVote(response http.ResponseWriter, request *http.Request) {

}
