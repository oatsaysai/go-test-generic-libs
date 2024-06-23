package mariaDB

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

func ConnectDB() (*sql.DB, error) {
	dbIP := viper.GetString("DB_IP")
	dbPort := viper.GetString("DB_PORT")
	dbUser := viper.GetString("DB_USER")
	dbPass := viper.GetString("DB_PASS")
	dbName := viper.GetString("DB_NAME")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		dbUser,
		dbPass,
		dbIP,
		dbPort,
		dbName,
	)

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	} else {
		log.Println("MariaDB connected")
	}

	return db, nil
}
