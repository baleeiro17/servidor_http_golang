# servidor_http_golang

# Steps for execute code:

1- Download the source code:
```
https://github.com/baleeiro17/servidor_http_golang
```

2- Install the dependencies:
```
cd servidor_http_golang
go mod download
```

3- Build the binary:
```
cd cmd
go build server.go
```

4- Execute the binary:
```
./server
```

5- API working in:
```
http://localhost:8080
```
6- Running with docker:
```
docker build -t portal .
docker run -d -p 8080:8080 portal
```
