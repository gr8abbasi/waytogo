package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main()  {
	readerChannel := make(chan string)
	validatorChannel := make(chan string)
	mapperChannel := make(chan string)
	translatorChannel := make(chan string)
	//writerChannel := make(chan string)

	wg.Add(5)
	go reader(&wg, readerChannel)
	go validator(&wg, readerChannel, validatorChannel)
	go mapper(&wg, validatorChannel, mapperChannel)
	go translator(&wg, mapperChannel, translatorChannel)
	go writer(&wg, translatorChannel)

	wg.Wait()
	
}

func reader(wg *sync.WaitGroup, readerChannel chan string) {
	defer wg.Done()
	for i:=0; i < 10; i++ {
		go printData("Reader", i)
	}
}

func validator(wg *sync.WaitGroup, readerChannel chan string, validatorChannel chan string) {
	defer wg.Done()
	for i:=0; i < 10; i++ {
		go printData("Validator", i)
	}
}

func mapper(wg *sync.WaitGroup, validatorChannel chan string, mapperChannel chan string) {
	defer wg.Done()
	for i:=0; i < 10; i++ {
		go printData("Mapper", i)
	}
}

func translator(wg *sync.WaitGroup, mapperChannel chan string, transltranslatorChannelator chan string) {
	defer wg.Done()
	for i:=0; i < 10; i++ {
		go printData("Translator", i)
	}
}

func writer(wg *sync.WaitGroup, translatorChannel chan string) {
	defer wg.Done()
	for i:=0; i < 10; i++ {
		go printData("Writer", i)
	}
}

func printData(name string, i int) {
	fmt.Printf("%s # %d\n", name, i)
}

