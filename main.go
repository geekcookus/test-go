package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

func main() {
	//переменные
	var (
		elementMap      = make(map[string]int)
		elementMapMutex = sync.RWMutex{}
		wg              = sync.WaitGroup{}
	)
	//чтение папки
	files, err := ioutil.ReadDir("files/")
	if err != nil {
		panic(err)
	}
	//запуск горутины на один файл.
	for _, file := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			b, err := ioutil.ReadFile("files/" + file.Name())
			if err != nil {
				panic(err)
			}
			//счет символов
			for _, char := range string(b) {
				st := string(char)
				elementMapMutex.Lock()
				elementMap[st]++
				elementMapMutex.Unlock()
			}

		}()
	}
	wg.Wait()
	//вывод в консоль
	fmt.Printf("%s", elementMap)
	//go run -race main.go --> WARNING: DATA RACE
}
