package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func WriteUseCaseFile(projectFile string, initName string) error {
	initNameCapitalize := strings.Title(initName)
	domainContent :=
		`package usecases

	import (
		"` + projectFile + `/domain"
		"context"
	)
	
	type ` + initNameCapitalize + `Usecase struct {
		` + initNameCapitalize + `MySQLRepo domain.` + initNameCapitalize + `SQLRepo
		` + initNameCapitalize + `RedisRepo domain.` + initNameCapitalize + `RedisRepo
	}
	
	func New` + initNameCapitalize + `Usecase(` + initNameCapitalize + `MySQL domain.` + initNameCapitalize + `SQLRepo, ` + initNameCapitalize + `RedisRepo domain.` + initNameCapitalize + `RedisRepo) domain.` + initNameCapitalize + `UseCase {
		return &` + initNameCapitalize + `Usecase{
			` + initNameCapitalize + `MySQLRepo: ` + initNameCapitalize + `MySQL,
			` + initNameCapitalize + `RedisRepo: ` + initNameCapitalize + `RedisRepo,
		}
	}
	
	func (usecase *` + initNameCapitalize + `Usecase) ` + initNameCapitalize + `ExampleFunction(ctx context.Context, request domain.` + initNameCapitalize + `ExampleRequest) (response domain.` + initNameCapitalize + `ExampleResponse, err error) {
		return
	}
`

	filePath := fmt.Sprintf("%s/"+initName+"/usecase/"+initName+".go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}

func WriteUseCase(projectFile string, initName string) error {
	err := WriteUseCaseFile(projectFile, initName)
	if err != nil {
		return err
	}

	return nil
}
