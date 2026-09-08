package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		var user User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}

		if user.Name == "" || user.Age <= 0 {
			http.Error(w, "Missing required fields", http.StatusUnprocessableEntity)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "User %s successfully processed", user.Name)
	}

	brokenJSON := `{"name": "Ivan", "age": 25`
	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(brokenJSON))
	rr := httptest.NewRecorder()

	handler(rr, req)

	fmt.Printf("HTTP Status Code: %d\n", rr.Code)
	fmt.Printf("Server Response: %s\n", rr.Body.String())
}
