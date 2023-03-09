package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Post struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
type Response struct {
	Status  int
	Message string
	Data    interface{}
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/belajargolang")
	if err != nil {
		return nil, err
	}

	return db, nil
}
func handlerIndex(w http.ResponseWriter, r *http.Request) {
	// var filepath = path.Join("views", "index.html")
	// var tmpl, err = template.ParseFiles(filepath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	getData, err := db.Query("SELECT * FROM post")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer getData.Close()

	post := Post{}
	var result []Post

	for getData.Next() {
		var err = getData.Scan(&post.Id, &post.Title, &post.Content)

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, post)
	}

	// err = tmpl.Execute(w, result)
	// if err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	fmt.Println(err.Error())
	// }
	var response Response
	response.Status = 200
	response.Message = "Get Data Success"
	response.Data = result
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
func handlerDetail(w http.ResponseWriter, r *http.Request) {
	// var filepath = path.Join("views", "detail.html")
	// var tmpl, err = template.ParseFiles(filepath)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	vars := mux.Vars(r)
	id := vars["id"]
	// getId := r.URL.Query().Get("id")

	getData, err := db.Query("SELECT * FROM post WHERE id=?", id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer getData.Close()
	post := Post{}
	var result []Post

	for getData.Next() {
		err = getData.Scan(&post.Id, &post.Title, &post.Content)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, post)
	}

	// err = tmpl.Execute(w, post)
	// if err != nil {
	// 	// http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	fmt.Println(err.Error())
	// }
	var response Response
	response.Status = 200
	response.Message = "Get Data Success"
	response.Data = result
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
}
func handlerAdd(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "tambah.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.Execute(w, tmpl)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
	}
}
func handlerInsert(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// crud
	var post Post
	if r.Method == "POST" {
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		_, err = db.Exec("INSERT INTO post (title,content) VALUES (?,?)", post.Title, post.Content)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	defer db.Close()
	fmt.Println("Insert Succes")
	var response Response
	response.Status = 200
	response.Message = "Insert Succes"
	response.Data = post
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
	// http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func handlerEdit(w http.ResponseWriter, r *http.Request) {
	var filepath = path.Join("views", "edit.html")
	var tmpl, err = template.ParseFiles(filepath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	getId := r.URL.Query().Get("id")

	getData, err := db.Query("SELECT * FROM post WHERE id=?", getId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer getData.Close()
	post := Post{}
	for getData.Next() {
		err = getData.Scan(&post.Id, &post.Title, &post.Content)

		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	err = tmpl.Execute(w, post)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println(err.Error())
	}
}
func handlerUpdate(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var post Post
	if r.Method == "PUT" {
		err := json.NewDecoder(r.Body).Decode(&post)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// id := r.FormValue("id")
		vars := mux.Vars(r)
		id := vars["id"]
		_, err = db.Exec("UPDATE post SET title =? ,content =? WHERE id=?", post.Title, post.Content, id)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
	defer db.Close()
	fmt.Println("Update Succes")
	var response Response
	response.Status = 200
	response.Message = "Update Succes"
	response.Data = []Post{}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
	// http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func handlerDelete(w http.ResponseWriter, r *http.Request) {
	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// ambil id dari url html
	// getId := r.URL.Query().Get("id")
	vars := mux.Vars(r)
	id := vars["id"]

	_, err = db.Exec("DELETE FROM post WHERE id=?", id)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()
	fmt.Println("Delete Succes")
	var response Response
	response.Status = 200
	response.Message = "Insert Succes"

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(response)
	// http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
func main() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/", handlerIndex)
	r.HandleFunc("/detail/{id}", handlerDetail).Methods("GET")
	r.HandleFunc("/add", handlerAdd)
	r.HandleFunc("/insert", handlerInsert)
	r.HandleFunc("/edit/{id}", handlerEdit)
	r.HandleFunc("/update/{id}", handlerUpdate).Methods("PUT")
	r.HandleFunc("/delete/{id}", handlerDelete).Methods("DELETE")

	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
