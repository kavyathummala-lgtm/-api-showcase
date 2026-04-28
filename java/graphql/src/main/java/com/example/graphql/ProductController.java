package com.example.graphql;

import org.springframework.graphql.data.method.annotation.*;
import org.springframework.stereotype.Controller;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@Controller
public class ProductController {

    private final Map<String, Product> db = new ConcurrentHashMap<>();

    public record ProductPage(List<Product> data, int page, int limit, int total, int pages) {}

    @QueryMapping
    public ProductPage products(
            @Argument String name,
            @Argument String category,
            @Argument Double minPrice,
            @Argument Double maxPrice,
            @Argument Integer page,
            @Argument Integer limit) {
        int p = (page == null || page < 1) ? 1 : page;
        int l = (limit == null || limit < 1) ? 10 : limit;
        List<Product> filtered = db.values().stream()
                .filter(pr -> name == null || pr.getName().toLowerCase().contains(name.toLowerCase()))
                .filter(pr -> category == null || pr.getCategory().equalsIgnoreCase(category))
                .filter(pr -> minPrice == null || pr.getPrice() >= minPrice)
                .filter(pr -> maxPrice == null || pr.getPrice() <= maxPrice)
                .collect(java.util.stream.Collectors.toList());
        int total = filtered.size();
        int pages = Math.max(1, (int) Math.ceil((double) total / l));
        int start = Math.min((p - 1) * l, total);
        int end = Math.min(start + l, total);
        return new ProductPage(filtered.subList(start, end), p, l, total, pages);
    }

    @QueryMapping
    public Product product(@Argument String id) {
        return db.get(id);
    }

    @MutationMapping
    public Product createProduct(@Argument ProductInput input) {
        String id = UUID.randomUUID().toString();
        Product product = new Product(id, input.name(), input.category(), input.price(), input.stock());
        db.put(id, product);
        return product;
    }

    @MutationMapping
    public Product updateProduct(@Argument String id, @Argument ProductInput input) {
        if (!db.containsKey(id)) return null;
        Product product = new Product(id, input.name(), input.category(), input.price(), input.stock());
        db.put(id, product);
        return product;
    }

    @MutationMapping
    public boolean deleteProduct(@Argument String id) {
        if (!db.containsKey(id)) return false;
        db.remove(id);
        return true;
    }

    public record ProductInput(String name, String category, double price, int stock) {}
}
