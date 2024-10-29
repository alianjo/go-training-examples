package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

var done = make(chan bool)

func IPGetter(counts *int) (resp *http.Response, err error) {
	if *counts == 5 {
		done <- true
	}
	url := "https://ifconfig.io/ip"
	resp, err = http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("An error occured while getting the public ip %s", err)
	}
	*counts++
	return resp, nil
}

func main() {
	counts := 0
	ticker := time.NewTicker(time.Second * 2)

	go func() {
		for {
			select {
			case <-done:
				fmt.Println("YOU are Done")
				return
			case <-ticker.C:
				resp, err := IPGetter(&counts)
				defer resp.Body.Close()
				body, err := io.ReadAll(resp.Body)
				if err != nil {
					fmt.Println(err)
				}
				ip := strings.TrimSpace(string(body))
				fmt.Println("YOUR IP : ", ip)
			}
		}
	}()
	time.Sleep(time.Second * 5)
	<-done
	fmt.Println("We are Finished")
}
