package multithreadingSort

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"sync"
)

type Doctor struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phoneNumber" validate:"required"`
	Email       string `json:"email" validate:"required"`
	UserAgent   string `json:"userAgent" validate:"required"`
	HexColor    string `json:"hexColor" validate:"required"`
}

func SortArray(goroutines int) ([]Doctor, error) {

	data, err := readFile()
	if err != nil {
		panic("Incorrect work reading function")
	}

	chunkSize := len(data) / goroutines
	if len(data)&goroutines != 0 {
		chunkSize++
	}

	var wg sync.WaitGroup

	resultChanel := make(chan []Doctor)

	for i := 0; i < goroutines; i++ {
		start := i * chunkSize
		end := start + chunkSize

		if end > len(data) {
			end = len(data)
		}

		wg.Add(1)
		go sortPart(data[start:end], resultChanel, &wg)
	}

	go func() {
		wg.Wait()
		close(resultChanel)
	}()

	var sortedParts [][]Doctor
	for part := range resultChanel {
		sortedParts = append(sortedParts, part)
	}

	sortedDoctors := mergeSortedDoctors(sortedParts)

	return sortedDoctors, nil

}

func mergeSortedDoctors(parts [][]Doctor) []Doctor {
	var merged []Doctor
	for _, part := range parts {
		merged = append(merged, part...)
	}

	sort.Slice(merged, func(i, j int) bool {
		return merged[i].Name < merged[j].Name
	})

	return merged
}

func sortPart(arr []Doctor, result chan<- []Doctor, wg *sync.WaitGroup) {
	defer wg.Done()

	sort.Slice(arr, func(i, j int) bool {
		return arr[i].Name < arr[j].Name
	})

	result <- arr
}

func readFile() ([]Doctor, error) {

	byteValue, err := os.ReadFile("./500kb.json")
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err.Error())
	}

	var doctors []Doctor

	err = json.Unmarshal(byteValue, &doctors)
	if err != nil {
		return nil, err
	}

	return doctors, nil
}
