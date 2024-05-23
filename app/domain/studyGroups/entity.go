package studygroups

type StudyGroup struct {
	Name        string   `json:"name" bson:"name"`
	Subject     string   `json:"subject" bson:"subject"`
	Students    []string `json:"students" bson:"students"`
	Description string   `json:"description" bson:"description"`
	Warnings    []string `json:"warnings" bson:"warnings"`
}
