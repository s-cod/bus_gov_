package main

import (
	"bus_gov_go/pkg/bs"
	"bus_gov_go/pkg/fhd"
	"bus_gov_go/pkg/rd"
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {

	args := os.Args
	fmt.Println(args)

	if len(args) < 2 {
		fmt.Println("необходимо указать параметры для конвертации")
		fmt.Println("--fhd	для конвертации планов ФХД")
		fmt.Println("--bs	для конвертации Бюджетные средства")
		fmt.Println("--rd	для конвертации Результаты деятельности")
		return
	}
	var f func(string) error
	switch args[1] {
	case "--fhd":
		fmt.Println("Начинаем обрабатывать xls файлы Планов ФХД")
		f = fhd.ProcessFile
	case "--bs":
		fmt.Println("Начинаем обрабатывать xls файлы Бюджетные средства")
		f = bs.ProcessFile
	case "--rd":
		fmt.Println("Начинаем обрабатывать xls файлы Результаты деятельности")
		f = rd.ProcessFile
	default:
		fmt.Println("не верные параметры обработки")
		return
	}

	files, err := filepath.Glob("./in/*.xlsx")
	if err != nil {
		fmt.Println("Ошибка при чтении файлов:", err)
		fmt.Println("test")
		return
	}

	// Создаем канал для передачи файлов и канал для завершения работы работников
	fileChan := make(chan string, len(files))
	var wg sync.WaitGroup

	// Запускаем работников
	numWorkers := 4
	for range numWorkers {
		wg.Add(1)
		go worker(f, fileChan, &wg)
	}

	// Отправляем файлы в канал
	for _, file := range files {
		fileChan <- file
	}
	close(fileChan)

	// Ждем завершения всех работников
	wg.Wait()
	fmt.Println("Обработка завершена.")
}

func worker(f func(string) error, fileChan <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for file := range fileChan {
		fmt.Printf("Обработка файла: %s\n", file)
		if err := f(file); err != nil {
			fmt.Printf("Ошибка при обработке файла %s: %v\n", file, err)
		}
	}
}
