package app

import (
	"github.com/Darkhackit/banking-auth/domain"
	"github.com/Darkhackit/banking-auth/service"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

func Start() {
	dbClient := getDClient()
	handlers := loginHandlers{service.NewLoginService(domain.NewLoginRepositoryDB(dbClient))}
	router := mux.NewRouter()
	router.HandleFunc("/login", handlers.loginHandler).Methods("POST")
	router.HandleFunc("/register", handlers.RegisterHandler).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}

func getDClient() *sqlx.DB {
	client, err := sqlx.Open("mysql", "root:@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return client
}
