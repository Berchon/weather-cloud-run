package dto

type ViaCepDto struct {
	ZipCode      string `json:"cep,omitempty"`
	Street       string `json:"logradouro,omitempty"`
	Neighborhood string `json:"bairro,omitempty"`
	City         string `json:"localidade,omitempty"`
	State        string `json:"uf,omitempty"`
	Error        string `json:"erro,omitempty"`
}
