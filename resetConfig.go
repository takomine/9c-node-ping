package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func resetConfig() {
	// Download the JSON file
	resp, err := http.Get("https://release.nine-chronicles.com/9c-launcher-config.json")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Get the AppData directory
	appDataDir, err := os.UserConfigDir()
	if err != nil {
		fmt.Println("Error retrieving AppData directory:", err)
		return
	}

	// Create the Appdata/Nine Chronicles directory path
	roamingDir := filepath.Join(appDataDir, "Nine Chronicles")

	// Create the directory where the JSON file will be saved
	err = os.MkdirAll(roamingDir, 0755)
	if err != nil {
		panic(err)
	}

	// Create the file where the JSON data will be saved
	file, err := os.Create(filepath.Join(roamingDir, "config.json"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Copy the JSON data from the response body to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		panic(err)
	}
}
