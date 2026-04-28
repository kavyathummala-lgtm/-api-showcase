package main

import (
	"context"
	"log"
	"net"
	"strings"
	"sync"

	pb "github.com/example/go-grpc/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedProductServiceServer
	mu sync.RWMutex
	db map[string]*pb.Product
}

func (s *server) GetProducts(_ context.Context, req *pb.ProductFilter) (*pb.ProductList, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	all := make([]*pb.Product, 0, len(s.db))
	for _, p := range s.db {
		if req.Name != "" && !strings.Contains(strings.ToLower(p.Name), strings.ToLower(req.Name)) {
			continue
		}
		if req.Category != "" && !strings.EqualFold(p.Category, req.Category) {
			continue
		}
		if req.MinPrice > 0 && p.Price < req.MinPrice {
			continue
		}
		if req.MaxPrice > 0 && p.Price > req.MaxPrice {
			continue
		}
		all = append(all, p)
	}
	total := int32(len(all))
	page := req.Page
	if page < 1 {
		page = 1
	}
	limit := req.Limit
	if limit < 1 {
		limit = 10
	}
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
	return &pb.ProductList{Products: all[start:end], Total: total, Page: page, Pages: pages}, nil
}

func (s *server) GetProduct(_ context.Context, req *pb.ProductId) (*pb.Product, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	p, ok := s.db[req.Id]
	if !ok {
		return nil, status.Errorf(codes.NotFound, "product %s not found", req.Id)
	}
	return p, nil
}

func (s *server) CreateProduct(_ context.Context, req *pb.ProductInput) (*pb.Product, error) {
	product := &pb.Product{
		Id:       uuid.New().String(),
		Name:     req.Name,
		Category: req.Category,
		Price:    req.Price,
		Stock:    req.Stock,
	}
	s.mu.Lock()
	s.db[product.Id] = product
	s.mu.Unlock()
	return product, nil
}

func (s *server) UpdateProduct(_ context.Context, req *pb.UpdateProductRequest) (*pb.Product, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.db[req.Id]; !ok {
		return nil, status.Errorf(codes.NotFound, "product %s not found", req.Id)
	}
	product := &pb.Product{
		Id:       req.Id,
		Name:     req.Input.Name,
		Category: req.Input.Category,
		Price:    req.Input.Price,
		Stock:    req.Input.Stock,
	}
	s.db[req.Id] = product
	return product, nil
}

func (s *server) DeleteProduct(_ context.Context, req *pb.ProductId) (*pb.DeleteResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.db[req.Id]; !ok {
		return nil, status.Errorf(codes.NotFound, "product %s not found", req.Id)
	}
	delete(s.db, req.Id)
	return &pb.DeleteResponse{Success: true, Message: "product deleted"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductServiceServer(s, &server{db: make(map[string]*pb.Product)})
	log.Println("Go gRPC server running on :50052")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
