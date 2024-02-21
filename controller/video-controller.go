package controller

import (
	"ginFrameworkProject/entity"
	"ginFrameworkProject/service"
	"github.com/gin-gonic/gin"
)

// This is the handler. It can know the terminology of gen libraries(gin.Context, gin.Bind)
// This knows how to handle the request
type VideoController interface {
	// Functionalities supported by gin
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type videoController struct {
	// It is a kind of db connection...
	videoservice service.VideoService
}

// It is used to extract data from the context and save it in videoservice
func (v *videoController) Save(ctx *gin.Context) entity.Video {
	var video entity.Video
	// You are passing the object to which raw request to be loaded.
	ctx.BindJSON(&video)
	// load data
	v.videoservice.Save(video)
	return video
}

func (v *videoController) FindAll() []entity.Video {
	return v.videoservice.FindAll()
}

func New(videoService service.VideoService) VideoController {
	return &videoController{videoservice: videoService}
}
