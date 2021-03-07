package main

import (
	"apigo/controller/tarefa"
	"apigo/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const version = "0.1.0"

func main() {

	r := gin.Default()

	r.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"Versão": version})
	})

	r.GET("/tarefa", func(context *gin.Context) {
		tarefas, err := tarefa.GetTarefas()
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
			return
		}

		context.JSON(http.StatusOK, tarefas)
	})

	r.POST("/tarefa", func(context *gin.Context) {
		var taref models.Tarefa

		if err := context.BindJSON(&taref); err != nil {
			context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": http.StatusBadRequest, "msg": err.Error()})
			return
		}

		taref.DataCriacao = time.Now()
		taref.Concluida = false

		if err := tarefa.AddTarefa(taref); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"status": "internal error: " + err.Error()})
		}
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}

}
