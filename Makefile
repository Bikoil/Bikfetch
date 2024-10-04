PREFIX ?= /usr

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

all:
	@echo Run \'sudo make install\' to install bikfetch

install:
	@install -Dm 755 bikfetch $(PREFIX)/bin/bikfetch

uninstall:
	@rm -f $(PREFIX)/bin/bikfetch

.PHONY: install uninstall all
