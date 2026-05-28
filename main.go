package main

import (
	"flag"
	"fmt"
	"log"

	smbiosdata "github.com/grep-michael/SMBIOS_parser/SMBiosLib/SMBiosData"
	utility "github.com/grep-michael/SMBIOS_parser/SMBiosLib/Utility"
)

func main() {
	//file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY, 0666)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	//flag handling
	desired_struct := flag.Int("struct", 4, "structure to print")
	test_index := flag.Int("test", 0, "Test data to use")
	flag.Parse()

	fmt.Printf("Loading %d test, looking for %d structs\n", *test_index, *desired_struct)
	test_data := smbiosdata.GetTestData(*test_index)
	log.Println(test_data)

	err := test_data.LoadDMITable()
	if err != nil {
		log.Printf("Error loading dmi table: %v\n", err)
		return
	}
	table := test_data.DMITable
	log.Printf("Tabled:\n\t%v\n\n", test_data.DMITable)
	utility.PrintObj(table.Structs[*desired_struct])
}
