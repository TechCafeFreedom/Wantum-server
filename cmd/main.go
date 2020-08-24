package main

import (
	"fmt"
	"net/http"
	"wantum/db/mysql"
	"wantum/pkg/api/middleware"
	tx "wantum/pkg/infrastructure/mysql"
	"wantum/pkg/tlog"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	// .envファイルの読み込み
	if err := godotenv.Load(); err != nil {
		tlog.GetAppLogger().Error(fmt.Sprintf("failed to load .env file: %v", err))
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
	c := cors.New(
		cors.Options{
			AllowedHeaders:   []string{"*"},
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"*"},
			AllowCredentials: true,
		},
	)
	r := mux.NewRouter()

	// connection testAPI
	r.HandleFunc("/ping/{hoge}/{fuga}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		fmt.Fprintf(w, "Hello %v, world! %v", vars["hoge"], vars["fuga"])
	})

	// userAPI
	firebaseAuth := r.PathPrefix("").Subrouter()
	r.HandleFunc("/users", userAPI.GetAllUsers).Methods("GET")

	firebaseAuth.Use(firebaseClient.MiddlewareFunc)
	{
		firebaseAuth.HandleFunc("/users", userAPI.CreateNewUser).Methods("POST")
		firebaseAuth.HandleFunc("/users/self", userAPI.GetAuthorizedUser).Methods("GET")
	}

	// port: 8080
	srv := &http.Server{
		Addr:    ":8080",
		Handler: c.Handler(r),
	}
	if err := srv.ListenAndServe(); err != nil {
		tlog.GetAppLogger().Fatal(err.Error())
	}
}
