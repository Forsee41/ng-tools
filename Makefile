BIN=bin
CMD=cmd
GOFLAGS=
TARGETS=$(patsubst $(CMD)/%/main.go, $(BIN)/%, $(wildcard cmd/*/main.go))

$(BIN)/%: $(CMD)/%/main.go
	go build -o $@ $< $(GOFLAGS)

build: $(BIN) $(TARGETS)

$(BIN):
	mkdir -p bin
