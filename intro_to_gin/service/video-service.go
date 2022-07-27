package service

import "intro_to_gin_ep1/entity"

//create VideoService interface to export
type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

//inner service for private stuff
type videoService struct {
	videos []entity.Video
}

//allocate memory for VideoService
func New() VideoService {
	return &videoService{}
}
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}
func (service *videoService) Findall(video entity.Video) []entity.Video {
	return service.videos
}
