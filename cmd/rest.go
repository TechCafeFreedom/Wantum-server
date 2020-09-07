package main

import (
	"fmt"
	"net/http"
	restWire "wantum/cmd/rest"
	restMiddleware "wantum/pkg/api/middleware/rest"
	"wantum/pkg/domain/repository"
	"wantum/pkg/tlog"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func StartRestAPIServer(firebaseClient restMiddleware.FirebaseAuth, masterTxManager repository.MasterTxManager) {
	// APIインスタンスの作成
	userAPI := restWire.InitUserAPI(masterTxManager)

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
