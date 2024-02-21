package controller

import (
	"ginFrameworkProject/entity"
	"ginFrameworkProject/service"
	"ginFrameworkProject/validators"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

// Why do we use interface instead of struct?
// I think this is the style he used. I see an advantage. Who ever wants to use the service,
// they must call the function. They cannot use the data.

// This is the handler. It can know the terminology of gen libraries(gin.Context, gin.Bind)
// This knows how to handle the request
type VideoController interface {
	// Functionalities supported by gin
	Save(ctx *gin.Context) (*entity.Video, error)
	FindAll() []entity.Video
	ShowAll(ctx *gin.Context)
}

type videoController struct {
	// It is a kind of db connection...
	videoservice service.VideoService
}

func (v *videoController) ShowAll(ctx *gin.Context) {
	videos := v.videoservice.FindAll()
	data := gin.H{
		"title":  "Video view",
		"videos": videos,
	}
	// You created the data and now you are passing it
	ctx.HTML(http.StatusOK, "index.html", data)
}

// It is used to extract data from the context and save it in videoservice
func (v *videoController) Save(ctx *gin.Context) (*entity.Video, error) {
	var video entity.Video
	// You are passing the object to which raw request to be loaded.
	// We can bindjson, bindheader,BindXML etc
	//ctx.BindJSON(&video)
	// Here they are using golang struct tags to validate the input data. For that we can use differrent function
	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return nil, err
	}
	// this is not the part of the framework. This is the part of golang validator package.
	err1 := validate.Struct(video)
	if err1 != nil {
		return nil, err1
	}
	// load data
	v.videoservice.Save(video)
	return &video, nil
}

func (v *videoController) FindAll() []entity.Video {
	return v.videoservice.FindAll()
}

var validate *validator.Validate

func New(videoService service.VideoService) VideoController {
	// Adding custom validation is not part of framework but package that we use
	validate = validator.New()
	validate.RegisterValidation("is-cool", validators.ValidateCoolTitel)
	return &videoController{videoservice: videoService}
}
