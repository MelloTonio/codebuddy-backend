package access

type Credential struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

type Description struct {
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	AccountID string `json:"account_id"`
}
