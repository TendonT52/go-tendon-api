package models

type BriefLearningNode struct {
	LearningNodeId          string `json:"learningNodeId" bson:"learning_node_id"`
	LearningNodeName        string
	LearningNodeDescription string   `json:"learningNodeDescription" bson:"learning_node_description"`
	NextLearningNode        []string `json:"nextLearningNode" bson:"next_learning_node"`
	PrevLearningNode        []string `json:"prevLearningNode" bson:"prev_learning_node"`
}

type LearningNode struct {
	LearningNodeId          string
	LearningNodeName        string
	LearningNodeDescription string
	Node                    []Node
	Curriculum              []BriefCurriculum
	NextLearningNode        []string
	PrevLearningNode        []string
}
