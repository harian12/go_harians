package command

import (
	"fmt"
	"os"
	"strings"
)

func MakeRepository(arg string) {
	fmt.Println("fungsi create repository")

	split := strings.Split(arg, "=")
	tableName := strings.Title(split[0])

	f, err := os.OpenFile("app/repository/"+strings.ToLower(tableName)+"_repo.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(
		"package repository \n" +
			"\n" +
			"import (\n" +
			"	\"gorm.io/gorm\"\n" +
			"	\"go_harians/database/migrations\"\n" +
			")\n\n" +
			"type I" + tableName + "Repository interface{\n" +
			"	FindAll() ([]migrations." + tableName + ", error)\n" +
			"	FindById(id uint64) (migrations." + tableName + ", error)\n" +
			"	Create(data migrations." + tableName + ") (migrations." + tableName + ", error)\n" +
			"\n}\n\n" +
			"type dB struct {\n" +
			"	db *gorm.DB" +
			"\n}\n\n" +
			"func New" + tableName + "(db *gorm.DB) *dB {\n" +
			"	return &dB{db}" +
			"\n}\n\n" +
			"func (r *dB) FindAll() ([]migrations." + tableName + ", error) {\n" +
			"	var data []migrations." + tableName + "\n" +
			"	err := r.db.Find(&data).Error\n" +
			"	return data, err" +
			"\n}\n\n" +
			"func (r *dB) FindById(id uint64) (migrations." + tableName + ", error) {\n" +
			"	var data migrations." + tableName + "\n" +
			"	err := r.db.Find(&data,id).Error\n" +
			"	return data, err" +
			"\n}\n\n" +
			"func (r *dB) Create(data migrations." + tableName + ") (migrations." + tableName + ", error) {\n" +
			"	err := r.db.Create(&data).Error\n" +
			"	return data, err" +
			"\n}\n\n"); err != nil {
		panic(err)
	}
}
