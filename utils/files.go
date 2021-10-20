package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Alphabetic []string

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
	si := list[i]
	sj := list[j]
	siLower := strings.ToLower(si)
	sjLower := strings.ToLower(sj)
	if siLower == sjLower {
		return si < sj
	}
	return siLower < sjLower
}

func GetFiles(dir string) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, fmt.Errorf("Open directory error=%v\n", err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(file)

	list, _ := file.Readdirnames(0) // 0 to read all files and folders
	sort.Sort(Alphabetic(list))
	return list, nil
}

func ReadFile(file string) (string, error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func CreateFile(path, fileName string, message string) error {
	// Creating dir...
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err = os.MkdirAll(path, 0777)
		if err != nil {
			return err
		}
	}

	// Creating file...
	file, err := os.Create(path + "/" + fileName) // create file
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}

	_, err = file.WriteString(message) // writing to file
	if err != nil {
		err = file.Close() // close file
		if err != nil {
			return err
		}
		return err
	}

	return file.Close() // close file
}
