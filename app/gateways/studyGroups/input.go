package httpStudyGroup

type (
	StudyGroup struct {
		Name        string `json:"name" validate:"required"`
		Students    string `json:"students" validate:"required"`
		Subject     string `json:"subject" validate:"required"`
		Description string `json:"description" validate:"required"`
		Warning     string `json:"warning" `
	}

	StudyGroupRespBody struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}
)
