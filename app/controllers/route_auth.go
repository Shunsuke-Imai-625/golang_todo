package controllers

import (
	"Golang_udemy/todo_app/app/models"
	"log"
	"net/http"
)

//サインアップ画面に遷移する
func signup(w http.ResponseWriter, r *http.Request) {
	//GETで飛んできたときは、セッションの有無で作業を振り分ける
	if r.Method == "GET" {
		//セッションがあれば、サインアップ画面へ
		_, err := session(w, r)
		if err != nil {
			generateHTML(w, nil, "layout", "public_navbar", "signup")
			//なければ、TODOにリダイレクト
		} else {
			http.Redirect(w, r, "/todos", 302)
		}
		//POSTで飛んできたときは、サインアップ処理を行う
	} else if r.Method == "POST" {
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		user := models.User{
			Name:     r.PostFormValue("name"),
			Email:    r.PostFormValue("email"),
			Password: r.PostFormValue("password"),
		}
		if err := user.CreateUser(); err != nil {
			log.Println(err)
		}

		http.Redirect(w, r, "/", 302)

	}
}

//ログイン画面に遷移する
func login(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	//セッションが作成できればログイン画面へ
	if err != nil {
		generateHTML(w, nil, "layout", "public_navbar", "login")
		//できなければTODO画面へ
	} else {
		http.Redirect(w, r, "/todos", 302)
	}
}

//ログイン処理を行う
func authenticate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	//ユーザー情報を取得する
	user, err := models.GetUserByEmail(r.PostFormValue("email"))
	if err != nil {
		log.Println(err)
		http.Redirect(w, r, "/login", 302)
	}
	//持ってきたユーザーのパスワードと入力されたものを暗号化したものが等しいか確認
	if user.Password == models.Encrypt(r.PostFormValue("password")) {
		//等しければセッションを作成
		session, err := user.CreateSession()
		if err != nil {
			log.Println(err)
		}
		//クッキーにセッションを保存
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.UUID,
			HttpOnly: true,
		}
		http.SetCookie(w, &cookie)
		//topにリダイレクト
		http.Redirect(w, r, "/", 302)
	} else {
		//等しくなければログインにリダイレクト
		http.Redirect(w, r, "/login", 302)
	}

}

//ログアウト処理を行う
func logout(w http.ResponseWriter, r *http.Request) {
	//クッキーを取得
	cookie, err := r.Cookie("_cookie")
	if err != nil {
		log.Println(err)
	}
	//クッキーの値から対応するセッションを削除
	if err != http.ErrNoCookie {
		session := models.Session{
			UUID: cookie.Value,
		}
		session.DelsteSessionByUUID()
	}
	//ログイン画面へ
	http.Redirect(w, r, "/login", 302)
}
