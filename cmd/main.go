package main

import (
	"fmt"

	"backend-challenge-7solutions/pkg/maxsumpath"
)

func main() {

	challengeOne() // 1. จงหาเส้นทางที่มีค่ามากที่สุด
	challengeTwo() // 2. จับฉันให้ได้สิ ซ้าย-ขวา-เท่ากับ

}

func challengeOne() {
	url := "https://github.com/7-solutions/backend-challenge/blob/main/files/hard.json"
	data, err := maxsumpath.FetchTriangleData(url)
	if err != nil {
		fmt.Printf("Error fetching data: %v\n", err)
		return
	}
	// root := maxsumpath.BuildTree(data)
	// maxSum := maxsumpath.MaxSumPathBFS(root) <-- BFS is can't handle the large data (out of memory)
	maxSum := maxsumpath.MaxSumPathDP(data)
	fmt.Println(maxSum)
}

func challengeTwo() {

}
