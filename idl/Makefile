
OUTPUT_DIR="../idl-gen"

all:
	protoc -I=proto \
	--go_out="$OUTPUT_DIR" \
	--go-grpc_out="$OUTPUT_DIR" \
	--grpc-gateway_out="$OUTPUT_DIR" \
	proto/**/*.proto

user:
	protoc -I=proto \
	--go_out="$OUTPUT_DIR" \
	--go-grpc_out="$OUTPUT_DIR" \
	--grpc-gateway_out="$OUTPUT_DIR" \
	proto/user/*.proto

order:
	protoc -I=proto \
	--go_out="$OUTPUT_DIR" \
	--go-grpc_out="$OUTPUT_DIR" \
	--grpc-gateway_out="$OUTPUT_DIR" \
	proto/order/*.proto

ticket:
	protoc -I=proto \
	--go_out="$OUTPUT_DIR" \
	--go-grpc_out="$OUTPUT_DIR" \
	--grpc-gateway_out="$OUTPUT_DIR" \
	proto/ticket/*.proto




