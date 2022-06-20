package command

import (
	"fmt"
	"os"
	"strings"
)

func MakeDTO(arg string) {
	fmt.Println("fungsi create DTO")
	split := strings.Split(arg, "=")
	tableName := strings.ToLower(split[0])
	if tableName == "" {
		panic("Nama tabel tidak ditemukan")
	}
	tableField := split[1]
	fields := field(tableField, "dto")

	f, err := os.OpenFile("app/dto/"+tableName+"_dto.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(
		"package dto \n" +
			" import ( \n" +
			"	\"gorm.io/gorm\" \n" +
			"	\"time\"\n" +
			"  )\n" +
			"\n" +
			"type Create" + strings.Title(tableName) + " struct {\n" +
			fields +
			"\n}"); err != nil {
		panic(err)
	}
}
