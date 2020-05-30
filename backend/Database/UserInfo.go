package Database

import (
	"database/sql"
	"encoding/json"
	"fmt"

	_ "github.com/lib/pq"
)

type VerifyMessage struct {
	Success int    `json:"success"`
	Message string `json:"message"`
}

func VerifyLogin(reqBody string, db *sql.DB) VerifyMessage {
	var result map[string]interface{}
	json.Unmarshal([]byte(reqBody), &result)
	email := result["email"]
	password := result["password"]
	sqlStatement := `SELECT password,name FROM users WHERE EMAIL=$1`
	res := db.QueryRow(sqlStatement, email)
	fmt.Println("res: ", res)
	var pass string
	var name string
	err := res.Scan(&pass, &name)
	if err != nil {
		output := VerifyMessage{0, "Invalid Email"}
		return output
	}
	fmt.Println("pass: ", pass)
	if pass != password {
		output := VerifyMessage{0, "Incorrect Password"}
		return output
	}
	output := VerifyMessage{1, name}
	return output
}
