package controllers

import (
	"Golang_udemy/todo_app/app/models"
	"log"
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		generateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

//インデックス画面に遷移する
func index(w http.ResponseWriter, r *http.Request) {
	//セッションを取得
	sess, err := session(w, r)
	if err != nil {
		//取得できなければTOPに遷移
		http.Redirect(w, r, "/", 302)
	} else {
		//できればそこからユーザー情報を取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//取得したユーザーのTODO情報を取得して画面遷移
		todos, _ := user.GetTodosByUser()
		user.Todos = todos

		generateHTML(w, user, "layout", "private_navbar", "index")
	}

}

//TODO追加画面に遷移
func todoNew(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		//セッションがなければログイン画面へ
		http.Redirect(w, r, "/login", 302)
	} else {
		//セッションがあれば追加画面へ
		generateHTML(w, nil, "layout", "private_navbar", "todo_new")
	}
}

//記入したTODOを保存する
func todoSave(w http.ResponseWriter, r *http.Request) {
	//セッションを取得
	sess, err := session(w, r)
	if err != nil {
		//なければログイン画面へ遷移
		http.Redirect(w, r, "/login", 302)
	} else {
		//あればTODOを追加する
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//セッション情報からユーザーを取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//ユーザーのTODOに新しいTODOを追加
		content := r.PostFormValue("content")
		if err := user.CreateTodo(content); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)
	}

}

//TODO編集画面へ遷移
func todoEdit(w http.ResponseWriter, r *http.Request, id int) {
	//セッションを取得
	sess, err := session(w, r)
	if err != nil {
		//なければログイン画面へ
		http.Redirect(w, r, "/login", 302)
	} else {
		//あればユーザー情報を確認
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//あればIDからTODO情報を取得し、編集画面へ遷移
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		generateHTML(w, t, "layout", "private_navbar", "todo_edit")
	}

}

//TODOを更新する
func todoUpdate(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		//セッションがなければログイン画面へ
		http.Redirect(w, r, "/login", 302)
	} else {
		//あればフォームから値を取得
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		//セッションからユーザー情報を取得
		user, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		content := r.PostFormValue("content")
		//得た情報を用いてTODOをアップデート
		t := &models.Todo{ID: id, Content: content, UserID: user.ID}
		if err := t.UpdateTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)

	}

}

//TODOを削除する
func todoDelete(w http.ResponseWriter, r *http.Request, id int) {
	sess, err := session(w, r)
	if err != nil {
		//セッションがなければログイン画面へ
		http.Redirect(w, r, "/login", 302)
	} else {
		_, err := sess.GetUserBySession()
		if err != nil {
			log.Println(err)
		}
		//引数のIDに等しいTODOをDBから削除する
		t, err := models.GetTodo(id)
		if err != nil {
			log.Println(err)
		}
		if err := t.DeleteTodo(); err != nil {
			log.Println(err)
		}
		http.Redirect(w, r, "/todos", 302)

	}

}
