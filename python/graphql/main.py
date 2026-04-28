import strawberry
from strawberry.fastapi import GraphQLRouter
from fastapi import FastAPI
from typing import Optional
import uuid


@strawberry.type
class Product:
    id: str
    name: str
    category: str
    price: float
    stock: int


@strawberry.input
class ProductInput:
    name: str
    category: str
    price: float
    stock: int


@strawberry.type
class ProductPage:
    data: list[Product]
    page: int
    limit: int
    total: int
    pages: int


db: dict[str, Product] = {}


@strawberry.type
class Query:
    @strawberry.field
    def products(
        self,
        name: Optional[str] = None,
        category: Optional[str] = None,
        min_price: Optional[float] = None,
        max_price: Optional[float] = None,
        page: int = 1,
        limit: int = 10,
    ) -> ProductPage:
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

    @strawberry.field
    def product(self, id: str) -> Optional[Product]:
        return db.get(id)


@strawberry.type
class Mutation:
    @strawberry.mutation
    def create_product(self, input: ProductInput) -> Product:
        product_id = str(uuid.uuid4())
        product = Product(
            id=product_id,
            name=input.name,
            category=input.category,
            price=input.price,
            stock=input.stock,
        )
        db[product_id] = product
        return product

    @strawberry.mutation
    def update_product(self, id: str, input: ProductInput) -> Optional[Product]:
        if id not in db:
            return None
        product = Product(
            id=id,
            name=input.name,
            category=input.category,
            price=input.price,
            stock=input.stock,
        )
        db[id] = product
        return product

    @strawberry.mutation
    def delete_product(self, id: str) -> bool:
        if id not in db:
            return False
        del db[id]
        return True


schema = strawberry.Schema(query=Query, mutation=Mutation)
graphql_app = GraphQLRouter(schema)

app = FastAPI(title="Product GraphQL API", version="1.0.0")
app.include_router(graphql_app, prefix="/graphql")


@app.get("/health")
def health():
    return {"status": "ok", "service": "python-graphql"}
