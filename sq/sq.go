package sq

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"fmt"
)

func SqConnection() *sql.DB{
	db, err := sql.Open("mysql", "root:12345678@tcp(127.0.0.1:3306)/Books?parseTime=true")

	if err != nil{
		fmt.Println(err)
		return nil
	}
	return db
}