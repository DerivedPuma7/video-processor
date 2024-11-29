package domain_test

import (
	"testing"
	"time"

	"github.com/derivedpuma7/video-processor/domain"

	"github.com/stretchr/testify/require"
   uuid "github.com/satori/go.uuid"
)

func TestValidateIfVideoIsEmpty(t *testing.T) {
   video := domain.NewVideo()

   err := video.Validate()

   require.Error(t, err)
}

func TestVideoIdIsNotAnUuid(t *testing.T) {
   video := domain.NewVideo()
   video.ID = "any id"

   err := video.Validate()

   require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
   video := domain.NewVideo()
   video.ID = uuid.NewV4().String()
   video.ResourceID = "any resource id"
   video.FilePath = "any file path"
   video.CreatedAt = time.Now()

   err := video.Validate()

   require.Nil(t, err)
}
