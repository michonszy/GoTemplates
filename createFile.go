// hello.go

package main

import (
    "log"
    "os"
	"time"
	"strings"
	"os/user"
)

func main() {
	t := time.Now()
	
	formatedTime := t.Format("2006-01-02 15:04:05")
	nazwa2 := strings.Replace(string(formatedTime),".","-",3)

	user, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}

	username2 := user.Username
	username := strings.Replace(username2,"DOMAIN\\","",1)


	nazwa:= username+"-"+strings.Replace(string(nazwa2),":","-",3)+"-OUTPUT.txt"
    emptyFile, err := os.Create(nazwa)
    if err != nil {
        log.Fatal(err)
    }
    log.Println(emptyFile)
    emptyFile.Close()
}