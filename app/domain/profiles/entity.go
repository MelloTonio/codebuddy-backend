package profiles

type Profile struct {
	Username    string   `json:"username" bson:"username"`
	Password    string   `json:"password" bson:"password"`
	Groups      []string `json:"groups" bson:"groups"`
	ProfileType string   `json:"profile_type" bson:"profile_type"`
}
