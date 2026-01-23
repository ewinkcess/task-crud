package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Harga       int    `json:"harga"`
	Stok        int    `json:"stok"`
	Description string `json:"description"`
}

var category = []Category{
	{ID: 1, Name: "LED TV 43 inc", Harga: 10000000, Stok: 10, Description: "Google tv 4k UHD"},
	{ID: 2, Name: "Mesin cuci", Harga: 20000000, Stok: 20, Description: "Front Loading 10kg"},
	{ID: 3, Name: "Kulkas", Harga: 30000000, Stok: 30, Description: "Side by Side"},
	{ID: 4, Name: "AC", Harga: 40000000, Stok: 40, Description: "Split Standar"},
}

func getCategoriesByID(w http.ResponseWriter, r *http.Request) {
	idStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Categories ID", http.StatusBadRequest)
		return
	}
	for _, p := range category {
		if p.ID == id {
			w.Header().Set("Content-type", "application/json")
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	http.Error(w, "Categories belum ada", http.StatusNotFound)
}
func updateCategories(w http.ResponseWriter, r *http.Request) {
	idrStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idrStr)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	//get data dari request
	var updateCategories Category
	err = json.NewDecoder(r.Body).Decode(&updateCategories)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i := range category {
		if category[i].ID == id {
			updateCategories.ID = id       //keep id
			category[i] = updateCategories // update index i

			w.Header().Set("Content_type", "application/json")
			json.NewEncoder(w).Encode(updateCategories)
			return
		}
	}

	http.Error(w, "Produk belum ada", http.StatusNotFound)
}
func deleteCategories(w http.ResponseWriter, r *http.Request) {
	idrStr := strings.TrimPrefix(r.URL.Path, "/api/categories/")
	id, err := strconv.Atoi(idrStr)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	for i, p := range category {
		if p.ID == id {
			category = append(category[:i], category[i+1:]...)
			w.Header().Set("Content-Type", "application/jdon")
			json.NewEncoder(w).Encode(map[string]string{
				"message": "sukses delete",
			})
		}
	}
}
func main() {
	//GET by ID
	http.HandleFunc("/api/categories/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			getCategoriesByID(w, r)
		} else if r.Method == "PUT" {
			updateCategories(w, r)
		} else if r.Method == "DELETE" {
			deleteCategories(w, r)
		}
	})

	// GET ALL & Created Categories localhost:8080/api/catogeries
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(category)
		} else if r.Method == "POST" {
			var newCategory Category
			err := json.NewDecoder(r.Body).Decode(&newCategory)
			if err != nil {
				http.Error(w, "Invalid Request", http.StatusBadRequest)
				return
			}
			newCategory.ID = len(category) + 1
			category = append(category, newCategory)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newCategory)

		}
	})

	//localhost:8080/health
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"status":  "ok",
			"message": "Hello World!",
		})

	})
	fmt.Println("Hello World")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("gagal running server")
	}
}
