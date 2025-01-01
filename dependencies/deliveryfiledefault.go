package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func WriteDeliveryHandlerFile(projectFile string, initName string) error {
	initNameCapitalize := strings.Title(initName)
	domainContent :=
		`package handler

	import (
		"` + projectFile + `/domain"
		"github.com/go-playground/validator/v10"
		"github.com/gofiber/fiber/v2"
		"github.com/labstack/gommon/log"
		"github.com/valyala/fasthttp"
	)
	
	type ` + initNameCapitalize + `Handler struct {
		` + initNameCapitalize + `UseCase domain.` + initNameCapitalize + `UseCase
	}

	func (handler *` + initNameCapitalize + `Handler) ` + initNameCapitalize + `ExampleFunction(c *fiber.Ctx) (err error){
		// Example body parse
		var req domain.` + initNameCapitalize + `ExampleRequest
		err = c.BodyParser(&req)
		if err != nil {
			return c.SendStatus(fasthttp.StatusBadRequest)
		}

		// Example validate
		validate := validator.New()
		err = validate.Struct(req)
		if err != nil {
			log.Error(err)
			return c.SendStatus(fasthttp.StatusBadRequest)
		}
		return
	}
`

	filePath := fmt.Sprintf("%s/"+initName+"/delivery/http/handler/"+initName+".go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}

func WriteDeliveryAPIFile(projectFile string, initName string) error {
	initNameCapitalize := strings.Title(initName)
	domainContent :=
		`package handler

	import (
		"` + projectFile + "/" + initName + `/delivery/http/handler"
		"` + projectFile + `/domain"
		"github.com/gofiber/fiber/v2"
		"github.com/gofiber/fiber/v2/middleware/cors"
		"github.com/spf13/viper"
	)

	// @title Open API for Learning Information Master Data
	// @version 1.0.0
	// @description Open API for Learning Information Master Data
	// @host localhost:8099
	// @BasePath /

	// @securityDefinitions.apikey bearerAuth
	// @in header
	// @name Authorization
	// @description Input only value of barrer token prefix, e.g. "abcde12345".
	
	type ` + initNameCapitalize + `Handler struct {
		` + initNameCapitalize + `UseCase domain.` + initNameCapitalize + `UseCase
	}

	func RouterAPI(app *fiber.App, ` + initNameCapitalize + ` domain.` + initNameCapitalize + `UseCase) {
		handler` + initNameCapitalize + ` := &handler.` + initNameCapitalize + `Handler{` + initNameCapitalize + `UseCase: ` + initNameCapitalize + `}
	
		basePath := viper.GetString("server.base_path")
	
		cms := app.Group(basePath)
	
		cms.Use(cors.New(cors.Config{
			AllowOrigins: viper.GetString("middleware.allows_origin"),
		}))
	
		// Example request method
		cms.Post("/", handler` + initNameCapitalize + `.` + initNameCapitalize + `ExampleFunction)
		cms.Patch("/",handler` + initNameCapitalize + `.` + initNameCapitalize + `ExampleFunction)
		cms.Get("/",handler` + initNameCapitalize + `.` + initNameCapitalize + `ExampleFunction)
		cms.Delete("/",handler` + initNameCapitalize + `.` + initNameCapitalize + `ExampleFunction)
		cms.Get("/:id", handler` + initNameCapitalize + `.` + initNameCapitalize + `ExampleFunction)
	
	}
	
`

	filePath := fmt.Sprintf("%s/"+initName+"/delivery/http/api.go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}

func WriteDelivery(projectFile string, initName string) error {
	err := WriteDeliveryHandlerFile(projectFile, initName)
	if err != nil {
		return err
	}

	err = WriteDeliveryAPIFile(projectFile, initName)
	if err != nil {
		return err
	}

	return nil
}
