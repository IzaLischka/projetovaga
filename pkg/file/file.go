package file

import (
    "bufio"
    "log"
    "os"
)

type callbackType func(string)

func Read(path string, callback callbackType) {
	file, err := os.Open(path)
	if err != nil {
			log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		callback(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
}
