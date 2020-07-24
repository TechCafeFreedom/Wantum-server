package main

import (
	"fmt"
	"log"
	"net/http"
	"wantum/db/mysql"
	"wantum/pkg/api/middleware"
	tx "wantum/pkg/infrastructure/mysql"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		log.Printf("failed to load .env file: %v", err)
	}

	// DBインスタンスの作成
	dbInstance := mysql.CreateSQLInstance()
	defer dbInstance.Close()

	// トランザクションマネージャーの作成
	masterTxManager := tx.NewDBMasterTxManager(dbInstance)

	// APIインスタンスの作成
	userAPI := InitUserAPI(masterTxManager)

	// firebase middlewareの作成
	firebaseClient := middleware.CreateFirebaseInstance()

	// CORS対応
	r := mux.NewRouter()

	// connection testAPI
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello, world!")
	})

	// userAPI
	firebaseAuth := r.PathPrefix("").Subrouter()
	r.HandleFunc("/users", userAPI.GetAllUsers).Methods("GET")

	firebaseAuth.Use(firebaseClient.MiddlewareFunc)
	{
		firebaseAuth.HandleFunc("/users", userAPI.CreateNewUser).Methods("POST")
		firebaseAuth.HandleFunc("/users/self", userAPI.GetUserProfile).Methods("GET")
	}

	// port: 8080
	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	log.Fatal(srv.ListenAndServe())
}
