package api

import (
	"encoding/json"
	"fem/museum/data"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var newData data.Exhibition
		err := json.NewDecoder(r.Body).Decode(&newData)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		data.AddData(newData)
		w.WriteHeader(http.StatusCreated)
	} else {
		http.Error(w, "Unsupported Method", http.StatusMethodNotAllowed)
	}
}
