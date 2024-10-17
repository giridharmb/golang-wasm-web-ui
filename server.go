package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type FormDataServer struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Age       int      `json:"age"`
	Gender    string   `json:"gender"`
	Interests []string `json:"interests"`
	Comments  string   `json:"comments"`
}

func main() {
	http.HandleFunc("/submit", handleSubmit)

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSubmit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var formData FormDataServer
	err := json.NewDecoder(r.Body).Decode(&formData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process the data (in this example, we'll just perform some simple validations)
	errors := validateFormData(formData)

	response := make(map[string]interface{})

	if len(errors) > 0 {
		response["status"] = "error"
		response["errors"] = errors
	} else {
		response["status"] = "success"
		response["message"] = fmt.Sprintf("Thank you, %s! Your form has been processed.", formData.Name)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func validateFormData(data FormDataServer) []string {
	var errors []string

	if strings.TrimSpace(data.Name) == "" {
		errors = append(errors, "Name is required")
	}

	if strings.TrimSpace(data.Email) == "" {
		errors = append(errors, "Email is required")
	}

	if data.Age < 18 || data.Age > 120 {
		errors = append(errors, "Age must be between 18 and 120")
	}

	if strings.TrimSpace(data.Gender) == "" {
		errors = append(errors, "Gender is required")
	}

	if len(data.Interests) == 0 {
		errors = append(errors, "Please select at least one interest")
	}

	return errors
}
