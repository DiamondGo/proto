
for PROTO in chat.proto blob.proto
do
    echo "generating code for $PROTO"
    python -m grpc_tools.protoc -I. --python_out=service/python --grpc_python_out=service/python $PROTO
    python -m grpc_tools.protoc -I. --go_out=. --go-grpc_out=. $PROTO
done