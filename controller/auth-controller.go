package controller

import (
	"net/http"
)

type authController struct{}
type IAuthController interface {
	authenticate(response http.ResponseWriter, request *http.Request)
	register(response http.ResponseWriter, request *http.Request)
	verifyAccount(response http.ResponseWriter, request *http.Request)
	ResetPassword(response http.ResponseWriter, request *http.Request)
}

func ObjIAuthController() IAuthController {
	return &authController{}
}

func (*authController) authenticate(response http.ResponseWriter, request *http.Request) {

}
func (*authController) register(response http.ResponseWriter, request *http.Request) {

}
func (*authController) verifyAccount(response http.ResponseWriter, request *http.Request) {

}
func (*authController) ResetPassword(response http.ResponseWriter, request *http.Request) {

}

func setResponseData(resp http.ResponseWriter) {
	resp.Header().Set("Content-Type", "application/json")

}
