package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type userSignupForm struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	var form userSignupForm

	if err := json.NewDecoder(r.Body).Decode(&form); err != nil {
		app.serverError(w, err)
		return
	}

	app.renderTemplate(w, form, "signup.html")
}

func (app *application) registerUser(w http.ResponseWriter, r *http.Request) {
	var form userSignupForm

	err := json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Gist created successfully",
		"data":    form,
	})
	if err != nil {
		log.Println("Error sending response:", err)
	}
}

func (app *application) userLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logging in")
}

func (app *application) userLogout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Logging out")
}
