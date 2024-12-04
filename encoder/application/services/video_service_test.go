package services_test

import (
	"log"
	"testing"

	"github.com/derivedpuma7/video-processor/application/repositories"
	"github.com/derivedpuma7/video-processor/application/services"
	"github.com/derivedpuma7/video-processor/domain"
	"github.com/derivedpuma7/video-processor/framework/database"
	"github.com/stretchr/testify/require"

	"github.com/joho/godotenv"
)

func init() {
   err := godotenv.Load("../../.env")
   if err != nil {
      log.Fatalf("Error loading .env file")
   }
}

func prepare() (*domain.Video, repositories.VideoRepository) {
   video, _ := domain.NewVideo("any resource id", "convite.mp4")
   db := database.NewDbTest()
   defer db.Close()
   repo := repositories.VideoRepositoryDb{Db: db}
   return video, repo
}

func TestDownload(t *testing.T) {
   video, repo := prepare()
   videoService := services.NewVideoService(video, repo, "golang-encoder")

   err := videoService.Download()

   require.Nil(t, err)
}
