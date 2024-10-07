PREFIX ?= /usr
GO_FILES := bikfetch.go
BINARY := bikfetch

# Detect OS to set the correct PREFIX path
ifeq ($(shell uname), FreeBSD)
    PREFIX = /usr/local
endif

ifeq ($(shell uname), Darwin)
    PREFIX = /usr/local
endif

ifeq ($(shell uname), Linux)
    ifeq ($(shell uname -o), Android)
        PREFIX = /data/data/com.termux/files/usr
    endif
endif

all: $(BINARY)

$(BINARY): $(GO_FILES)
	go build -o $(BINARY) $(GO_FILES)
	@echo Built $(BINARY)

install: $(BINARY)
	@install -Dm 755 $(BINARY) $(PREFIX)/bin/$(BINARY)
	@rm ./bikfetch
	@echo Installed $(BINARY) to $(PREFIX)/bin, Thank you for installing!

uninstall:
	@rm -f $(PREFIX)/bin/$(BINARY)
	@echo Uninstalled $(BINARY) from $(PREFIX)/bin

clean:
	@rm -f $(BINARY)
	@echo Cleaned $(BINARY)

.PHONY: all install uninstall clean
