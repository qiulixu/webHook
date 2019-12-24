build:
	GOOS=linux GOARCH=amd64 go build -tags netgo -o greeter-srv  main.go
run:
	nohup ./webhook &&