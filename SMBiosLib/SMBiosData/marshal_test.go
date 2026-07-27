package smbiosdata

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshal(t *testing.T) {
	data, err := loadLocalSMBIOS()
	if err != nil {
		t.Fatal(err)
	}
	err = data.LoadDMITable()
	if err != nil {
		t.Fatal(err)
	}
	js, err := json.Marshal(data)
	fmt.Println(string(js))
	if err != nil {
		t.Fatal(err)
	}
	var data2 SMBiosData
	err = json.Unmarshal(js, &data2)
	if err != nil {
		t.Fatal(err)
	}
	chunks := data2.DMITable.Structs[4]
	fmt.Println(chunks[0].Data)
}
