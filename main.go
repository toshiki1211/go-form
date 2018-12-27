package main

import (
	_"fmt"
    "html/template"
    "log"
	"net/http"
	"strings"
	"unicode/utf8"
)


type FormData struct {
	ValidateError bool
	ErrorMessages map[string]string
	Username string
	Password string
}

func (d *FormData) index(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 初期化
	d.ValidateError = false
	d.Username = ""
	d.Password = ""
	
	if r.Method == "GET" {
        t, _ := template.ParseFiles("public/index.tpl")
        t.Execute(w, d)
    } else {
		d.Username = r.FormValue("username")
		d.Password = r.FormValue("password")
		d.ValidateError = true

		t, _ := template.ParseFiles("public/index.tpl")
        t.Execute(w, d)
	}
}

func (d *FormData) confirm(w http.ResponseWriter, r *http.Request) {
	d.ValidateError = false
	
    if r.Method == "GET" {
        t, _ := template.ParseFiles("public/confirm.tpl")
        t.Execute(w, d)
    } else {
		err := r.ParseForm()
		if err != nil {
			// エラー処理
		}

		d.Username = strings.TrimSpace(r.FormValue("username"))
		d.Password = strings.TrimSpace(r.FormValue("password"))

		d.ErrorMessages = d.validate()

		if(d.ValidateError == false) {
			t, _ := template.ParseFiles("public/confirm.tpl") 
        	t.Execute(w, d)
		} else {
			t, _ := template.ParseFiles("public/index.tpl") 
        	t.Execute(w, d)
		}
    }
}

func  (d *FormData) complate(w http.ResponseWriter, r *http.Request) {
	d.ValidateError = false

	d.Username = strings.TrimSpace(r.FormValue("username"))
	d.Password = strings.TrimSpace(r.FormValue("password"))

	d.ErrorMessages = d.validate()

	if(d.ValidateError == false) {
		t, _ := template.ParseFiles("public/complate.tpl")
		t.Execute(w, d)
	} else {
		t, _ := template.ParseFiles("public/index.tpl")
		t.Execute(w, d)
	}
}

func (d *FormData) validate() (ErrorMessages map[string]string) {
	ErrorMessages = make(map[string]string)

	if(utf8.RuneCountInString(d.Username) == 0) {
		d.ValidateError = true
		ErrorMessages["username"] = "ユーザ名は必須です"
	}

	if(utf8.RuneCountInString(d.Password) == 0) {
		d.ValidateError = true
		ErrorMessages["password"] = "パスワードは必須です"
	}

	return ErrorMessages
}


func main() {	
	FormData := FormData{}

	//アクセスのルーティングを設定
    http.HandleFunc("/", FormData.index)
	http.HandleFunc("/confirm", FormData.confirm)
	http.HandleFunc("/complate", FormData.complate)
	
    err := http.ListenAndServe(":5000", nil) //監視するポートを設定
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}