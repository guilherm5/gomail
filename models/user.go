package models

type User struct {
	IDUsuario int    `json:"id_usuario"`
	Nome      string `json:"nome"`
	Email     string `json:"email"`
	Senha     string `json:"senha"`
}
