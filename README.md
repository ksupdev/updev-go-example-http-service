# updev-go-example-http-service

## setup project 
1. Clone project from git ``git clone https://github.com/ksupdev/updev-go-example-http-service.git``

```powershell
git clone https://github.com/ksupdev/updev-go-example-http-service.git
Cloning into 'updev-go-example-http-service'...
remote: Enumerating objects: 4, done.
remote: Counting objects: 100% (4/4), done.
remote: Compressing objects: 100% (3/3), done.
remote: Total 4 (delta 0), reused 0 (delta 0), pack-reused 0
Unpacking objects: 100% (4/4), done.
```

2. setup module ``go mod init github.com/ksupdev/updev-go-example-http-service``
```powershell
% go mod init github.com/ksupdev/updev-go-example-http-service
go: creating new go.mod: module github.com/ksupdev/updev-go-example-http-service
---- file on current directory ----
ls
README.md       go.mod
```

พอทำการ run คำสั่ง go จะทำการสร้าง file go.mod

3. ทดสอบ project go โดยเราจะทำการสร้าง file ``main.go`` และให้แสดง ``hello world`` ออกมาที่ terminal

```go
[file:main.go]
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

```powershell
% go run main.go
hello world
---- build and excute ----
% go build main.go
% ./main
hello world
```

## Implement

### Create interface 
1. Create IContext interface : ใช้สำหรับ handle request and response context

```golang
type IContext interface {
	Log(message string) // log
	Param(name string) string // Get path param
	QueryParam(name string) string // Get query paran
	ReadInput() string // Read request body
	Response(responseCode int, responseData interface{}) //Response body
}
```

2. Create IMicroservice interface : ใช้สำหรับกำหนดการทำงานทั้งหมดของ microservice


```golang
type IMicroservice interface {
	Start() error
	Cleanup() error

	// HTTP Services
	GET(path string, h ServiceHandleFunc)
	POST(path string, h ServiceHandleFunc)
	PUT(path string, h ServiceHandleFunc)
	PATCH(path string, h ServiceHandleFunc)
	DELETE(path string, h ServiceHandleFunc)
}

//Function types
// define this type for implement function
type ServiceHandleFunc func(ctx IContext) error
```

> การที่เราทำการ Implement interface เป็นส่วนหนึ่งของ design pattern นั้นก็คือ ``design pattern proxy``



