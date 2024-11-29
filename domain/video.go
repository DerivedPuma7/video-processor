package domain

import (
	"time"

	govalidator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Video struct {
	ID         string    `valid:"uuid"`
	ResourceID string    `valid:"notnull"`
	FilePath   string    `valid:"notnull"`
	CreatedAt  time.Time `valid:"-"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo(resourceID string, filePath string) (*Video, error) {
	video := &Video{
      ID: uuid.NewV4().String(),
      ResourceID: resourceID,
      FilePath: filePath,
      CreatedAt: time.Now(),
   }
   err := video.Validate()
   if err != nil {
      return nil, err
   }
   return video, nil
}

func (v *Video) Validate() error {
	_, err := govalidator.ValidateStruct(v)
	return err
}
