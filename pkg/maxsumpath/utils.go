package maxsumpath

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func FetchTriangleData(url string) ([][]int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// I try too get it from github but
	// cause of Error parsing JSON:invalid character '<' looking for beginning of value T^T
	var data [][]int
	if err := json.Unmarshal(body, &data); err != nil {

		dir, _ := os.Getwd()
		hardFile, err := os.ReadFile(dir + "/hard.json")
		if err != nil {
			return nil, fmt.Errorf("failed to read hard.json file: %w", err)
		}
		if err := json.Unmarshal(hardFile, &data); err != nil {
			return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
		}
	}

	return data, nil
}

func FetchFormFile(filename string) ([][]int, error) {

	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %v", err)
	}

	var data [][]int
	if err := json.Unmarshal(byteValue, &data); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}
	return data, nil
}
