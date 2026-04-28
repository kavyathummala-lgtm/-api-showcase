package com.example.rest;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.*;

import java.util.*;
import java.util.concurrent.ConcurrentHashMap;

@RestController
@RequestMapping("/products")
public class ProductController {

    private final Map<String, Product> db = new ConcurrentHashMap<>();

    @GetMapping
    public Map<String, Object> getAll(
            @RequestParam(required = false) String name,
            @RequestParam(required = false) String category,
            @RequestParam(required = false) Double minPrice,
            @RequestParam(required = false) Double maxPrice,
            @RequestParam(defaultValue = "1") int page,
            @RequestParam(defaultValue = "10") int limit) {
        List<Product> filtered = db.values().stream()
                .filter(p -> name == null || p.getName().toLowerCase().contains(name.toLowerCase()))
                .filter(p -> category == null || p.getCategory().equalsIgnoreCase(category))
                .filter(p -> minPrice == null || p.getPrice() >= minPrice)
                .filter(p -> maxPrice == null || p.getPrice() <= maxPrice)
                .collect(java.util.stream.Collectors.toList());
        int total = filtered.size();
        int pages = Math.max(1, (int) Math.ceil((double) total / limit));
        int start = Math.min((page - 1) * limit, total);
        int end = Math.min(start + limit, total);
        return Map.of("data", filtered.subList(start, end), "page", page, "limit", limit, "total", total, "pages", pages);
    }

    @GetMapping("/{id}")
    public ResponseEntity<Product> getById(@PathVariable String id) {
        Product p = db.get(id);
        return p != null ? ResponseEntity.ok(p) : ResponseEntity.notFound().build();
    }

    @PostMapping
    public ResponseEntity<Product> create(@RequestBody Product input) {
        String id = UUID.randomUUID().toString();
        Product product = new Product(id, input.getName(), input.getCategory(), input.getPrice(), input.getStock());
        db.put(id, product);
        return ResponseEntity.status(HttpStatus.CREATED).body(product);
    }

    @PutMapping("/{id}")
    public ResponseEntity<Product> update(@PathVariable String id, @RequestBody Product input) {
        if (!db.containsKey(id)) return ResponseEntity.notFound().build();
        Product product = new Product(id, input.getName(), input.getCategory(), input.getPrice(), input.getStock());
        db.put(id, product);
        return ResponseEntity.ok(product);
    }

    @DeleteMapping("/{id}")
    public ResponseEntity<Map<String, String>> delete(@PathVariable String id) {
        if (!db.containsKey(id)) return ResponseEntity.notFound().build();
        db.remove(id);
        return ResponseEntity.ok(Map.of("message", "Product " + id + " deleted"));
    }
}
