package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	followers, err := getLinks("followers.json")
	if err != nil {
		fmt.Println("can not read followers")
	}

	followings, err := getLinks("following.json")
	if err != nil {
		fmt.Println("can not read followings")
	}

	var result []string
	for k := range followings {
		if ok := followers[k]; !ok {
			result = append(result, k)
		}
	}

	saveToFile("output.html", result)

}

func getLinks(path string) (map[string]bool, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open the file: %v", err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, fmt.Errorf("failed to read the file: %v", err)
	}

	var result []map[string]interface{}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}
	links := make(map[string]bool)
	for _, item := range result {
		profile_url := (((item["string_list_data"].([]interface{}))[0]).(map[string]interface{})["href"]).(string)
		links[profile_url] = true
	}

	return links, nil
}

func saveToFile(path string, result []string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	for _, str := range result {
		_, err := file.WriteString("<a target='_blank' href='" + str + "'>" + str + "</a> </br>" + "\n")
		if err != nil {
			return fmt.Errorf("error writing to file: %v", err)
		}
	}
	fmt.Println("String array saved to file successfully.")
	return nil
}
