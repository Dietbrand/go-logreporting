package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	//Nginx uses log_format to format logs. We would need the $status from the nginx.conf file, delimited with ;

	nginxConf, err := os.Open("nginx.conf")
	check(err)
	fmt.Printf("opening %s", nginxConf.Name())
	reader := bufio.NewReader(nginxConf)
	for {
		line, err := reader.ReadString(';')
		check(err)
		if strings.Contains(line, "log_format") {
			strArray := strings.Fields(line)
			fmt.Println(strArray)
		}
	}
	defer nginxConf.Close()
}
