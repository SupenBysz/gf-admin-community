ROOT_DIR    = $(shell pwd)
NAMESPACE   = "github.com/SupenBysz/gf-admin-community"

# 安装/更新最新的 CLI 工具
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf -fr \
	ln -s /usr/local/bin/gf /bin/gf


# 检查并安装 CLI 工具
.PHONY: cli.install
cli.install:
	@set -e; \
	gf -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "GoFame CLI 未安装，开始自动安装..."; \
		make cli; \
	fi;


# 生成 DAO/DO/Entity 的 Go 文件
.PHONY: dao
dao: cli.install
#	@gf gen dao -p sys_model -o sys_do -e sys_entity -d sys_dao -t1 hack/tpls/dao_template_dao.tpl -t2 hack/tpls/dao_internal_template.tpl -t3 hack/tpls/dao_template_do.tpl -t4 hack/tpls/dao_template_entity.tpl
	@gf gen dao


# 生成 sys_service 的 Go 文件
.PHONY: service
service: cli.install
#	@gf gen service
	@gf gen service -d ./sys_service

# 构建相关变量
PLATFORMS := linux/amd64 windows/amd64 darwin/amd64 darwin/arm64
SOURCE_FILE := example/main.go
BINARY_NAME := main
RESOURCES := example/resource example/i18n example/manifest example/.env.example

# 清理构建目录
.PHONY: build.clean
build.clean:
	@echo "正在清理旧的构建目录..."
	@rm -rf ./build
	@mkdir -p ./build

# 准备构建环境
.PHONY: build.prepare
build.prepare:
	@echo "正在准备构建环境..."
	@echo "将构建以下平台: $(PLATFORMS)"

# 编译特定平台的二进制文件
.PHONY: build.compile
build.compile:
	@for PLATFORM in $(PLATFORMS); do \
		OS=$${PLATFORM%/*}; \
		ARCH=$${PLATFORM#*/}; \
		OUTPUT_DIR="./build/$${OS}_$${ARCH}"; \
		OUTPUT_BINARY="$${OUTPUT_DIR}/$${BINARY_NAME}"; \
		if [ "$${OS}" = "windows" ]; then \
			OUTPUT_BINARY="$${OUTPUT_BINARY}.exe"; \
		fi; \
		echo "--> 正在构建 $${OS}/$${ARCH}..."; \
		mkdir -p "$${OUTPUT_DIR}"; \
		env GOOS=$${OS} GOARCH=$${ARCH} CGO_ENABLED=0 go build -v -ldflags="-s -w" -o "$${OUTPUT_BINARY}" $(SOURCE_FILE); \
		if [ $$? -ne 0 ]; then \
			echo "构建失败: $${OS}/$${ARCH}"; \
			exit 1; \
		fi; \
	done

# 复制资源文件到构建目录
.PHONY: build.resource
build.resource:
	@for PLATFORM in $(PLATFORMS); do \
		OS=$${PLATFORM%/*}; \
		ARCH=$${PLATFORM#*/}; \
		OUTPUT_DIR="./build/$${OS}_$${ARCH}"; \
		echo "--> 正在复制资源文件到 $${OUTPUT_DIR}..."; \
		cp -R $(RESOURCES) "$${OUTPUT_DIR}/"; \
	done

# 为不同平台打包发布文件
.PHONY: build.package
build.package:
	@for PLATFORM in $(PLATFORMS); do \
		OS=$${PLATFORM%/*}; \
		ARCH=$${PLATFORM#*/}; \
		echo "--> 正在为 $${OS}/$${ARCH} 创建压缩包..."; \
		if [ "$${OS}" = "windows" ]; then \
			ZIP_NAME="./build/gf-admin-community_$${OS}_$${ARCH}.zip"; \
			ARCHIVE_DIR="$${OS}_$${ARCH}"; \
			echo "----> 正在创建 $${ZIP_NAME}..."; \
			(cd ./build && zip -rq "../$${ZIP_NAME}" "$${ARCHIVE_DIR}") || exit 1; \
			echo "----> 已创建 $${ZIP_NAME}"; \
		else \
			TAR_NAME="./build/gf-admin-community_$${OS}_$${ARCH}.tar.gz"; \
			ARCHIVE_DIR="$${OS}_$${ARCH}"; \
			echo "----> 正在创建 $${TAR_NAME}..."; \
			(cd ./build && tar -czvf "../$${TAR_NAME}" "$${ARCHIVE_DIR}") || exit 1; \
			echo "----> 已创建 $${TAR_NAME}"; \
		fi; \
	done

# 清理临时构建文件
.PHONY: build.cleanup
build.cleanup:
	@echo "正在清理临时构建文件..."
	@for PLATFORM in $(PLATFORMS); do \
		OS=$${PLATFORM%/*}; \
		ARCH=$${PLATFORM#*/}; \
		BUILD_DIR="./build/$${OS}_$${ARCH}"; \
		echo "--> 正在删除临时目录: $${BUILD_DIR}..."; \
		rm -rf "$${BUILD_DIR}"; \
	done
	@echo "清理完成。"

# 主构建目标
.PHONY: build
build: cli.install build.clean build.prepare build.compile build.resource build.package build.cleanup
	@echo "所有平台构建完成。"



