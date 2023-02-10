package main

import (
	"banking/handler"
	"banking/logs"
	"banking/repository"
	"banking/service"
	"fmt"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {

	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepositoryDb := repository.NewCustomerRepositoryDb(db)
	customerService := service.NewCustomerService(customerRepositoryDb)
	customerHandler := handler.NewCustomerHandler(customerService)

	accountRepositoryDb := repository.NewAccountRepositoryDb(db)
	accountService := service.NewAccountService(accountRepositoryDb)
	accountHandler := handler.NewAccountHandler(accountService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.GetCustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.GetCustomer).Methods(http.MethodGet)

	router.HandleFunc("/customers/{customerID:[0-9]+}/account", accountHandler.GetAccounts).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}/account", accountHandler.NewAccount).Methods(http.MethodPost)

	// log.Printf("Banking service starting at port %v", viper.GetInt("app.port"))
	logs.Info("Banking service starting at port" + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)
}

func initConfig() {

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}
	time.Local = ict

}

func initDatabase() *sqlx.DB {

	// ** sql config setup
	dataSource := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)

	db, err := sqlx.Open(viper.GetString("db.driver"), dataSource)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)

	return db
}
