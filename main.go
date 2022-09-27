package main

import (
	"io/ioutil"
	"fmt"
	"os"
	"bufio"
	"strings"
	"path/filepath"
	
)

func main() {
	var folder string
	if len(os.Args) > 0{
		folder = os.Args[1]
	}else{
		fmt.Println("First argument is the folder path")
	}
	files, err := ioutil.ReadDir(folder)
    if err != nil {
        fmt.Println(err)
    }

    for _, file := range files {
		extension := filepath.Ext(file.Name())
		if extension == ".yaml" {
			file, err := os.Open(file.Name())
			if err != nil {
				fmt.Println(err)
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			var kind string
			var count int = 0
			var statement string = ""
			for scanner.Scan() {
				splitted := strings.Split(scanner.Text(), ":")
				if splitted[0] == "kind" {
					kind = strings.TrimSpace(splitted[1])
				}
				if kind == "Service" {
					if strings.TrimSpace(splitted[0]) == "type"{
						statement = statement + " of type " + strings.TrimSpace(splitted[1])
					}
					if strings.TrimSpace(splitted[0]) == "name" {
						statement = "A service with name " + statement + strings.TrimSpace(splitted[1]) + " will be created"
					}
				}else if kind == "Deployment" {
					if strings.TrimSpace(splitted[0]) == "replicas"{
						statement = statement + " with number of replicas " + strings.TrimSpace(splitted[1])
					}
					if count < 1 {
						if strings.TrimSpace(splitted[0]) == "name" {
							count += 1
							statement = "A deployment with name " + statement + strings.TrimSpace(splitted[1]) + " will be created"
						}
					}
				}	
			}
			fmt.Println(statement)

			if err := scanner.Err(); err != nil {
				fmt.Println(err)
			}
		}
	}
    
}
