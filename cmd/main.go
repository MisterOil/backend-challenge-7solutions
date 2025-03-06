package main

import (
	"fmt"
	"log"

	"backend-challenge-7solutions/pkg/leftrightequal"
	"backend-challenge-7solutions/pkg/maxsumpath"
	"backend-challenge-7solutions/pkg/piefiredire"

	"github.com/gin-gonic/gin"
)

func main() {

	challengeOne()   // 1. จงหาเส้นทางที่มีค่ามากที่สุด
	challengeTwo()   // 2. จับฉันให้ได้สิ ซ้าย-ขวา-เท่ากับ
	challengeThree() // 3.พาย ไฟ ได - Pie Fire Dire

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
	var encoded string
	// fmt.Scanln(&encoded);
	encoded = "LLRR="

	result, err := leftrightequal.MinDigitSequence(encoded)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}

func challengeThree() {
	router := gin.Default()
	router.GET("/", piefiredire.BeefSummaryHandler)

	if err := router.Run(":3002"); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
