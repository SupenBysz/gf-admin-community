#!/bin/bash
echo "正在为以下平台构建: $PLATFORMS"

for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    OUTPUT_PATH="$OUTPUT_DIR/${GOOS}_${GOARCH}"
    BINARY_FILENAME="$BINARY_NAME"
    if [ "$GOOS" = "windows" ]; then
        BINARY_FILENAME="${BINARY_NAME}.exe"
    fi
    FINAL_OUTPUT="$OUTPUT_PATH/$BINARY_FILENAME"

    echo "正在构建 $GOOS/$GOARCH..."
    mkdir -p "$OUTPUT_PATH"
    env GOOS=$GOOS GOARCH=$GOARCH CGO_ENABLED=0 go build -ldflags="-s -w" -o "$FINAL_OUTPUT" "$SOURCE_FILE"
    if [ $? -ne 0 ]; then
        echo "构建失败: $GOOS/$GOARCH"
        exit 1
    fi
    echo "已构建: $FINAL_OUTPUT"

    echo "正在复制资源文件到 $OUTPUT_PATH..."
    for RESOURCE in $RESOURCES; do
        if [ -d "$RESOURCE" ] || [ -f "$RESOURCE" ]; then
            cp -R "$RESOURCE" "$OUTPUT_PATH/"
            echo "已复制: $RESOURCE"
        else
            echo "警告: 资源文件不存在: $RESOURCE"
        fi
    done

    echo "正在为 $GOOS/$GOARCH 创建发布包..."
    if [ "$GOOS" = "windows" ]; then
        ZIP_NAME="$OUTPUT_DIR/${BINARY_NAME}_${GOOS}_${GOARCH}.zip"
        echo "正在创建 ZIP 包: $ZIP_NAME"
        (cd "$OUTPUT_DIR" && zip -rq "../$ZIP_NAME" "${GOOS}_${GOARCH}")
        echo "已创建 ZIP 包: $ZIP_NAME"
    else
        TAR_NAME="$OUTPUT_DIR/${BINARY_NAME}_${GOOS}_${GOARCH}.tar.gz"
        echo "正在创建 TAR.GZ 包: $TAR_NAME"
        (cd "$OUTPUT_DIR" && tar -czf "../$TAR_NAME" "${GOOS}_${GOARCH}")
        echo "已创建 TAR.GZ 包: $TAR_NAME"
    fi
done

echo "所有平台构建完成。"

echo "正在清理临时文件..."
for PLATFORM in $PLATFORMS; do
    GOOS=${PLATFORM%/*}
    GOARCH=${PLATFORM#*/}
    BUILD_DIR="$OUTPUT_DIR/${GOOS}_${GOARCH}"
    if [ -d "$BUILD_DIR" ]; then
        echo "正在删除临时目录: $BUILD_DIR"
        rm -rf "$BUILD_DIR"
    fi
done
echo "清理完成。"
