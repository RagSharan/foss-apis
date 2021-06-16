package router

import (
	"github.com/ragsharan/foss-apis/controller"
)

type mapping struct{}

type IMapping interface {
	UriMappings()
}

var (
	httpRouter IRouter = NewMuxRouter()

	postController    controller.IPostController    = controller.ObjIPostController()
	commentController controller.ICommentController = controller.ObjICommentController()
	voteController    controller.IVoteController    = controller.ObjIVoteController()
	profileController controller.IProfileController = controller.ObjIProfileController()
	authController    controller.IAuthController    = controller.ObjIAuthController()
)

func InstMapping() IMapping {
	return mapping{}
}
func (mapping) UriMappings() {

	postMapping()
	commentMapping()
	voteMapping()
	profileMapping()

}

func postMapping() {
	httpRouter.GET("/post", postController.GetPost)
	httpRouter.GET("/postList", postController.GetPostList)
	httpRouter.POST("/post", postController.AddPost)
	httpRouter.POST("/postList", postController.AddPostList)
	httpRouter.PUT("/post", postController.UpdatePost)
	httpRouter.PUT("/postList", postController.UpdatePostList)
	httpRouter.DELETE("/post", postController.RemovePost)
}

func commentMapping() {
	httpRouter.GET("/comment", commentController.GetComment)
	httpRouter.POST("/comment", commentController.AddComment)
	httpRouter.DELETE("/comment", commentController.RemoveComment)
}

func voteMapping() {
	httpRouter.GET("/vote", voteController.GetUpVote)
	httpRouter.POST("/vote", voteController.AddUpVote)
	httpRouter.DELETE("/vote", voteController.GetDownVote)
	httpRouter.DELETE("/vote", voteController.AddDownVote)
}
func profileMapping() {
	httpRouter.GET("/profile", profileController.GetProfile)
	httpRouter.POST("/profile", profileController.AddProfile)
	httpRouter.PUT("/profile", profileController.UpdateProfile)
	httpRouter.PUT("/profile/pic", profileController.ChangePic)
	httpRouter.PUT("/profile/cover", profileController.UpdateCover)
	httpRouter.DELETE("/profile", profileController.RemoveProfile)
}
