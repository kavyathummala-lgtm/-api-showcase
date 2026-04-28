package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

type Product struct {
	ID       string  `json:"id"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
}

var (
	db  = map[string]Product{}
	mu  sync.RWMutex
)

type ProductPage struct {
	Data  []Product `json:"data"`
	Page  int       `json:"page"`
	Limit int       `json:"limit"`
	Total int       `json:"total"`
	Pages int       `json:"pages"`
}

var productPageType = graphql.NewObject(graphql.ObjectConfig{
	Name: "ProductPage",
	Fields: graphql.Fields{
		"data":  &graphql.Field{Type: graphql.NewList(productType)},
		"page":  &graphql.Field{Type: graphql.Int},
		"limit": &graphql.Field{Type: graphql.Int},
		"total": &graphql.Field{Type: graphql.Int},
		"pages": &graphql.Field{Type: graphql.Int},
	},
})

var productType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Product",
	Fields: graphql.Fields{
		"id":       &graphql.Field{Type: graphql.String},
		"name":     &graphql.Field{Type: graphql.String},
		"category": &graphql.Field{Type: graphql.String},
		"price":    &graphql.Field{Type: graphql.Float},
		"stock":    &graphql.Field{Type: graphql.Int},
	},
})

func buildSchema() (graphql.Schema, error) {
	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"products": {
				Type: productPageType,
				Args: graphql.FieldConfigArgument{
					"name":     &graphql.ArgumentConfig{Type: graphql.String},
					"category": &graphql.ArgumentConfig{Type: graphql.String},
					"minPrice": &graphql.ArgumentConfig{Type: graphql.Float},
					"maxPrice": &graphql.ArgumentConfig{Type: graphql.Float},
					"page":     &graphql.ArgumentConfig{Type: graphql.Int},
					"limit":    &graphql.ArgumentConfig{Type: graphql.Int},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					mu.RLock()
					defer mu.RUnlock()
					name, _ := p.Args["name"].(string)
					category, _ := p.Args["category"].(string)
					minPrice, hasMin := p.Args["minPrice"].(float64)
					maxPrice, hasMax := p.Args["maxPrice"].(float64)
					page, _ := p.Args["page"].(int)
					limit, _ := p.Args["limit"].(int)
					if page < 1 {
						page = 1
					}
					if limit < 1 {
						limit = 10
					}
					all := make([]Product, 0, len(db))
					for _, v := range db {
						if name != "" && !strings.Contains(strings.ToLower(v.Name), strings.ToLower(name)) {
							continue
						}
						if category != "" && !strings.EqualFold(v.Category, category) {
							continue
						}
						if hasMin && v.Price < minPrice {
							continue
						}
						if hasMax && v.Price > maxPrice {
							continue
						}
						all = append(all, v)
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
					return ProductPage{Data: all[start:end], Page: page, Limit: limit, Total: total, Pages: pages}, nil
				},
			},
			"product": {
				Type: productType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					mu.RLock()
					defer mu.RUnlock()
					if v, ok := db[id]; ok {
						return v, nil
					}
					return nil, nil
				},
			},
		},
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createProduct": {
				Type: productType,
				Args: graphql.FieldConfigArgument{
					"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"category": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"price":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
					"stock":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					product := Product{
						ID:       uuid.New().String(),
						Name:     p.Args["name"].(string),
						Category: p.Args["category"].(string),
						Price:    p.Args["price"].(float64),
						Stock:    p.Args["stock"].(int),
					}
					mu.Lock()
					db[product.ID] = product
					mu.Unlock()
					return product, nil
				},
			},
			"updateProduct": {
				Type: productType,
				Args: graphql.FieldConfigArgument{
					"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"name":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"category": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"price":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
					"stock":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Int)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					mu.Lock()
					defer mu.Unlock()
					if _, ok := db[id]; !ok {
						return nil, nil
					}
					product := Product{
						ID:       id,
						Name:     p.Args["name"].(string),
						Category: p.Args["category"].(string),
						Price:    p.Args["price"].(float64),
						Stock:    p.Args["stock"].(int),
					}
					db[id] = product
					return product, nil
				},
			},
			"deleteProduct": {
				Type: graphql.Boolean,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id := p.Args["id"].(string)
					mu.Lock()
					defer mu.Unlock()
					if _, ok := db[id]; !ok {
						return false, nil
					}
					delete(db, id)
					return true, nil
				},
			},
		},
	})

	return graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}

func main() {
	schema, err := buildSchema()
	if err != nil {
		log.Fatalf("failed to build schema: %v", err)
	}

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"status": "ok", "service": "go-graphql"})
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query     string `json:"query"`
			Variables map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		result := graphql.Do(graphql.Params{
			Schema:         schema,
			RequestString:  params.Query,
			VariableValues: params.Variables,
		})
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	})

	log.Println("Go GraphQL API running on :8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
