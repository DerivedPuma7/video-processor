package domain

import (
	"time"

	govalidator "github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Video struct {
	ID         string    `json:"encoded_video_folder" valid:"uuid" gorm:"type:uuid;primary_key"`
	ResourceID string    `json:"resource_id" valid:"notnull" gorm:"type:varchar(255)"`
	FilePath   string    `json:"file_path" valid:"notnull" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"-" valid:"-"`
	Jobs       []*Job    `json:"-" valid:"-" gorm:"ForeignKey:VideoID"`
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

func NewVideo(resourceID string, filePath string) (*Video, error) {
	video := &Video{
		ID:         uuid.NewV4().String(),
		ResourceID: resourceID,
		FilePath:   filePath,
		CreatedAt:  time.Now(),
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
