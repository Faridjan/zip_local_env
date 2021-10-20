package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"ziper_local_env/utils"
)

var allENV string

const PATH = "."

const delimiter = "("
const delimiter2 = ")\n"

const envName = "result"

func main() {

	log.Println(utils.Yellow("Inter the encryption KEY:"))
	var PHRASE string
	_, _ = fmt.Scan(&PHRASE)

	log.Println(utils.Yellow("Zip or unzip? (1, 2)"))
	var scanln string
	_, _ = fmt.Scan(&scanln)

	for scanln != "1" && scanln != "2" {
		log.Println(utils.Yellow("Inter '1' or '2'"))
		_, _ = fmt.Scan(&scanln)
	}

	if scanln == "1" {
		zip(PHRASE)
	} else if scanln == "2" {
		unzip(PHRASE)
	}
}

func zip(PHRASE string) {
	// getting all services...
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() {
			continue
		}

		// Read .env in service
		envContent, err := utils.ReadFile("./" + f.Name() + "/.env")
		if os.IsNotExist(err) {
			log.Println(utils.Yellow(f.Name()), "- doesn't have .evn file.")
			continue
		}

		allENV += delimiter + f.Name() + delimiter2 + envContent + "\n"
	}

	// write file
	if err := utils.CreateFile(PATH, envName, string(utils.Encrypt([]byte(allENV), PHRASE))); err != nil {
		log.Fatal(err.Error())
	}
}

func unzip(PHRASE string) {
	file, err := utils.ReadFile(PATH + "/" + envName)
	if err != nil {
		return
	}

	decryptedENV := string(utils.Decrypt([]byte(file), PHRASE))
	serviceEnv := strings.Split(decryptedENV, delimiter)

	for _, chunkEnvService := range serviceEnv[1:] {
		serviceWithEnv := strings.Split(chunkEnvService, delimiter2)

		if len(serviceEnv) < 2 {
			continue
		}

		serviceName := serviceWithEnv[0]
		envContent := serviceWithEnv[1]

		if err := utils.CreateFile(PATH, serviceName+"/.env", envContent); err != nil {
			return
		}
	}
}
