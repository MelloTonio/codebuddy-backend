package challenges

type Challenge struct {
	Name      string   `json:"name" bson:"name"`
	Group     string   `json:"group" bson:"group"`
	Text      string   `json:"text" bson:"text"`
	Input     string   `json:"input" bson:"input"`
	Output    string   `json:"output" bson:"output"`
	Difficult string   `json:"difficulty" bson:"difficulty"`
	Answer    []Answer `json:"answer" bson:"answer"`
}

type Answer struct {
	AlumniName string `json:"name" bson:"name"`
	Text       string `json:"text" bson:"text"`
}
