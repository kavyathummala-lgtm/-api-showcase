from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from typing import Optional
import uuid

app = FastAPI(title="Product REST API", version="1.0.0")


class Product(BaseModel):
    name: str
    category: str
    price: float
    stock: int


class ProductResponse(Product):
    id: str


class ProductPage(BaseModel):
    data: list[ProductResponse]
    page: int
    limit: int
    total: int
    pages: int


db: dict[str, ProductResponse] = {}


@app.get("/health")
def health():
    return {"status": "ok", "service": "python-rest"}


@app.get("/products", response_model=ProductPage)
def get_all_products(
    name: Optional[str] = None,
    category: Optional[str] = None,
    min_price: Optional[float] = None,
    max_price: Optional[float] = None,
    page: int = 1,
    limit: int = 10,
):
    results = list(db.values())
    if name:
        results = [p for p in results if name.lower() in p.name.lower()]
    if category:
        results = [p for p in results if p.category.lower() == category.lower()]
    if min_price is not None:
        results = [p for p in results if p.price >= min_price]
    if max_price is not None:
        results = [p for p in results if p.price <= max_price]
    total = len(results)
    pages = max(1, (total + limit - 1) // limit)
    start = (page - 1) * limit
    return ProductPage(data=results[start:start + limit], page=page, limit=limit, total=total, pages=pages)


@app.get("/products/{product_id}", response_model=ProductResponse)
def get_product(product_id: str):
    if product_id not in db:
        raise HTTPException(status_code=404, detail="Product not found")
    return db[product_id]


@app.post("/products", response_model=ProductResponse, status_code=201)
def create_product(product: Product):
    product_id = str(uuid.uuid4())
    new_product = ProductResponse(id=product_id, **product.model_dump())
    db[product_id] = new_product
    return new_product


@app.put("/products/{product_id}", response_model=ProductResponse)
def update_product(product_id: str, product: Product):
    if product_id not in db:
        raise HTTPException(status_code=404, detail="Product not found")
    updated = ProductResponse(id=product_id, **product.model_dump())
    db[product_id] = updated
    return updated


@app.delete("/products/{product_id}")
def delete_product(product_id: str):
    if product_id not in db:
        raise HTTPException(status_code=404, detail="Product not found")
    del db[product_id]
    return {"message": f"Product {product_id} deleted successfully"}
