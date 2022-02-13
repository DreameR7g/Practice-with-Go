package main

import (
	"fmt"
	"log"
)

func main() {
	err := Testing()
	if err != nil {
		log.Printf("Testing through error %s", err)
	}
	collection := getCollection()
	for _, item := range collection {
		log.Println(item)
	}
	mapCollection := getMap()
	log.Println(mapCollection[50])
	log.Println(mapCollection[80])
	log.Println(mapCollection[100])
	log.Println(mapCollection[110])
}

func Testing() (err error) {
	err = fmt.Errorf("You screwed up")
	return
}

func getCollection() (collection []string) {
	collection = make([]string, 0)
	for i := 0; i < 10; i++ {
		item := fmt.Sprintf("Item #%d", i)
		collection = append(collection, item)
	}
	return
}

func getMap() (collection map[int]string) {
	collection = make(map[int]string)
	for i := 0; i <= 100; i = i + 10 {
		item := fmt.Sprintf("Item #%d", i)
		collection[i] = item
	}
	return
}
