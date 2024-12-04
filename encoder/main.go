package main

import (
	"fmt"

	"github.com/derivedpuma7/video-processor/application/repositories"
	"github.com/derivedpuma7/video-processor/application/services"
	"github.com/derivedpuma7/video-processor/domain"
	"github.com/derivedpuma7/video-processor/framework/database"
	log "github.com/sirupsen/logrus"
)

func prepare() (*domain.Video, repositories.VideoRepository) {
   video, _ := domain.NewVideo("any resource id", "convite.mp4")
   db := database.NewDbTest()
   defer db.Close()
   repo := repositories.VideoRepositoryDb{Db: db}
   return video, repo
}

func main() {
	log.Info("ola mundo")
   video, repo := prepare()
   videoService := services.NewVideoService(video, repo, "golang-encoder")

   err := videoService.Download()
   fmt.Println(err)
}
