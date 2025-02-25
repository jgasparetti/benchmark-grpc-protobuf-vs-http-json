package httpjson

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/mail"
)

// Start entrypoint
func Start() {
	http.HandleFunc("/", CreateUser)
	log.Println(http.ListenAndServe(":60001", nil))
}

// User type
type User struct {
	ID       string    `json:"id"`
	Email    string    `json:"email"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Other    string    `json:"other"`
	Field1   int64     `json:"field1"`
	Field2   float64   `json:"field2"`
	Field3   []string  `json:"field3"`
	Field4   []int64   `json:"field4"`
	Field5   []float32 `json:"field5"`
}

// Response type
type Response struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	User    *User  `json:"user"`
}

// CreateUser handler
func CreateUser(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		return
	}
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")
	validationErr := validate(user)
	if validationErr != nil {
		err = json.NewEncoder(w).Encode(Response{
			Code:    500,
			Message: validationErr.Error(),
		})
		if err != nil {
			return
		}
		return
	}

	user.ID = "1000000"
	err = json.NewEncoder(w).Encode(Response{
		Code:    200,
		Message: "OK",
		User:    &user,
	})
	if err != nil {
		return
	}
	return
}

func validate(in User) error {
	_, err := mail.ParseAddress(in.Email)
	if err != nil {
		return err
	}

	if len(in.Name) < 4 {
		return errors.New("Name is too short")
	}

	if len(in.Password) < 4 {
		return errors.New("Password is too weak")
	}

	return nil
}
