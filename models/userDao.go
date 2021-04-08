package models

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

var (
	server   = "192.168.25.197"
	port     = 1433
	user     = "sa"
	password = "vncert"
	database = "ThuRac"
)

func Connection() *sql.DB{
	connString := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%d;database=%s;", server, user, password, port, database)
	conn, _ := sql.Open("mssql", connString)
	return conn
}

// checkDateExist
func CheckDateExist(date string) int{
	result, err := Connection().Query("SELECT id from dataColletion where date = ?", date)

	if err != nil {
		panic(err.Error())
	}

	count := 0

	for result.Next() {
		var id int

		result.Scan(&id, &date)

		if err != nil{
			return -1
		}
		count++
	}

	result.Close()
	return count
}

func CheckRole(user string,password string) int{
	var role int

	result, err := Connection().Query("SELECT [role] from users where account = ? and passwd = ?", user,password)

	if err != nil {
		panic(err.Error())
	}

	for result.Next(){
		result.Scan(&role)
	}

	return role
}

func (u Users) Exists(id int) bool{
	exists := false
		for _, user := range u{
			if user.Id == id{
				return true
			}
		}
	return exists
}

func (u Users) FindByName(name string) (User, error){
	for _, user := range u{
		if user.Username == name{
			return user, nil
		}
	}
	return User{}, errors.New("NOT_FOUND_USER")
}


