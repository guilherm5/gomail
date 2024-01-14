<h1 align="center">Envio de Email com golang</h1>
<p align="center">
  <img src="img/go.png" alt="Mascote do golang entregando uma carta na caixinha de correio">
</p>

<h2 align="center">Tecnologias usadas</h2>

> GO 1.20 

> Gin Framework

> LIBPQ 

> JWT 

> Postgresql

> PLPGSQL

> Render (Para deploy)

> S3 Bucket (Para armazenar os arquivos enviados)


<h2 align="center">Instalação</h2>

```
Ter o Go instalado em sua maquina 
Link: https://go.dev/doc/install

Rodar este comando no terminal (dentro da pasta raiz do projeto):
go mod tidy

Configurar arquivo example.env com as informações do seu banco de dados e email
```

<h2 align="center">Funcionalidades</h2>

> Cadastrar usuario

> Logar usuario 

> Crud completo de usurio 

> Enviar email com ou sem arquivos

> Ver emails enviados

> Ver emails que certo usuario enviou (pesquisa pelo nome)

<h2 align="center">Sobre o projeto</h2>
<p>Este projeto nasceu da necessidade de enviar emails, eu não queria entrar toda hora no google e acessar o gmail e poder enviar emails, então, criei minha própria aplicação para enviar emails. 

O projeto segue um fluxo interessante, onde eu posso cadastrar um usuário(que pode ser do tipo user ou admin). Após o cadastro do usuário eu posso me autenticar(usando a rota de login que me devolve um JWT para me autenticar nas rotas seguintes), após autenticado, posso finalmente enviar email.

Também tenho rotas para acessar os emails enviados por mim, rotas para atualizar meu usuário, deletar meu usuário e ver informações sobre meu usuáio. Todas estas infomações são armazenadas em um banco de dados Postgresql(que esta na nuvem [plataforma neondb]).

O projeto possui um grau de "autorização" de acesso em certas rotas, caso voce for do tipo "admin", você consegue acessar rotas extras, rotas onde você pode ver emails enviados de TODOS usuários do sistema, deletar QUALQUER usuário do sistema, dar upload em QUALQUER usuário do sistema e fazer todas as outras operações restantes do sitema. Mas caso você seja do tipo "user", você só pode acessar rotas que enviam email e que realizam CRUD em você mesmo.
</p>

<h2 align="center">Rotas para user</h2>
<h4 align="center">Todos os usuarios devem estar autenticados com um token que obtemos na rota de login</h4>

```
Cadastro -> /api/user

Rota usada para cadastrar um usuario, aceitando os campos "nome, email, senha e tipo_usuario" 
```

```
Login -> /api/login

Rota usada para logar seu usuario, aceitando os campos "email e senha
```

```
Login -> /api/mail

Rota usada para enviar email aceitando os campos: destinatario, assunto, conteudo(o remetende sera padrão, deve ser preenchido na roda example.env)
```

```
Ver emails enviados -> /api/mail-user

Rota usada para verificar os emails que você enviou, voce recebera um JSON como resposta, com os campos: email, assunto, conteudo
```

```
Pegar informações do seu usuario -> /api/my-user

Rota usada para verificar as informações do usuario, você recebera um JSON contendo com os campos: nome, email, id_usuario, tipo_usuario
```

```
Atualizar o nome do seu usuario -> /api/update-name-my-user

Rota usada para atualizar o nome do seu usuario
```

```
Atualizar senha do seu usuario -> /api/update-secret-my-user

Rota usada para atualizar a senha do seu usuario
```

```
Deleter seu usuario -> /api/delete-my-user

Rota usada para deletar seu usuario
```

```
Enviar email com arquivo -> /api/file-mail

Rota usada para enviar email contendo arquivos
```

```
Ver emails recebidos(rota em desenvolvimento) -> /api/mail-received

Rota usada para ver emails que foram recebidos
```

<h2 align="center">Rotas para admin</h2>
<h4 align="center">Todos os usuarios devem estar autenticados com um token que obtemos na rota de login</h4>

```
O admin pode realizar operações em todas as rotas acima, mas além das rotas acima ele pode acessar as seguintes rotas:
```

```
Ver emails enviados por todos usuario -> /api/mails

Rota usada para ver emails que foram enviados
```

```
Ver todos usuarios do sistema -> /api/users

Rota para ver todos os usuarios armazenados no sistema, recebendo um JSON com os campos: nome, email, id_usuario, tipo_usuario
```

```
Deletar usuario -> api/delete-user

Rota para deletar algum usuario, passando um JSON contendo o ID do usuario a excluir
```

```
Atualizar usuario(qualquer usuario)-> /api/atualizar-user

Rota para atualizar algum usuario, passando um JSON contendo os campos: nome, email, senha, tipo_usuario, id_usuario
```

<h2 align="center">Collection para o Postman esta no arquivo: <strong>Gomail.json</stron></h2>

