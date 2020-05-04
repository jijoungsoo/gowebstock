package business_service

import (
	"fmt"
)

const (
	// Initialize connection constants.
	HOST        = "localhost"
	DB_DATABASE = "postgres"
	DB_USER     = "postgres"
	DB_PASSWORD = "pwd"
)

func Get_connection_string() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", HOST, DB_USER, DB_PASSWORD, DB_DATABASE)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
