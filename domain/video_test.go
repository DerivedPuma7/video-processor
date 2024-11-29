package domain_test

import (
	"testing"

	"github.com/derivedpuma7/video-processor/domain"
	"github.com/stretchr/testify/require"
)

func TestVideoIdIsNotAnUuid(t *testing.T) {
   video := domain.NewVideo("any resource id", "any path")
   video.ID = "any id"

   err := video.Validate()

   require.Error(t, err)
}

func TestVideoResourceIdIsNotNull(t *testing.T) {
   video := domain.NewVideo("", "any path")

   err := video.Validate()

   require.Error(t, err)
}

func TestVideoFilePathIsNotNull(t *testing.T) {
   video := domain.NewVideo("any resource id", "")

   err := video.Validate()

   require.Error(t, err)
}

func TestVideoValidation(t *testing.T) {
   video := domain.NewVideo("any resource id", "any path")

   err := video.Validate()

   require.Nil(t, err)
}
