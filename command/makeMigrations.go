package command

import (
	"errors"
	"os"
	"strings"
)

func MakeMigration(arg string) {

	split := strings.Split(arg, "=")
	tableName := strings.Title(split[0])
	if tableName == "" {
		panic("Nama tabel tidak ditemukan")
	}
	tableField := split[1]
	fields := field(tableField, "migrations")

	/*ccek apakah file sudah ada di folder migrations*/
	if _, err := os.Stat("database/migrations/" + tableName + "_table.go"); !errors.Is(err, os.ErrNotExist) {
		panic("migrations tabel sudah ada")
	}

	f, err := os.OpenFile("database/migrations/"+tableName+"_table.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(
		"package migrations \n" +
			"\n" +
			"import (\n" +
			"	\"gorm.io/gorm\"\n" +
			"	\"time\"" +
			")\n\n" +
			"type " + tableName + " struct {\n" +
			"	ID uint64 `json:\"" + strings.ToLower(tableName) + "\" gorm:\"primaryKey;uniqueIndex\"`\n" +
			fields +
			"	DeletedAt gorm.DeletedAt `json:\"deleted_at\"`\n" +
			"	CreatedAt time.Time `json:\"created_at\"`\n" +
			"	UpdatedAt time.Time `json:\"updated_at\"`\n" +
			"\n}"); err != nil {
		panic(err)
	}
}

func field(fl string, fungsi string) string {
	fieldSplit := strings.Split(fl, ",")
	fields := ""
	for _, a := range fieldSplit {
		aSplit := strings.Split(a, "?")
		name := aSplit[0]

		ruleField := "type:string"

		if len(aSplit) > 1 {
			ruleField = aSplit[1]
		}

		var nameSplit string
		if len(strings.Split(name, "_")) != 1 {
			nameSplit = toUppderFirst(strings.Split(name, "_"))
		} else {
			nameSplit = strings.Title(strings.Split(name, "_")[0])
		}

		/*type data field*/
		tipeField := strings.Split(strings.Split(ruleField, "type:")[1], ";")[0]
		if fungsi == "dto" {
			tipeField = strings.Replace(tipeField, "*", "", 1)
			fields += "		" + nameSplit + " " + tipeField + " `json:\"" + name + "\" validate:\"required\"`\n"
		} else {
			fields += "		" + nameSplit + " " + tipeField + " `json:\"" + name + "\"`\n"
		}

	}
	return fields
}

func toUppderFirst(data []string) string {
	temp := make([]string, len(data))
	for i, v := range data {
		temp[i] = strings.Title(v)
	}
	namaField := strings.Join(temp, "")
	return namaField
}
