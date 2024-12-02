package main

import (
	"fmt"
	"net/http"
	"time"
)

const monitoringTimes = 4
const delay = 2

func main() {

	startMonitor()
}

func startMonitor() {
	sites := []string{"https://httpstat.us/random/200,201,500-504", "https://www.alura.com.br", "https://www.caelum.com.br"}
	fmt.Println("----------------------------------")

	for j := 1; j < monitoringTimes; j++ {
		fmt.Println(" ")
		fmt.Println("*** Testing", j, "/3 times. ***")
		fmt.Println(" ")
		for i, site := range sites {
			fmt.Println("Position:", i, "// Site:", site)
			testSite(site)
			fmt.Println("----------------------------------")
		}
		time.Sleep(delay * time.Second)
	}

}

func testSite(site string) {

	resp, _ := http.Get(site)
	respStatusCode := resp.StatusCode

	if respStatusCode == 200 {
		fmt.Println("Website loaded successfully (StatusCode):", respStatusCode)
	} else {
		fmt.Println("Website error (StatusCode):", respStatusCode)

	}

}
