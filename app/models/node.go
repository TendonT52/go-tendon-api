package models

import (
	"time"
)

const (
	Required  = "required"
	Extension = "extension"
	Optional  = "optional"
)

const (
	Text  = "text"
	Image = "image"
	Video = "video"
	Pdf   = "pdf"
	Sound = "sound"
)

type Node struct {
	Type      string
	NodeId    string
	Data      string
	Priority  string
	CreatedAt time.Time `json:"createdAt" bson:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" bson:"updated_at"`
}
