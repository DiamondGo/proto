python -m grpc_tools.protoc -I. --python_out=service/python --grpc_python_out=service/python chat.proto
python -m grpc_tools.protoc -I. --go_out=. --go-grpc_out=. chat.proto
