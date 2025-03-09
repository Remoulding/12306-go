OUTPUT_DIR="../idl-gen"

# -I 参数指定了 protoc 命令在搜索依赖文件时的搜索路径
protoc -I=proto \
--go_out="$OUTPUT_DIR" \
--go-grpc_out="$OUTPUT_DIR" \
--grpc-gateway_out="$OUTPUT_DIR" \
proto/**/*.proto

# 等待
echo "生成代码完成"
read -p "按任意键继续..." var