package repositories_test

import (
	"testing"

	"github.com/derivedpuma7/video-processor/application/repositories"
	"github.com/derivedpuma7/video-processor/domain"
	"github.com/derivedpuma7/video-processor/framework/database"
	"github.com/stretchr/testify/require"
)

func TestJobInsert(t *testing.T) {
   db := database.NewDbTest()
   defer db.Close()
   video, _ := domain.NewVideo("any resource id", "any path")
   videoRepo := repositories.NewVideoRepository(db)
   videoRepo.Insert(video)

   job, _ := domain.NewJob("any output", "any status", video)
   jobRepo := repositories.NewJobRepository(db)
   jobRepo.Insert(job)
   j, err  :=  jobRepo.Find(job.ID)

   require.Nil(t, err)
   require.NotEmpty(t, j.ID)
   require.Equal(t, j.ID, job.ID)
   require.Equal(t, j.VideoID, video.ID)
}

func TestJobUpdate(t *testing.T) {
   db := database.NewDbTest()
   defer db.Close()
   video, _ := domain.NewVideo("any resource id", "any path")
   videoRepo := repositories.NewVideoRepository(db)
   videoRepo.Insert(video)

   job, _ := domain.NewJob("any output", "any status", video)
   jobRepo := repositories.NewJobRepository(db)
   jobRepo.Insert(job)
   job.Status = "some other status"
   jobRepo.Update(job)
   j, err  :=  jobRepo.Find(job.ID)

   require.Nil(t, err)
   require.NotEmpty(t, j.ID)
   require.Equal(t, j.Status, "some other status")
}
