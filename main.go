package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	smbiosdata "github.com/grep-michael/SMBIOS_parser/SMBiosLib/SMBiosData"
)

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.SetOutput(file)

	//flag handling
	desired_struct := flag.Int("struct", 4, "Structureto print")
	test_index := flag.Int("test", 0, "Test data to use")
	flag.Parse()
	fmt.Printf("Loading %d test, looking for %d structs\n", *test_index, *desired_struct)
	test_data := smbiosdata.GetTestData(*test_index)
	err = test_data.DecodeBase64Fields()
	err = test_data.LoadEPSStruct()
	printObj(test_data.EPS)
}

func printObj(obj any) {
	js, err := json.MarshalIndent(obj, "", "  ")
	if err != nil {
		log.Println("Error unmarshaling")
	}
	fmt.Println(string(js))
}
