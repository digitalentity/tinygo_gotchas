BUILD  := build
BINARY := demo

TINYGO_VERSION := 0.41.1
TOOLCHAIN_DIR  := toolchain
TINYGO         := $(TOOLCHAIN_DIR)/tinygo/bin/tinygo

GOOPTS := 	-size=short \
			-opt=1 \
			-panic=trap \
			-gc=conservative \
			-scheduler=tasks \
			-print-allocs=.

GOTARGET := rp2040

# ── toolchain ──────────────────────────────────────────────────────────────────

test-toolchain:
	@if [ -f $(TINYGO) ]; then \
		$(TINYGO) version; \
	else \
		echo "TinyGo toolchain not found at $(TINYGO). Run 'make install-toolchain' first."; \
		exit 1; \
	fi

install-toolchain:
	mkdir -p $(TOOLCHAIN_DIR)
	curl -L https://github.com/tinygo-org/tinygo/releases/download/v$(TINYGO_VERSION)/tinygo$(TINYGO_VERSION).linux-amd64.tar.gz | tar -xzC $(TOOLCHAIN_DIR)

# ── demo applications ──────────────────────────────────────────────────────────

build-alloc-woes: test-toolchain | $(BUILD)
	$(TINYGO) build $(GOOPTS) -target=$(TARGET) -o $(BUILD)/$(BINARY)-$*.hex ./src/alloc_woes

test-alloc-woes:
	$(TINYGO) test -v ./src/alloc_woes/


# ── misc ───────────────────────────────────────────────────────────────────────

$(BUILD):
	mkdir -p $@

clean:
	rm -rf $(BUILD)
