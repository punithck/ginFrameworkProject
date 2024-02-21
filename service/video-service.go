package service

import "ginFrameworkProject/entity"

// It is not tightly related to http service. It can be used by any service. Think of it as db layer
// It exposes interfaces, Save and FindAll function

type VideoService interface {
	Save(video entity.Video) entity.Video
	FindAll() []entity.Video
}

// In golang if any struct implement the functions defined by the interfaces.
// Then we can inheritance concept. i.e representing struct from interface.
type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (v *videoService) Save(video entity.Video) entity.Video {
	v.videos = append(v.videos, video)
	return video
}

func (v *videoService) FindAll() []entity.Video {
	return v.videos
}
