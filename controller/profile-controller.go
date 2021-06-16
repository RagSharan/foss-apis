package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ragsharan/foss-apis/service"
)

type profileController struct{}
type IProfileController interface {
	GetProfile(res http.ResponseWriter, req *http.Request)
	AddProfile(res http.ResponseWriter, req *http.Request)
	UpdateProfile(res http.ResponseWriter, req *http.Request)
	ChangePic(res http.ResponseWriter, req *http.Request)
	UpdateCover(res http.ResponseWriter, req *http.Request)
	RemoveProfile(res http.ResponseWriter, req *http.Request)
}

func ObjIProfileController() IProfileController {
	return &profileController{}
}

var (
	profileService service.IProfileService = service.InstProfileService()
)

func (*profileController) GetProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	result, err := profileService.GetProfile(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err)
	} else {
		json.NewEncoder(res).Encode(result)
	}
}
func (*profileController) AddProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	result, err := profileService.AddProfile(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err)
	} else {
		json.NewEncoder(res).Encode(result)
	}
}
func (*profileController) UpdateProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	result, err := profileService.UpdateProfile(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err)
	} else {
		json.NewEncoder(res).Encode(result)
	}

}
func (*profileController) RemoveProfile(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	result, err := profileService.RemoveProfile(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err)
	} else {
		json.NewEncoder(res).Encode(result)
	}
}
func (*profileController) ChangePic(res http.ResponseWriter, req *http.Request) {

}
func (*profileController) UpdateCover(res http.ResponseWriter, req *http.Request) {

}
