package routes

import (
	"fmt"
	"net/http"

	jwtHelper "github.com/quanhnv/eLotus-challenges/auth"
	sqliteHelper "github.com/quanhnv/eLotus-challenges/database"
)

func Register(w http.ResponseWriter, r *http.Request) {
	//Check request method must be Post
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userName := r.PostFormValue("username")
	passWord := r.PostFormValue("password")

	//validate input
	isValid, message := validateUser(userName, passWord)
	if !isValid {
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	err := sqliteHelper.InsertUser(userName, passWord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	//w.Write([]byte("Register successful!"))
	fmt.Fprintf(w, "Register successful!")
}

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	userName := r.PostFormValue("username")
	passWord := r.PostFormValue("password")

	fmt.Println("%s login", userName)

	isValid, message := validateUser(userName, passWord)
	if !isValid {
		http.Error(w, message, http.StatusBadRequest)
		return
	}

	err := sqliteHelper.Login(userName, passWord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, errCreate := jwtHelper.CreateAndSetCookie(w, userName)
	if errCreate != nil {
		http.Error(w, "Generate token failed", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Login successful!")
}

func validateUser(username string, password string) (bool, string) {
	isValid := true
	message := ""
	if username == "" {
		isValid = false
		message = "Username not empty "
	}
	if password == "" {
		isValid = false
		message += "Password not empty "
	}
	return isValid, message
}
