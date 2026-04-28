#!/bin/bash
# Run this once to generate gRPC stubs from the proto file
python -m grpc_tools.protoc \
  -I. \
  --python_out=. \
  --grpc_python_out=. \
  product.proto
echo "Generated product_pb2.py and product_pb2_grpc.py"
