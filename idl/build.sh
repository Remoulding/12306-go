# 定义proto文件所在目录
PROTO_DIR="./idl/proto"

# 定义生成的Go代码输出目录（这里是相对路径，实际结合后续操作对应到my-gen-code-repository仓库目录）
OUTPUT_DIR="./idl-gen"

## 循环处理每个proto文件
find "$PROTO_DIR" -name "*.proto" | while read -r proto_file; do
    # 生成Go代码
    protoc --proto_path="$PROTO_DIR" \
    --go_out="$OUTPUT_DIR" \
    --go-grpc_out="$OUTPUT_DIR" \
    "$proto_file"
done