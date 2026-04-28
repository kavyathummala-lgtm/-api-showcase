FROM python:3.12-slim
WORKDIR /app
COPY requirements.txt .
RUN pip install --no-cache-dir -r requirements.txt
COPY product.proto server.py client.py generate_proto.sh ./
RUN python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. product.proto
EXPOSE 50051
HEALTHCHECK --interval=30s --timeout=5s --start-period=15s --retries=3 \
  CMD python -c "import socket; s=socket.socket(); s.settimeout(3); s.connect(('localhost',50051)); s.close()" || exit 1
CMD ["python", "server.py"]
