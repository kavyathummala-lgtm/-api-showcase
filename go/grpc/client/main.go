package main

import (
	"context"
	"log"
	"time"

	pb "github.com/example/go-grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Println("=== Create Product ===")
	created, err := client.CreateProduct(ctx, &pb.ProductInput{
		Name:     "Laptop",
		Category: "Electronics",
		Price:    999.99,
		Stock:    50,
	})
	if err != nil {
		log.Fatalf("CreateProduct: %v", err)
	}
	log.Printf("Created: %+v", created)

	log.Println("\n=== Get All Products ===")
	list, err := client.GetProducts(ctx, &pb.ProductFilter{})
	if err != nil {
		log.Fatalf("GetProducts: %v", err)
	}
	for _, p := range list.Products {
		log.Printf("  %s: %s - $%.2f", p.Id, p.Name, p.Price)
	}

	log.Println("\n=== Search Products (Electronics only) ===")
	filtered, err := client.GetProducts(ctx, &pb.ProductFilter{Category: "Electronics"})
	if err != nil {
		log.Fatalf("GetProducts filtered: %v", err)
	}
	for _, p := range filtered.Products {
		log.Printf("  %s: %s - $%.2f", p.Id, p.Name, p.Price)
	}

	log.Println("\n=== Update Product ===")
	updated, err := client.UpdateProduct(ctx, &pb.UpdateProductRequest{
		Id: created.Id,
		Input: &pb.ProductInput{
			Name:     "Laptop Pro",
			Category: "Electronics",
			Price:    1299.99,
			Stock:    30,
		},
	})
	if err != nil {
		log.Fatalf("UpdateProduct: %v", err)
	}
	log.Printf("Updated: %+v", updated)

	log.Println("\n=== Delete Product ===")
	resp, err := client.DeleteProduct(ctx, &pb.ProductId{Id: created.Id})
	if err != nil {
		log.Fatalf("DeleteProduct: %v", err)
	}
	log.Printf("Deleted: %v - %s", resp.Success, resp.Message)
}
