#!/bin/bash

help() {
    echo "type -b for backend, -f for frontend, -a for admin or nothing for all"
    echo "All isnt recommended as it opens all in one term"
}

backend() {
    cd backend
    go run main.go
    cd ..
}

case "$1" in
"-h")
    help
    ;;
"-b")
    backend
    ;;
"-f")
    ;;
"-a")
    admin
    ;;
*)
    admin &
    backend
    ;;

esac
