package controller

import (
	"net/http"
)

type commentController struct{}
type ICommentController interface {
	GetComment(response http.ResponseWriter, request *http.Request)
	AddComment(response http.ResponseWriter, request *http.Request)
	RemoveComment(response http.ResponseWriter, request *http.Request)
}

func ObjICommentController() ICommentController {
	return &commentController{}
}

func (*commentController) GetComment(response http.ResponseWriter, request *http.Request) {

}
func (*commentController) AddComment(response http.ResponseWriter, request *http.Request) {

}
func (*commentController) RemoveComment(response http.ResponseWriter, request *http.Request) {

}
