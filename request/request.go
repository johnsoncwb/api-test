package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		for j := 1; j < 11; j++ {
			callService(j)
			time.Sleep(1 * time.Second)
		}
	}
}

func callService(i int) {
	client := &http.Client{
		Timeout: 20 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://192.168.18.8:8090/%d", i), nil)
	if err != nil {
		log.Fatal(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
}
