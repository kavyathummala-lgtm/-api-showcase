import grpc
from concurrent import futures
import uuid
import product_pb2
import product_pb2_grpc


db: dict[str, dict] = {}


class ProductServiceServicer(product_pb2_grpc.ProductServiceServicer):

    def GetProducts(self, request, context):
        results = []
        for p in db.values():
            if request.name and request.name.lower() not in p["name"].lower():
                continue
            if request.category and request.category.lower() != p["category"].lower():
                continue
            if request.min_price > 0 and p["price"] < request.min_price:
                continue
            if request.max_price > 0 and p["price"] > request.max_price:
                continue
            results.append(product_pb2.Product(
                id=p["id"], name=p["name"], category=p["category"],
                price=p["price"], stock=p["stock"]
            ))
        total = len(results)
        page = max(1, request.page if request.page > 0 else 1)
        limit = max(1, request.limit if request.limit > 0 else 10)
        pages = max(1, (total + limit - 1) // limit)
        start = (page - 1) * limit
        return product_pb2.ProductList(products=results[start:start + limit], total=total, page=page, pages=pages)

    def GetProduct(self, request, context):
        p = db.get(request.id)
        if not p:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Product not found")
            return product_pb2.Product()
        return product_pb2.Product(
            id=p["id"], name=p["name"], category=p["category"],
            price=p["price"], stock=p["stock"]
        )

    def CreateProduct(self, request, context):
        product_id = str(uuid.uuid4())
        db[product_id] = {
            "id": product_id, "name": request.name,
            "category": request.category, "price": request.price,
            "stock": request.stock,
        }
        return product_pb2.Product(
            id=product_id, name=request.name, category=request.category,
            price=request.price, stock=request.stock
        )

    def UpdateProduct(self, request, context):
        if request.id not in db:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Product not found")
            return product_pb2.Product()
        inp = request.input
        db[request.id] = {
            "id": request.id, "name": inp.name,
            "category": inp.category, "price": inp.price,
            "stock": inp.stock,
        }
        return product_pb2.Product(
            id=request.id, name=inp.name, category=inp.category,
            price=inp.price, stock=inp.stock
        )

    def DeleteProduct(self, request, context):
        if request.id not in db:
            context.set_code(grpc.StatusCode.NOT_FOUND)
            context.set_details("Product not found")
            return product_pb2.DeleteResponse(success=False, message="Product not found")
        del db[request.id]
        return product_pb2.DeleteResponse(success=True, message=f"Product {request.id} deleted")


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    product_pb2_grpc.add_ProductServiceServicer_to_server(ProductServiceServicer(), server)
    server.add_insecure_port("[::]:50051")
    server.start()
    print("gRPC server running on port 50051")
    server.wait_for_termination()


if __name__ == "__main__":
    serve()
