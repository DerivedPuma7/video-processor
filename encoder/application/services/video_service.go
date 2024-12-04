package services

import (
	"context"
	"io"
	"log"
	"os"
	"os/exec"

	"cloud.google.com/go/storage"
	"github.com/derivedpuma7/video-processor/application/repositories"
	"github.com/derivedpuma7/video-processor/domain"
)


type VideoService struct {
	Video *domain.Video
   VideoRepository repositories.VideoRepository
   localStoragePath string
   bucket string
}

func NewVideoService(video *domain.Video, videoRepository repositories.VideoRepository, bucket string) VideoService {
   return VideoService{
      Video: video,
      VideoRepository: videoRepository,
      bucket: bucket,
      localStoragePath: os.Getenv("localStoragePath") + "/",
   }
}

func (v *VideoService) Download() error {
   ctx := context.Background()
   client, err := storage.NewClient(ctx)
   if err != nil {
      return err
   }

   bucket := client.Bucket(v.bucket)
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

func (v *VideoService) Fragment() error {
   err := os.Mkdir(v.localStoragePath + v.Video.ID, os.ModePerm)
   if err != nil {
      return err
   }

   source := v.localStoragePath + v.Video.ID + ".mp4"
   target := v.localStoragePath + v.Video.ID + ".frag"

   cmd := exec.Command("mp4fragment", source, target)
   output, err := cmd.CombinedOutput()
   if err != nil {
      return err
   }
   printOutput(output)
   return nil
}

func printOutput(out []byte) {
   if len(out) > 0 {
      log.Printf("=====> Output: %s\n", string(out))
   }
}
