# List of SQL source files
SRC = $(wildcard *.sql) run.go

# File to track the last migration run
TIMESTAMP = .last_migration

# Default target
all: migration

# Target to run the Go script if any SQL file is newer than the timestamp
migration: $(TIMESTAMP)

# Update the timestamp after running the Go script
$(TIMESTAMP): $(SRC)
	touch $(TIMESTAMP)
	go run run.go

# Phony targets to avoid conflicts with file names
.PHONY: all migration
