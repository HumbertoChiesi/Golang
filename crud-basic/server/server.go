package server

import (
	db "crud/DB"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type user struct {
	ID    uint32 `json:"id"`
	Name  string `json:"nome"`
	Email string `json:"email"`
}

//CreateUser insert a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	//reading the request
	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Failed to read the request body"))
		return
	}

	var user user

	//putting the request content in an user struct variable
	if erro := json.Unmarshal(requestBody, &user); erro != nil {
		w.Write([]byte("ERROR while trying to conver user to struct"))
		return
	}

	fmt.Println(user)

	//connecting to the DB
	db, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("ERROR while trying to connect to the DB"))
		return
	}
	defer db.Close()

	//inserting the content into the DB
	statement, erro := db.Prepare("insert into usuarios (nome, email) values(?,?)")
	if erro != nil {
		w.Write([]byte("ERROR while trying to create the statement"))
		return
	}
	defer statement.Close()

	insertion, erro := statement.Exec(user.Name, user.Email)
	if erro != nil {
		w.Write([]byte("ERROR while trying to execute the statement"))
		return
	}

	insertedId, erro := insertion.LastInsertId()
	if erro != nil {
		w.Write([]byte("ERROR while trying to get the inserted ID"))
		return
	}

	//STATUS CODE
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("Successfully inserted the user!! Id: %d", insertedId)))
}

//SearchUsers get all the users saved in the DB
func SearchUsers(w http.ResponseWriter, r *http.Request) {

	//connecting with the DB
	db, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("FAILED to connect with the DB"))
		return
	}
	defer db.Close()

	lines, erro := db.Query("Select * from usuarios")
	if erro != nil {
		w.Write([]byte("ERROR trying to get the users from the DB"))
		return
	}

	var users []user
	for lines.Next() {
		var u user
		if erro := lines.Scan(&u.ID, &u.Name, &u.Email); erro != nil {
			w.Write([]byte("ERROR trying to scan the users"))
			return
		}
		users = append(users, u)
	}

	w.WriteHeader(http.StatusOK)
	if erro := json.NewEncoder(w).Encode(users); erro != nil {
		w.Write([]byte("ERROR while converting users to json"))
		return
	}
}

//SearchUser gets an especif user from the DB
func SearchUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ERROR while converting the parameter to int"))
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("FAILED to connect with the DB"))
		return
	}
	defer db.Close()

	line, erro := db.Query("select * from usuarios where id = ?", ID)
	if erro != nil {
		w.Write([]byte("ERROR while seaching the user"))
		return
	}

	var u user
	if line.Next() {
		if erro := line.Scan(&u.ID, &u.Name, &u.Email); erro != nil {
			w.Write([]byte("ERROR while scanning the user"))
			return
		}
	}

	if erro := json.NewEncoder(w).Encode(u); erro != nil {
		w.Write([]byte("ERROR while converting user to JSON"))
		return
	}
}

//UpdateUser changes an user data in the DB
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameters["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ERROR while converting the parameter to int"))
		return
	}

	requestBody, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("ERROR while reading the request body"))
		return
	}

	var u user
	if erro := json.Unmarshal(requestBody, &u); erro != nil {
		w.Write([]byte("ERROR while converting user to struct"))
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("ERROR trying to connect with the DB"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("update usuarios set nome = ?, email = ? where id = ?")
	if erro != nil {
		w.Write([]byte("ERROR trying to create the statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(u.Name, u.Email, ID); erro != nil {
		w.Write([]byte("ERROR while updating the user"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	parameter := mux.Vars(r)

	ID, erro := strconv.ParseUint(parameter["id"], 10, 32)
	if erro != nil {
		w.Write([]byte("ERROR while converting the parameter to int"))
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		w.Write([]byte("ERROR trying to connect with the DB"))
		return
	}
	defer db.Close()

	statement, erro := db.Prepare("delete from usuarios where id = ?")
	if erro != nil {
		w.Write([]byte("ERROR trying to create the statement"))
		return
	}
	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil {
		w.Write([]byte("ERROR while trying to execute the statement"))
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
