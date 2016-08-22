package users

import (
	"encoding/json"
	"fmt"
	"github.com/Kedarnag13/Sample-Go-Application/api/v1/models"
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

	fmt.Printf("Id:%v", user.Id)
	fmt.Printf("Name:%v", user.Name)
}
