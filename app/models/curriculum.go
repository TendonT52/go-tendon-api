package models

import (
	"time"
)

const (
	PublicAccess    = "public"
	ProtectedAccess = "protected"
	PrivateAccess   = "private"
)

type Curriculum struct {
	CurriculumId          string `json:"id" bson:"_id,omitempty"`
	CurriculumName        string `json:"curriculumName " bson:"curriculum_name"`
	CurriculumDescription string
	BriefLearningNode     []BriefLearningNode `json:"BriefLearningNode" bson:"BriefLearningNode"`
	Access                string              `json:"access" bson:"access"`
	CreateBy              string              `json:"createBy" bson:"create_by"`
	CreatedAt             time.Time           `json:"createdAt" bson:"created_at"`
	UpdatedAt             time.Time           `json:"updatedAt" bson:"updated_at"`
}

type BriefCurriculum struct {
	CurriculumId          string `json:"id" bson:"_id,omitempty"`
	CurriculumName        string `json:"curriculumName " bson:"curriculum_name"`
	CurriculumDescription string
	Access                string `json:"access" bson:"access"`
}
