package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/ragsharan/foss-apis/service"
)

type postController struct{}
type IPostController interface {
	GetPost(res http.ResponseWriter, req *http.Request)
	GetPostList(res http.ResponseWriter, req *http.Request)
	AddPost(res http.ResponseWriter, req *http.Request)
	AddPostList(res http.ResponseWriter, req *http.Request)
	RemovePost(res http.ResponseWriter, req *http.Request)
	UpdatePost(res http.ResponseWriter, req *http.Request)
	UpdatePostList(res http.ResponseWriter, req *http.Request)
}

func ObjIPostController() IPostController {
	return &postController{}
}

var (
	postService service.IPostService = service.InstPostService()
)

func (*postController) GetPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	post, err := postService.FindPost(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}
func (*postController) GetPostList(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Println(err)
	}
	post, err := postService.FindPostList(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}
func (*postController) AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	post, err := postService.CreatePost(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}

func (*postController) AddPostList(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	post, err := postService.CreatePostList(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}

func (*postController) RemovePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	post, err := postService.DeletePost(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}
func (*postController) UpdatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	post, err := postService.UpdatePost(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}
func (*postController) UpdatePostList(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	data, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatal(err)
	}
	post, err := postService.UpdatePostList(data)
	if err != nil {
		log.Println(err)
		json.NewEncoder(res).Encode(err.Error())
	} else {
		json.NewEncoder(res).Encode(post)
	}
}
