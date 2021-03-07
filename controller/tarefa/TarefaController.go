package tarefa

import (
	"apigo/database"
	"apigo/models"
	"time"
)

func GetTarefas() ([]models.Tarefa, error) {
	db := database.OpenConnection()
	const query = `SELECT id, usuario, tarefa, datacriacao, concluida FROM tarefas ORDER BY dataCriacao DESC`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	results := make([]models.Tarefa, 0)

	for rows.Next() {
		var id int
		var usuario, tarefa string
		var dataCriacao time.Time
		var concluida bool

		err = rows.Scan(&id, &usuario, &tarefa, &dataCriacao, &concluida)
		if err != nil {
			return nil, err
		}

		results = append(results, models.Tarefa{Usuario: usuario, Tarefa: tarefa, DataCriacao: dataCriacao, Concluida: concluida})
	}

	defer db.Close()
	defer rows.Close()

	return results, nil
}

func AddTarefa(tarefa models.Tarefa) error {
	db := database.OpenConnection()
	const query = `INSERT INTO tarefas(usuario, tarefa, datacriacao, concluida) VALUES ($1, $2, $3, $4)`
	_, err := db.Exec(query, tarefa.Usuario, tarefa.Tarefa, tarefa.DataCriacao, tarefa.Concluida)
	defer db.Close()
	return err
}
