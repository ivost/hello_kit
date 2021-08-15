package main

import service "github.com/ivost/hello_kit/cmd/service"

/*
 http GET :8081/foo s=bar
*/
func main() {
	service.Run()
}
