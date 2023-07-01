package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			sendData(fmt.Sprintf("data/%d.json", i))
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
}

func sendData(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	req, _ := http.NewRequest("POST", "http://localhost:5080/api/default/test2/_json", bytes.NewBuffer(content))
	req.SetBasicAuth("root@example.com", "Complexpass#123")
	resp, err := http.DefaultClient.Do(req)
	fmt.Println(resp, err)
}
