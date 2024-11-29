package domain_test

import (
	"testing"

	"github.com/derivedpuma7/video-processor/domain"
	"github.com/stretchr/testify/require"
)

func TestNewJob(t *testing.T) {
	video, _ := domain.NewVideo("any resource id", "any filepath")

	job, err := domain.NewJob("any output", "any status", video)

	require.Nil(t, err)
	require.NotNil(t, job)
}
