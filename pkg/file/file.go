package file

import (
    "bufio"
    "log"
		"os"
		"fmt"
)

type callbackType func(string)

func Read(path string, callback callbackType) {
	file, err := os.Open(path)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	counter := 0
	for scanner.Scan() {
		fmt.Print("\033[G\033[K") //limpa linha e retorna cursor para a posição 0
		fmt.Println("Line: ", counter)
		fmt.Print("\033[A")
		callback(scanner.Text())
		counter++
	}


	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}
