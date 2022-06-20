package command

import (
	"fmt"
	"os"
	"strings"
)

func MakeService(arg string) {
	fmt.Println("fungsi create service")

	split := strings.Split(arg, "=")
	tableName := strings.Title(split[0])
	tableField := split[1]
	fields := field(tableField, "service")

	f, err := os.OpenFile("app/service/"+strings.ToLower(tableName)+"_service.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(
		"package service \n" +
			"\n" +
			"import (\n" +
			"	\"go_harians/app/dto\"\n" +
			"	\"go_harians/app/repository\"\n" +
			"	\"go_harians/database/migrations\"\n" +
			")\n\n" +
			"type I" + tableName + "Service interface{\n" +
			"	FindAll() ([]migrations." + tableName + ", error)\n" +
			"	FindById(id uint64) (migrations." + tableName + ", error)\n" +
			"	Create(data dto.Create" + tableName + ") (migrations." + tableName + ", error)\n" +
			"\n}\n\n" +
			"type service struct {\n" +
			"	service repository.I" + tableName + "Repository" +
			"\n}\n\n" +
			"func NewService" + tableName + "(repository repository.I" + tableName + "Repository) *service {\n" +
			"	return &service{repository}" +
			"\n}\n\n" +
			"func (s *service) FindAll() ([]migrations." + tableName + ", error) {\n" +
			"	return s.service.FindAll()" +
			"\n}\n\n" +
			"func (s *service) FindById(id uint64) (migrations." + tableName + ", error) {\n" +
			"	return s.service.FindById(id)" +
			"\n}\n\n" +
			"func (s *service) Create(data dto.Create" + tableName + ") (migrations." + tableName + ", error) {\n" +
			"	datas := migrations." + tableName + "{\n" +
			fields +
			"	\n}\n" +
			"	return s.service.Create(datas)" +
			"\n}\n\n"); err != nil {
		panic(err)
	}
}
