BIN=quiz
SRCS=main.go
VERB=0

ifeq ($(VERB),1)
	V= 
else
	V=@
endif

.PHONY: clean test

all: test ${BIN}

$(BIN): $(SRCS)
	$(V)go build -o $@ $<

test: $(SRCS)
	$(V)go test -v $<

run: $(BIN)
	$(V)./$(BIN)

clean:
	$(V)go clean
	$(V)rm -rf $(BIN)
