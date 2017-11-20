package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	buckets := make([][]string, 12)

	scanner := bufio.NewScanner(res.Body)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word := scanner.Text()
		n := HashBucket(word, 12)
		buckets[n] = append(buckets[n], word)
	}

	for i := 0; i < 12; i++ {
		fmt.Println(len(buckets[i]))
	}

	// fmt.Println(buckets[5])

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}

// HashBucket algorithm
func HashBucket(word string, size int) int {
	sum := 0
	for _, v := range word {
		sum += int(v)
	}
	return sum % size
}
