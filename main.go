package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/user/login", login)
	http.HandleFunc("user/get_info", getInfo)

	port := 8080
	fmt.Printf("Server listening on http://localhost::%d...\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	userID := queryParams.Get("user_id")
	password := queryParams.Get("password")

	if r.Method != http.MethodGet {
		data := map[string]interface{}{
			"status":  false,
			"ret_str": "404 Invalid Route",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(data)
		return
	} 

	if userID == "" || password == "" {
		data := map[string]interface{}{
			"status":  false,
			"ret_str": "User ID and Password Required",
		}

		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(data)
		return
	}

	if(userID != "kamrul" || password != "1234"){
		data := map[string]interface{}{
			"status":  false,
			"ret_str": "Invalid user",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(data)
		return
	}

	data := map[string]interface{}{
		"status":  true,
		"ret_str": "Login Successful",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
func getInfo(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		data := map[string]interface{}{
			"status":  false,
			"ret_str": "404 Invalid Route",
		}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(data)
		return
	} 

	data := map[string]interface{}{
		"status": true,
		"user_info": map[string]interface{}{
			"name":   "Kamrul Hasan",
			"mobile": "01874007261",
			"email":  "kamrulcse7@gmail.com",
		},
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
