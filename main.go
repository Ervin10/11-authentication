package main

import (
	"7-routing/connection"
	"7-routing/handler"
	"7-routing/middleware"
	"fmt"
	"net/http"

	"github.com/gorilla/mux" //npm routing
)

func main() {
	connection.DatabaseConnect()
	route := mux.NewRouter() // buat router dengan mux
	// port := "5000"

	// route untuk menginisialisasi folder public agar dapat dibaca
	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/", http.FileServer(http.Dir("./public/"))))
	route.PathPrefix("/uploads/").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads/"))))

	route.HandleFunc("/", handler.HandleHome).Methods("GET") // routing jadikan function home
	route.HandleFunc("/contact", handler.HandleContact).Methods("GET")
	route.HandleFunc("/project", handler.HandleProject).Methods("GET")
	route.HandleFunc("/add-project", middleware.UploadFile(handler.HandleAddProject)).Methods("POST") // package middleware: upload file
	route.HandleFunc("/project-detail/{id}", handler.HandleDetailProject).Methods("GET")
	route.HandleFunc("/edit-project/{id}", handler.HandleEditProject).Methods("GET")
	route.HandleFunc("/edit-project/{id}", middleware.UploadFile(handler.HandleEditProject)).Methods("POST")
	route.HandleFunc("/delete/{id}", handler.HandleDeleteProject).Methods("GET")
	route.HandleFunc("/register", handler.HandleRegister)
	route.HandleFunc("/login", handler.HandleLogin)
	route.HandleFunc("/logout", handler.HandleLogout)

	fmt.Println("Server sedang berjalan di port 4700")
	http.ListenAndServe("Localhost:4700", route) // panggil untuk dapat diakses di browser par1: string, par2: route
}

// Cara insert table manual
// insert into tb_blogs (id, project_name, start_date, end_date, "desc", tech, image) values(3, 'Dumbways','2022-11-25','2024-01-25','Halo Dumbways','{node,react,angular}','3.png')
