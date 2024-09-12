package model

type Response struct {
	Message string `json:"message"`
}

func (Response) MensagemErroId(erro string) {

	respota := Response{
		Message: "Id não pode ser nulo",
	}
	erro = respota.Message
}
func (Response) MensagemErroNoBanco(erro string) {

	respota := Response{
		Message: "Erro no banco tente novamente.",
	}
	erro = respota.Message
}
