package httpChallenge

type (
	Challenge struct {
		Name      string   `json:"name" bson:"name"`
		Text      string   `json:"text" bson:"text"`
		Input     string   `json:"input" bson:"input"`
		Output    string   `json:"output" bson:"output"`
		Group     string   `json:"group" bson:"group"`
		Difficult string   `json:"difficulty" bson:"difficulty"`
		Answer    []Answer `json:"answer" bson:"answer"`
	}

	ChallengeRespBody struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	Answer struct {
		AlumniName string `json:"name" bson:"name"`
		Text       string `json:"text" bson:"text"`
	}
)
