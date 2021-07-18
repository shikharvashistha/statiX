// net/https package is used for network IO including http
// os package is an interface to OS functionality
// The design is unix like so it works on any platform
package main

import (
	"net/http"
	"os"
)

func main(){
	//In order to serve the static content in the current directory.
	dir, _ := os.Getwd()//Get the current working directory same as pwd(unix print working directory)
	//This function return two values, the first is a string and the second is an error
	//_ is error we don't care about it's like ignoring it.
	//Now we've the name of the directory we need to serve all the static flies in it.
	http.ListenAndServe(":8080", http.FileServer(http.Dir(dir+"/static")))
	//Listen and serve on this http port :8080 and serve the directory static in the current directory
	//open https://localhost:8080/ after running main.go 
}