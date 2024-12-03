package services

import (
	"context"
   "io"
	"log"
	"os"

	"cloud.google.com/go/storage"
	"github.com/derivedpuma7/video-processor/application/repositories"
	"github.com/derivedpuma7/video-processor/domain"
)

type VideoService struct {
	Video *domain.Video
   VideoRepository repositories.VideoRepository
   localStoragePath string
}

func NewVideoService() VideoService {
   return VideoService{
      localStoragePath: os.Getenv("localStoragePath/"),
   }
}

func (v *VideoService) Download(bucketName string) error {
   ctx := context.Background()
   client, err := storage.NewClient(ctx)
   if err != nil {
      return err
   }

   bucket := client.Bucket(bucketName)
   obj := bucket.Object(v.Video.FilePath)
   reader, err := obj.NewReader(ctx)
   if err != nil {
      return err
   }
   defer reader.Close()

   body, err := io.ReadAll(reader)
   if err != nil {
      return err
   }

   file, err := os.Create(v.localStoragePath + v.Video.ID + ".mp4")
   if err != nil {
      return err
   }
   defer file.Close()

   _, err = file.Write(body)
   if err != nil {
      return err
   }

   log.Printf("video %v has been downloaded", v.Video.ID)
   return nil
}
