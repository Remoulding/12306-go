# 设置 proto 文件根目录和生成文件目录
PROTO_DIR="./proto"
OUT_DIR="./gen"

# 删除之前生成的文件
rm -rf "$OUT_DIR"

# 确保输出目录存在
mkdir -p "$OUT_DIR"

# 遍历 proto 目录中的所有 .proto 文件
find "$PROTO_DIR" -name "*.proto" | while read -r proto_file; do
    # 获取相对于 proto 根目录的路径`
    relative_path=$(dirname "${proto_file#$PROTO_DIR/}")
    # 设置生成文件的输出子目录
    output_subdir="$OUT_DIR/$relative_path"
    mkdir -p "$output_subdir"

    # 编译 proto 文件
    protoc --proto_path="$PROTO_DIR" \
           --go_out="$OUT_DIR" \
           --go-grpc_out="$OUT_DIR" \
           "$proto_file"
    echo "Compiled $proto_file to $output_subdir"
done

echo "All proto files compiled successfully!"