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



