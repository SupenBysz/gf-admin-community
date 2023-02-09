ROOT_DIR    = $(shell pwd)
NAMESPACE   = "github.com/SupenBysz/gf-admin-community"

# Install/Update to the latest CLI tool.
.PHONY: cli
cli:
	@set -e; \
	wget -O gf https://github.com/gogf/gf/releases/latest/download/gf_$(shell go env GOOS)_$(shell go env GOARCH) && \
	chmod +x gf && \
	./gf install -y && \
	rm ./gf -fr \
	ln -s /usr/local/bin/gf /bin/gf


# Check and install CLI tool.
.PHONY: cli.install
cli.install:
	@set -e; \
	gf -v > /dev/null 2>&1 || if [[ "$?" -ne "0" ]]; then \
  		echo "GoFame CLI is not installed, start proceeding auto installation..."; \
		make cli; \
	fi;


# Generate Go files for DAO/DO/Entity.
.PHONY: dao
dao: cli.install
	@gf gen dao -p sys_model -o sys_do -e sys_entity -d sys_dao -t1 hack/tpls/dao_template_dao.tpl -t2 hack/tpls/dao_internal_template.tpl -t3 hack/tpls/dao_template_do.tpl -t4 hack/tpls/dao_template_entity.tpl


# Generate Go files for sys_service.
.PHONY: service
service: cli.install
	@gf gen service -d ./sys_service
