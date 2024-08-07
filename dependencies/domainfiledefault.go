package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func WriteDomainFile(projectFile string, initName string) error {
	initName = strings.Title(initName)
	domainContent := `package domain

import (
	"context"
	"time"
)

// Delete if you don't use this struct
type ` + initName + `ExampleRequest struct {
	StringType  string    ` + "`" + `json:"string_type" form:"string_type"` + "`" + `
	Inttype     int       ` + "`" + `json:"int_type" form:"int_type"` + "`" + `
	TimeType    time.Time ` + "`" + `json:"time" form:"time"` + "`" + `
	Float32Type float32   ` + "`" + `json:"float32_type" form:"float32_type"` + "`" + `
	Float64Type float64   ` + "`" + `json:"float64_type" form:"float64_type"` + "`" + `
	BooleanType bool      ` + "`" + `json:"boolean_type" form:"boolean_type"` + "`" + `
	// Another object
}

// Delete if you don't use this struct
type ` + initName + `ExampleResponse struct {
	StringType  string    ` + "`" + `json:"string_type"` + "`" + `
	Inttype     int       ` + "`" + `json:"int_type"` + "`" + `
	TimeType    time.Time ` + "`" + `json:"time"` + "`" + `
	Float32Type float32   ` + "`" + `json:"float32_type"` + "`" + `
	Float64Type float64   ` + "`" + `json:"float64_type"` + "`" + `
	BooleanType bool      ` + "`" + `json:"boolean_type"` + "`" + `
	// Another object
}


// Delete if you don't use this use case
type ` + initName + `UseCase interface {
	` + initName + `ExampleFunction(ctx context.Context, request ` + initName + `ExampleRequest) (response ` + initName + `ExampleResponse, err error)
}

// Delete if you don't use this repository
type ` + initName + `SQLRepo interface {
	` + initName + `ExampleFunction(ctx context.Context, request ` + initName + `ExampleRequest) (response ` + initName + `ExampleResponse, err error)
}

// Delete if you don't use this redis
type ` + initName + `RedisRepo interface {
	` + initName + `ExampleFunction(ctx context.Context, request ` + initName + `ExampleRequest) (response ` + initName + `ExampleResponse, err error)
}
`

	filePath := fmt.Sprintf("%s/domain/"+initName+".go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}
