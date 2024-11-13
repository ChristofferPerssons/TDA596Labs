package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Need to specify port")
		return
	}
	port := ":" + os.Args[1]
	fmt.Println(port)
	portInt, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("Port is not an integer")
		log.Fatal(err)
	}

	if portInt > 9999 && portInt < 0 {
		fmt.Println("Not a valid port number")
		log.Fatal(err)
	}

	//Testing get, post and unsupported methods
	//Testing file extensions and unsupported extensions
	//TODO: Testing badly formatted requests

	var wg sync.WaitGroup

	for i, m := range []string{"GET", "POST", "RANDOM"} {
		for k, e := range []string{".html", ".txt", ".gif", ".jpeg", ".jpg", ".css", "", ".", ".R"} {
			wg.Add(1)
			go func(method string, ext string, id int) {
				defer wg.Done()
				requestURL := "http://localhost" + port + "/hello" + ext
				req, err := http.NewRequest(method, requestURL, nil)
				if err != nil {
					fmt.Printf("client "+strconv.Itoa(id)+": could not create request: %s\n", err)
					os.Exit(1)
				}

				res, err := http.DefaultClient.Do(req)
				if err != nil {
					fmt.Printf("client "+strconv.Itoa(id)+": error making http request: %s\n", err)
					os.Exit(1)
				}
				fmt.Println("client " + strconv.Itoa(id) + ": sent " + method + " " + ext)
				fmt.Printf("client " + strconv.Itoa(id) + ": got response!\n")
				fmt.Printf("client "+strconv.Itoa(id)+": status code: %d\n", res.StatusCode)

				_, err = io.ReadAll(res.Body)
				if err != nil {
					fmt.Printf("client "+strconv.Itoa(id)+": could not read response body: %s\n", err)
					os.Exit(1)
				}
			}(m, e, i*10+k)
		}
	}
	wg.Wait()
}
