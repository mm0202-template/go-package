FROM golang

COPY main.go main.go
COPY main_test.go main_test.go

ENTRYPOINT [ "go", "run", "main.go" ]
