package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func checkPing() {
	// URLs to ping based from this "https://release.nine-chronicles.com/9c-launcher-config.json"
	urls := []string{
		"http://9c-main-rpc-1.nine-chronicles.com",
		"http://9c-main-rpc-2.nine-chronicles.com",
		"http://9c-main-rpc-3.nine-chronicles.com",
		"http://9c-main-rpc-4.nine-chronicles.com",
		"http://9c-main-rpc-5.nine-chronicles.com",
		"http://tky-nc-1.ninodes.com",
		"http://sgp-nc-1.ninodes.com",
		"http://nj-nc-1.ninodes.com",
		"http://la-nc-1.ninodes.com",
		"http://fra-nc-1.ninodes.com",
	}

	fmt.Println("Check response time (ping) of Nine Chronicles nodes")

	// Ping each URL
	for _, url := range urls {
		// Initialize a counter variable
		counter := 0

		// Loop until the counter reaches 3
		for counter < 3 {
			start := time.Now()
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println(err)
				// Increment the counter
				counter++
				continue
			}
			resp.Body.Close()
			elapsed := time.Since(start)
			convertSeconds := elapsed.Seconds()

			// remove "http://" and "com" from output
			wordsToRemove := []string{"http://", ".com"}
			cleanUrl := removeWords(url, wordsToRemove)

			// Print output
			fmt.Printf("%s responded in %.2f ms \n", cleanUrl, convertSeconds)

			// Reset the counter
			counter = 0

			// Break the loop
			break
		}
	}

	fmt.Println("Press ENTER to continue...")
	reader := bufio.NewReader(os.Stdin)
	_, _ = reader.ReadString('\n')
}
