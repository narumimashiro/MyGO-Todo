# makefile

.PHONY build run clean

build:
    go run build -o MyGOTodo

run:
    go run main.go

clean:
    rm -rf MyGoTodo