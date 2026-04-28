import grpc
import product_pb2
import product_pb2_grpc


def run(host="localhost", port=50051):
    with grpc.insecure_channel(f"{host}:{port}") as channel:
        stub = product_pb2_grpc.ProductServiceStub(channel)

        print("=== Create Product ===")
        created = stub.CreateProduct(product_pb2.ProductInput(
            name="Laptop", category="Electronics", price=999.99, stock=50
        ))
        print(f"Created: {created}")

        print("\n=== Get All Products ===")
        products = stub.GetProducts(product_pb2.ProductFilter())
        for p in products.products:
            print(f"  {p.id}: {p.name} - ${p.price}")

        print("\n=== Search Products (Electronics only) ===")
        filtered = stub.GetProducts(product_pb2.ProductFilter(category="Electronics"))
        for p in filtered.products:
            print(f"  {p.id}: {p.name} - ${p.price}")

        print("\n=== Get Single Product ===")
        fetched = stub.GetProduct(product_pb2.ProductId(id=created.id))
        print(f"Fetched: {fetched}")

        print("\n=== Update Product ===")
        updated = stub.UpdateProduct(product_pb2.UpdateProductRequest(
            id=created.id,
            input=product_pb2.ProductInput(
                name="Laptop Pro", category="Electronics", price=1299.99, stock=30
            )
        ))
        print(f"Updated: {updated}")

        print("\n=== Delete Product ===")
        result = stub.DeleteProduct(product_pb2.ProductId(id=created.id))
        print(f"Deleted: {result.success} - {result.message}")


if __name__ == "__main__":
    run()
