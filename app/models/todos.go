package models

import (
	"log"
	"time"
)

type Todo struct {
	ID        int
	Content   string
	UserID    int
	CreatedAt time.Time
}

//TODOテーブルに情報を追加
func (u *User) CreateTodo(content string) (err error) {
	cmd := `INSERT INTO todos (
		content,
		user_id,
		created_at) VALUES ($1, $2, $3)`

	_, err = Db.Exec(cmd, content, u.ID, time.Now())
	if err != nil {
		log.Fatalln(err)
	}
	return err
}

//IDから対応するTODOを取得する
func GetTodo(id int) (todo Todo, err error) {
	todo = Todo{}
	cmd := `SELECT id, content, user_id, created_at FROM todos
	WHERE id = $1S`
	err = Db.QueryRow(cmd, id).Scan(
		&todo.ID,
		&todo.Content,
		&todo.UserID,
		&todo.CreatedAt)

	return todo, err

}

func GetTodos() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at from todos`
	rows, err := Db.Query(cmd)
	if err != nil {
		log.Fatalln(err)
	}

	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err

}

func (u *User) GetTodosByUser() (todos []Todo, err error) {
	cmd := `SELECT id, content, user_id, created_at FROM todos
	WHERE user_id = $1`

	rows, err := Db.Query(cmd, u.ID)
	if err != nil {
		log.Fatalln(err)
	}
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID,
			&todo.Content,
			&todo.UserID,
			&todo.CreatedAt)
		if err != nil {
			log.Fatalln(err)
		}
		todos = append(todos, todo)
	}
	rows.Close()

	return todos, err
}

//TODOの内容をアップデートする
func (t *Todo) UpdateTodo() (err error) {
	cmd := `UPDATE todos SET content = $1, user_id = $2
	WHERE id = $3`
	_, err = Db.Exec(cmd, t.Content, t.UserID, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err

}

//TODOを削除する
func (t *Todo) DeleteTodo() (err error) {
	cmd := `DELETE FROM todos WHERE id = $1`
	_, err = Db.Exec(cmd, t.ID)
	if err != nil {
		log.Fatalln(err)
	}
	return err
}
