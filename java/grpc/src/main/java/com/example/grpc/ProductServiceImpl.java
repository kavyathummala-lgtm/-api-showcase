package com.example.grpc;

import com.example.grpc.proto.*;
import io.grpc.Status;
import io.grpc.stub.StreamObserver;
import net.devh.boot.grpc.server.service.GrpcService;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;
import java.util.stream.Collectors;

@GrpcService
public class ProductServiceImpl extends ProductServiceGrpc.ProductServiceImplBase {

    private final Map<String, Product> db = new ConcurrentHashMap<>();

    @Override
    public void getProducts(ProductFilter request, StreamObserver<ProductList> responseObserver) {
        List<Product> filtered = db.values().stream()
                .filter(p -> request.getName().isEmpty() || p.getName().toLowerCase().contains(request.getName().toLowerCase()))
                .filter(p -> request.getCategory().isEmpty() || p.getCategory().equalsIgnoreCase(request.getCategory()))
                .filter(p -> request.getMinPrice() <= 0 || p.getPrice() >= request.getMinPrice())
                .filter(p -> request.getMaxPrice() <= 0 || p.getPrice() <= request.getMaxPrice())
                .collect(Collectors.toList());
        int total = filtered.size();
        int page = request.getPage() > 0 ? request.getPage() : 1;
        int limit = request.getLimit() > 0 ? request.getLimit() : 10;
        int pages = Math.max(1, (int) Math.ceil((double) total / limit));
        int start = Math.min((page - 1) * limit, total);
        int end = Math.min(start + limit, total);
        ProductList list = ProductList.newBuilder()
                .addAllProducts(filtered.subList(start, end))
                .setTotal(total)
                .setPage(page)
                .setPages(pages)
                .build();
        responseObserver.onNext(list);
        responseObserver.onCompleted();
    }

    @Override
    public void getProduct(ProductId request, StreamObserver<Product> responseObserver) {
        Product p = db.get(request.getId());
        if (p == null) {
            responseObserver.onError(Status.NOT_FOUND
                    .withDescription("Product not found: " + request.getId())
                    .asRuntimeException());
            return;
        }
        responseObserver.onNext(p);
        responseObserver.onCompleted();
    }

    @Override
    public void createProduct(ProductInput request, StreamObserver<Product> responseObserver) {
        String id = UUID.randomUUID().toString();
        Product product = Product.newBuilder()
                .setId(id)
                .setName(request.getName())
                .setCategory(request.getCategory())
                .setPrice(request.getPrice())
                .setStock(request.getStock())
                .build();
        db.put(id, product);
        responseObserver.onNext(product);
        responseObserver.onCompleted();
    }

    @Override
    public void updateProduct(UpdateProductRequest request, StreamObserver<Product> responseObserver) {
        if (!db.containsKey(request.getId())) {
            responseObserver.onError(Status.NOT_FOUND
                    .withDescription("Product not found: " + request.getId())
                    .asRuntimeException());
            return;
        }
        ProductInput inp = request.getInput();
        Product product = Product.newBuilder()
                .setId(request.getId())
                .setName(inp.getName())
                .setCategory(inp.getCategory())
                .setPrice(inp.getPrice())
                .setStock(inp.getStock())
                .build();
        db.put(request.getId(), product);
        responseObserver.onNext(product);
        responseObserver.onCompleted();
    }

    @Override
    public void deleteProduct(ProductId request, StreamObserver<DeleteResponse> responseObserver) {
        if (!db.containsKey(request.getId())) {
            responseObserver.onError(Status.NOT_FOUND
                    .withDescription("Product not found: " + request.getId())
                    .asRuntimeException());
            return;
        }
        db.remove(request.getId());
        responseObserver.onNext(DeleteResponse.newBuilder()
                .setSuccess(true)
                .setMessage("Product " + request.getId() + " deleted")
                .build());
        responseObserver.onCompleted();
    }
}
