package main
    
import (
	"strconv"
	"time"
	"fmt"
)

func main() {
	start := time.Now()
	// Create `sample.csv` in current directory
	csv, err := NewCsvWriter("sample.csv")
	if err != nil {
		panic("Could not open `sample.csv` for writing")
	}

	// Flush pending writes and close file upon exit of main()
	defer csv.Close()

	count := 1000000

	done := make(chan bool)

	for i := count; i > 0; i-- {
		go func(i int) {
			csv.Write([]string{strconv.Itoa(i), "bottles", "of", "beer"})
			done <- true
		}(i)
	}

	for i := 0; i < count; i++ {
		<-done
	}

	fmt.Print("Procession Time for %s records is %s", count, time.Since(start))
}