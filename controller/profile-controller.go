package controller

import (
	"net/http"
)

type profileController struct{}
type IProfileController interface {
	GetProfile(response http.ResponseWriter, request *http.Request)
	UpdateProfile(response http.ResponseWriter, request *http.Request)
	ChangePic(response http.ResponseWriter, request *http.Request)
	UpdateCover(response http.ResponseWriter, request *http.Request)
	//	RemoveProfile(response http.ResponseWriter, request *http.Request)
}

func ObjIProfileController() IProfileController {
	return &profileController{}
}

func (*profileController) GetProfile(response http.ResponseWriter, request *http.Request) {

}
func (*profileController) UpdateProfile(response http.ResponseWriter, request *http.Request) {

}
func (*profileController) ChangePic(response http.ResponseWriter, request *http.Request) {

}
func (*profileController) UpdateCover(response http.ResponseWriter, request *http.Request) {

}
