#!/bin/bash
set -e

echo "🚀 Starting cross-compilation for bot..."
echo ""

# --- 配置 ---
LDFLAGS="-s -w"
SOURCE_FILE="main.go"
OUTPUT_BASE="skygo"
OUTPUT_PATH="dist"
TARGETS=(
    "windows/amd64"
    "darwin/amd64"
    "darwin/arm64"
    "linux/amd64"
    "linux/arm64"
)
# ---

mkdir -p "$OUTPUT_PATH"

# --- 循环构建 ---
for target in "${TARGETS[@]}"; do
    GOOS=$(echo "$target" | cut -d'/' -f1)
    GOARCH=$(echo "$target" | cut -d'/' -f2)

    # 确定输出文件名
    OUTPUT_NAME="${OUTPUT_PATH}/${OUTPUT_BASE}_${GOOS}_${GOARCH}"
    # 如果是为Windows构建，添加.exe后缀
    if [ "$GOOS" = "windows" ]; then
        OUTPUT_NAME+=".exe"
    fi

    echo "----------------------------------------"
    echo "Building for ${GOOS}/${GOARCH}..."

    # 执行构建命令，将GOOS和GOARCH作为该命令的临时环境变量
    GOOS=$GOOS GOARCH=$GOARCH go build -ldflags="$LDFLAGS" -o "$OUTPUT_NAME" "$SOURCE_FILE"

    echo "✅ Successfully built: ${OUTPUT_NAME}"
done

echo "----------------------------------------"
echo "🎉 All targets built successfully!"
echo "Find your binaries in the '${OUTPUT_PATH}' directory."