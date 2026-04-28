package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/google/uuid"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

type ProductInput struct {
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

type ProductPage struct {
	Data  []Product `json:"data"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	Total int       `json:"total"`
	Pages int       `json:"pages"`
}

var (
	db  = map[string]Product{}
	mu  sync.RWMutex
)

func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

func handleProducts(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		defer mu.RUnlock()
		q := r.URL.Query()
		name := strings.ToLower(q.Get("name"))
		category := strings.ToLower(q.Get("category"))
		minPrice := q.Get("min_price")
		maxPrice := q.Get("max_price")
		pageStr := q.Get("page")
		limitStr := q.Get("limit")
		page, _ := strconv.Atoi(pageStr)
		limit, _ := strconv.Atoi(limitStr)
		if page < 1 {
			page = 1
		}
		if limit < 1 {
			limit = 10
		}
		all := make([]Product, 0, len(db))
		for _, p := range db {
			if name != "" && !strings.Contains(strings.ToLower(p.Name), name) {
				continue
			}
			if category != "" && strings.ToLower(p.Category) != category {
				continue
			}
			if minPrice != "" {
				if v, err := strconv.ParseFloat(minPrice, 64); err == nil && p.Price < v {
					continue
				}
			}
			if maxPrice != "" {
				if v, err := strconv.ParseFloat(maxPrice, 64); err == nil && p.Price > v {
					continue
				}
			}
			all = append(all, p)
		}
		total := len(all)
		pages := (total + limit - 1) / limit
		if pages < 1 {
			pages = 1
		}
		start := (page - 1) * limit
		end := start + limit
		if start > total {
			start = total
		}
		if end > total {
			end = total
		}
		writeJSON(w, http.StatusOK, ProductPage{Data: all[start:end], Page: page, Limit: limit, Total: total, Pages: pages})

	case http.MethodPost:
		var input ProductInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		product := Product{ID: uuid.New().String(), Name: input.Name, Category: input.Category, Price: input.Price, Stock: input.Stock}
		mu.Lock()
		db[product.ID] = product
		mu.Unlock()
		writeJSON(w, http.StatusCreated, product)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleProduct(w http.ResponseWriter, r *http.Request) {
	id := strings.TrimPrefix(r.URL.Path, "/products/")

	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		p, ok := db[id]
		mu.RUnlock()
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
			return
		}
		writeJSON(w, http.StatusOK, p)

	case http.MethodPut:
		mu.RLock()
		_, ok := db[id]
		mu.RUnlock()
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
			return
		}
		var input ProductInput
		if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		product := Product{ID: id, Name: input.Name, Category: input.Category, Price: input.Price, Stock: input.Stock}
		mu.Lock()
		db[id] = product
		mu.Unlock()
		writeJSON(w, http.StatusOK, product)

	case http.MethodDelete:
		mu.Lock()
		_, ok := db[id]
		if ok {
			delete(db, id)
		}
		mu.Unlock()
		if !ok {
			writeJSON(w, http.StatusNotFound, map[string]string{"error": "product not found"})
			return
		}
		writeJSON(w, http.StatusOK, map[string]string{"message": "product deleted"})

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok", "service": "go-rest"})
	})
	mux.HandleFunc("/products", handleProducts)
	mux.HandleFunc("/products/", handleProduct)

	log.Println("Go REST API running on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
