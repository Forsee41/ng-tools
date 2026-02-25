BIN=bin
CMD=cmd
GOFLAGS=
TARGETS=$(patsubst $(CMD)/%/main.go, $(BIN)/%, $(wildcard cmd/*/main.go))

$(BIN)/%: $(CMD)/%/main.go
	go build -o $@ $< $(GOFLAGS)

build: $(BIN) $(TARGETS)

$(BIN):
	mkdir -p bin

.PHONY: watch/templ
watch/templ:
	templ generate --watch

.PHONY: watch/air_server
watch/air_server:
	air -c air_server.toml

.PHONY: watch/air_builder
watch/air_builder:
	air -c air_builder.toml

.PHONY: watch/vite
watch/vite:
	npm run dev

.PHONY: watch
watch:
	make -j 3 watch/templ watch/air_builder watch/air_server
