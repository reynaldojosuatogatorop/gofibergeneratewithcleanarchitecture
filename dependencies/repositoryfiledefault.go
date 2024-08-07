package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func WriteRepositoryRedisFile(projectFile string, initName string) error {
	initNameCapitalize := strings.Title(initName)
	domainContent :=
		`package redis

	import (
		"` + projectFile + `/domain"
		"context"
		"github.com/go-redis/redis/v8"
	)
	
	type redis` + initNameCapitalize + `Repository struct {
		Conn *redis.Client
	}
	
	func NewRedis` + initNameCapitalize + `Repository(Conn *redis.Client) domain.` + initNameCapitalize + `RedisRepo {
		return &redis` + initNameCapitalize + `Repository{Conn}
	}
	
	func (redis *redis` + initNameCapitalize + `Repository) ` + initNameCapitalize + `ExampleFunction(context context.Context, request domain.` + initNameCapitalize + `ExampleRequest) (response domain.` + initNameCapitalize + `ExampleResponse, err error) {
		return response, nil
	}
`

	filePath := fmt.Sprintf("%s/"+initName+"/repository/redis/"+initName+".go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}

func WriteRepositorySQLFile(projectFile string, initName string) error {
	initNameCapitalize := strings.Title(initName)
	domainContent :=
		`package sql

	import (
		"` + projectFile + `/domain"
		"context"
		"database/sql"
	)
	type SQL` + initNameCapitalize + `Repository struct {
		Conn *sql.DB
	}
	
	func NewSQL` + initNameCapitalize + `Repository(Conn *sql.DB) domain.` + initNameCapitalize + `SQLRepo {
		return &SQL` + initNameCapitalize + `Repository{Conn}
	}
	
	func (sql *SQL` + initNameCapitalize + `Repository) ` + initNameCapitalize + `ExampleFunction(ctx context.Context, request domain.` + initNameCapitalize + `ExampleRequest) (response domain.` + initNameCapitalize + `ExampleResponse, err error) {
		return
	}
	
	
`
	filePath := fmt.Sprintf("%s/"+initName+"/repository/sql/"+initName+".go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}

func WriteRepository(projectFile string, initName string) error {
	err := WriteRepositoryRedisFile(projectFile, initName)
	if err != nil {
		return err
	}

	err = WriteRepositorySQLFile(projectFile, initName)
	if err != nil {
		return err
	}

	return nil
}
