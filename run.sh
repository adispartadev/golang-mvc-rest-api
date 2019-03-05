if [ "$1" = "m" ]; then
    export GO111MODULE=on
fi

if [ "$1" = "r" ]; then
    go run main.go
fi