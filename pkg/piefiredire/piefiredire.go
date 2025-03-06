package piefiredire

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const baconIpsumURL = "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"

func BeefSummaryHandler(c *gin.Context) {
	resp, err := http.Get(baconIpsumURL)
	if err != nil {
		log.Printf("Error fetching text: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch text"})
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch text"})
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response body: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to read text"})
		return
	}

	summary := countMeatTypes(string(body))

	c.JSON(http.StatusOK, gin.H{"beef": summary})
}
func countMeatTypes(text string) map[string]int {
	text = strings.ToLower(text)

	// Split the text into words based on whitespace.
	words := strings.Fields(text)
	counts := make(map[string]int)
	/*
		{word: count,....}
	*/
	for _, word := range words {
		// Trim punctuation characters.
		word = strings.Trim(word, ",.!?;:")
		if word == "" {
			continue
		}
		counts[word]++
	}
	return counts
}
