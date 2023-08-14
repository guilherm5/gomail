package models

type Mail struct {
	IDMail       int    `json:"id_mail"`
	Conteudo     string `json:"conteudo"`
	Assunto      string `json:"assunto"`
	Destinatario string `json:"destinatario"`
	Remetente    string `json:"remetente"`
	IDUsuario    int    `json:"id_usuario"`
	JsonData     string `json:"json_data"`
}
