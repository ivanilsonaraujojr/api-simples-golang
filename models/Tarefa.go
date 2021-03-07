package models

import "time"

type Tarefa struct {
	Usuario     string    `json:"usuario" binding:"required"`
	Tarefa      string    `json:"tarefa" binding:"required"`
	DataCriacao time.Time `json:"dataCriacao"`
	Concluida   bool      `json:"concluida"`
}
