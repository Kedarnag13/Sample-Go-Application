package users

import (
	"database/sql"
	"encoding/json"
	"github.com/Kedarnag13/Sample-Go-Application/api/v1/models"
	_ "github.com/lib/pq"
	"io/ioutil"
	"net/http"
)

type User struct{}

var Foo User

func (u User) Create(rw http.ResponseWriter, req *http.Request) {
	// This line will help in reading
	// 1. Read the contents that is coming in the body of the request.
	// 2. The library we are using for this is "ioutil".
	// 3. "ioutil" uses one of the methods, "ReadAll" to read the contents that is sent in the body of the request.
	body, err := ioutil.ReadAll(req.Body)

	// This line helps us create a global variable (struct variable) to access all the fields that are in the struct.
	var user models.User

	// This line is for error handling, there are two ways of handling the errors.
	// 1. "panic" - Compile.
	// 2. "Fatal" - Runtime.
	if err != nil {
		panic(err)
	}

	// To Decode the json we use the following method.
	err = json.Unmarshal(body, &user)

	if err != nil {
		panic(err)
	}

	// To open a connection with the postgres database, we should mention all the required information such as dbname, password, host and sslmode
	db, err := sql.Open("postgres", "dbname=go_training password=password host=localhost sslmode=disable")

	if err != nil {
		panic(err)
	}

	// To insert data into the relation/table, insert command is used and syntax is as below.
	// The reason for multiple assignment is, the insert command can behave in 2 ways
	// Success - If the values are inserted, insert_user will have the success result.
	// Failure - If the values are not inserted, err will have the failure message.
	insert_user, err := db.Query("INSERT INTO users(id, name) VALUES($1, $2)", user.Id, user.Name)
	if err != nil || insert_user == nil {
		panic(err)
	}

}
