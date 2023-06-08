package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

func overwriteJson(node string) {
	downloadJson()

	// Get the Temp directory
	tempDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("Error retrieving AppData directory:", err)
		return
	}

	// Create the Cache directory path
	cacheDir := filepath.Join(tempDir, "node-json")

	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		panic(err)
	}

	files, err := ioutil.ReadDir(cacheDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		dir := cacheDir // Replace with the directory path you want to search

		files, err := ioutil.ReadDir(dir)
		if err != nil {
			fmt.Println("Error reading directory:", err)
			return
		}

		// Sort files by modification time in descending order
		sort.Slice(files, func(i, j int) bool {
			return files[i].ModTime().After(files[j].ModTime())
		})

		// Find the most recent file
		var mostRecentFile os.FileInfo
		if len(files) > 0 {
			mostRecentFile = files[0]
			fmt.Println("Most recent file:", mostRecentFile.Name())
			// Process the most recent file if needed
			// ...
			fileName := cacheDir + "/" + mostRecentFile.Name()
			fmt.Println("Found file:", fileName)

			// Read the JSON data from the response body
			data, err := ioutil.ReadFile(fileName)
			if err != nil {
				panic(err)
			}

			// Unmarshal the JSON data into a map for RemoteNodeList
			var jsonNode map[string]interface{}
			err = json.Unmarshal(data, &jsonNode)
			if err != nil {
				panic(err)
			}

			// Get the slice value of the "nodeList" key and type assert it
			nodeList := jsonNode["RemoteNodeList"].([]interface{})

			// Array. node = RPC node
			nodeList[0] = node
			nodeList[1] = node
			nodeList[2] = node
			nodeList[3] = node
			nodeList[4] = node
			nodeList[5] = node
			nodeList[6] = node
			nodeList[7] = node
			nodeList[8] = node
			nodeList[9] = node

			jsonNode["RemoteNodeList"] = nodeList

			// Marshal the updated map back into JSON
			newDataNodeList, err := json.MarshalIndent(jsonNode, "", "  ")
			if err != nil {
				panic(err)
			}

			// Get the AppData directory
			appDataDir, err := os.UserConfigDir()
			if err != nil {
				fmt.Println("Error retrieving AppData directory:", err)
				return
			}

			// Create the Appdata/Nine Chronicles directory path
			roamingDir := filepath.Join(appDataDir, "Nine Chronicles")

			err = os.MkdirAll(roamingDir, 0755)
			if err != nil {
				panic(err)
			}

			// Save the updated JSON to a specific directory
			err = ioutil.WriteFile(filepath.Join(roamingDir, "config.json"), newDataNodeList, 0644)
			if err != nil {
				panic(err)
			}

		} else {
			fmt.Println("No files found in the directory.")
		}

	}

}

func downloadJson() {
	// Get the response from the URL
	resp, err := http.Get("https://release.nine-chronicles.com/9c-launcher-config.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Read the response body as a byte slice
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the Temp directory
	tempDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("Error retrieving AppData directory:", err)
		return
	}

	// Create the Cache directory path
	cacheDir := filepath.Join(tempDir, "node-json")

	err = os.MkdirAll(cacheDir, 0755)
	if err != nil {
		panic(err)
	}

	// Create a temporary file in the current directory with a prefix of "tmp"
	tmpFile, err := ioutil.TempFile(cacheDir, "*-tempjson")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get the name of the temporary file
	tmpName := tmpFile.Name()

	// Write the byte slice to the temporary file
	err = ioutil.WriteFile(tmpName, data, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Do something with the temporary file
	fmt.Println("Downloaded and saved JSON file as", tmpName)

	//ConfigVersion
	// Unmarshal the JSON data into a map for ConfigVersion
	var jsonConfig map[string]interface{}
	err = json.Unmarshal(data, &jsonConfig)
	if err != nil {
		panic(err)
	}

	// Increment the specific value
	jsonConfig["ConfigVersion"] = jsonConfig["ConfigVersion"].(float64) + 1

	// Marshal the updated map back into JSON
	newDataConfig, err := json.MarshalIndent(jsonConfig, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(filepath.Join("", tmpName), newDataConfig, 0644)
	if err != nil {
		panic(err)
	}

}

/* backup codes...


 */
