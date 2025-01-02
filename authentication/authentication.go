package authentication

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	hashing "practise/authentication/Hash"
	"practise/controllers/connection"
	models "practise/models"

	"time"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = connection.GetUserCollection()

func Authorized(r *http.Request) error {

	username := r.FormValue("username")
	var AuthError = errors.New("Unauthorized")

	filter := bson.M{"username": username}

	var user models.Login
	err := collection.FindOne(context.Background(), filter).Decode(&user)

	if err != nil {
		return AuthError
	}

	st, err := r.Cookie("session_token")
	if err != nil || st.Value == "" || st.Value != user.SessionToken {
		return AuthError
	}

	csrf := r.Header.Get("X-CSRF-Token")
	if csrf != user.CSRFToken || csrf == "" {
		return AuthError
	}
	fmt.Println("CSRF Validation complete, Welcome :" + username)
	return nil

}

func SignUpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", err)
		return
	}

	var userData models.Login

	username := r.FormValue("username")
	password := r.FormValue("password")
	// _ = json.NewDecoder(r.Body).Decode(&userData)

	if len(username) < 5 || len(password) < 8 {
		err := http.StatusNotAcceptable
		http.Error(w, "Invalid username/password", err)
		return
	}
	hashedPassword, _ := hashing.HashPassword(password)
	userData.Username = username
	userData.HashedPassword = hashedPassword

	_, err := collection.InsertOne(context.Background(), userData)

	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode("Sign up successfull")

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", err)
		return
	}

	// var user models.Employee
	// _ = json.NewDecoder(r.Body).Decode(&user)
	username := r.FormValue("username")
	password := r.FormValue("password")

	// json.NewEncoder(w).Encode(user)
	filter := bson.M{"username": username}

	var findedUser models.Login
	err := collection.FindOne(context.Background(), filter).Decode(&findedUser)

	if err != nil {
		json.NewEncoder(w).Encode("No user found")
		return

	} else if !hashing.CheckPasswordHash(password, findedUser.HashedPassword) || findedUser.Username != username {
		json.NewEncoder(w).Encode("Username or Password Incorrect")
		return
	}

	sessionToken := hashing.GenerateToken(32)
	csrfToken := hashing.GenerateToken(32)

	http.SetCookie(w, &http.Cookie{
		Name:     "session_token",
		Value:    sessionToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "csrf_token",
		Value:    csrfToken,
		Expires:  time.Now().Add(24 * time.Hour),
		HttpOnly: false,
	})

	findedUser.SessionToken = sessionToken
	findedUser.CSRFToken = csrfToken

	update := bson.M{"$set": findedUser}
	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Login Successfull")
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		err := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", err)
		return
	}

	err := Authorized(r)
	if err != nil {
		err := http.StatusUnauthorized
		http.Error(w, "Unauthorized", err)
		return
	}
	username := r.FormValue("username")
	filter := bson.M{"username": username}

	var findedUser models.Login
	err = collection.FindOne(context.Background(), filter).Decode(&findedUser)

	if err != nil {
		log.Fatal(err)
	}

	findedUser.CSRFToken = ""
	findedUser.SessionToken = ""

	update := bson.M{"$set": findedUser}

	_, err = collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(w).Encode("Logout Successfull")

}
