package dtos

type PassageiroDto struct {
	Nome  string `json: Nome`
	Cpf   string `json: CPF`
	Email string `json: Email`
	Senha string `json: Senha`
}
