# API SIMPLES GOLANG
> API simples escrito em Golang utilizando o framework [Gin](https://github.com/gin-gonic/gin)

## 💻 Pré-requisitos
Antes de começar, verifique se você atendeu aos seguintes requisitos:
* Golang 1.16
* Um banco de dados postgresql

## ☕ Desenvolvimento

1 - Clonar o projeto do GitHub num diretório de sua preferência:
```shell
$ cd "diretorio de sua preferencia"
$ git clone https://github.com/ivanilsonaraujojr/api-simples-golang.git
```
2 - Abra a pasta do projeto na sua IDE

3 - Ter um banco de dados postgresql instalado em sua maquina(ou rodando em um container)
#### Comando para rodar um container docker de banco de dados postgresql na porta 5432 em localhost:
```shell
$ docker run -d --name apigo -e "POSTGRES_PASSWORD=postgres" -e "POSTGRES_USER=postgres" -e "POSTGRES_DB=apigo" -p 5432:5432 -v "${HOME}/apigo:/var/lib/postgresql/data" postgres:12-alpine
```
4 - Abra o terminal e digite: 
```shell 
$ go run api.go
```
__A aplicação ficará disponível na porta `8080`__