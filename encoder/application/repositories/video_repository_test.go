package repositories_test

import (
	"testing"

	"github.com/derivedpuma7/video-processor/application/repositories"
	"github.com/derivedpuma7/video-processor/domain"
	"github.com/derivedpuma7/video-processor/framework/database"
	"github.com/stretchr/testify/require"
)

func TestVideoInsert(t *testing.T) {
   db := database.NewDbTest()
   defer db.Close()
   video, _ := domain.NewVideo("any resource id", "any path")

   repo := repositories.NewVideoRepository(db)
   repo.Insert(video)
   v, err := repo.Find(video.ID)

   require.Nil(t, err)
   require.NotEmpty(t, v.ID)
   require.Equal(t, v.ID, video.ID)
}
