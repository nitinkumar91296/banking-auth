package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nitkumar91296/banking-auth/domain"
	"github.com/nitkumar91296/banking-auth/service"
)

func Start() {
	// SanityCheck()

	router := mux.NewRouter()

	authRepository := domain.NewAuthRepository(getDbClient())
	ah := AuthHandler{repo: authRepository, service: service.NewAuthService(authRepository)}SELECT list is not in GROUP BY clause and contains nonaggregated column 'banking.u.username' which is not functionally dependent on columns in GROUP BY clause; this is incompatible with sql_mode=only_full_group_by


	router.HandleFunc("/auth/login", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/register", ah.Login).Methods(http.MethodPost)
	router.HandleFunc("/auth/verify", ah.Login).Methods(http.MethodGet)

	// addr := os.Getenv("AUTH_SERVER_ADDRESS")
	// port := os.Getenv("AUTH_SERVER_PORT")
	addr, port := "localhost", "8010"
	log.Println(fmt.Sprintf("Starting OAuth server at %s:%s", addr, port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", addr, port), router))
}

func SanityCheck() {
	if os.Getenv("AUTH_SERVER_ADDRESS") == "" || os.Getenv("AUTH_SERVER_PORT") == "" {
		log.Fatal("server variables are not defined...")
	}

	if os.Getenv("DB_USER") == "" ||
		os.Getenv("DB_PASSWORD") == "" ||
		os.Getenv("DB_ADDR") == "" ||
		os.Getenv("DB_PORT") == "" ||
		os.Getenv("DB_NAME") == "" {
		log.Fatal("db variables are not defined ...")
	}
}

func getDbClient() *sqlx.DB {
	// dbUser := os.Getenv("DB_USER")
	// dbPasswd := os.Getenv("DB_PASSWORD")
	// dbAddr := os.Getenv("DB_ADDR")
	// dbPort := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPasswd, dbAddr, dbPort, dbName)
	dsn := "root:Nitin@91296@tcp(localhost:3306)/banking"
	client, err := sqlx.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
