package dependencies

import (
	"fmt"
	"os"
	"strings"
)

func WriteMainFile(projectFile string, initName string) error {
	initNameCapitalize := strings.Title(initName)
	domainContent :=
		`package main

		import (
			"context"
			"crypto/tls"
			"database/sql"
			"flag"
			"fmt"
			"net"
			"strconv"
		
			_DeliveryHTTP "` + projectFile + `/` + initName + `/delivery/http"
			_RepoRedis` + initNameCapitalize + ` "` + projectFile + `/` + initName + `/repository/redis"
			_RepoMySQL` + initNameCapitalize + ` "` + projectFile + `/` + initName + `/repository/sql"
			_Usecase` + initNameCapitalize + ` "` + projectFile + `/` + initName + `/usecase"
			"` + projectFile + `/config"
		
			"github.com/go-redis/redis/v8"
			_ "github.com/go-sql-driver/mysql" // MySQL driver
			"github.com/gofiber/fiber/v2"
			"github.com/gofiber/fiber/v2/middleware/cors"
			"github.com/gofiber/fiber/v2/middleware/recover"
			"github.com/golang-migrate/migrate/v4"
			"github.com/golang-migrate/migrate/v4/database/sqlserver"
			_ "github.com/golang-migrate/migrate/v4/source/file"
			log "github.com/sirupsen/logrus"
			"github.com/spf13/viper"
			"google.golang.org/grpc"
		)
		
		func main() {
			// CLI options parse
			configFile := flag.String("c", "config.yaml", "Config file")
			flag.Parse()
		
			// Config file
			config.ReadConfig(*configFile)
			viper.AutomaticEnv()
		
			// Set log level
			switch viper.GetString("server.log_level") {
			case "error":
				log.SetLevel(log.ErrorLevel)
			case "warning":
				log.SetLevel(log.WarnLevel)
			case "info":
				log.SetLevel(log.InfoLevel)
			case "debug":
				log.SetLevel(log.DebugLevel)
			default:
				log.SetLevel(log.InfoLevel)
			}
		
			// Initialize database connection
			connection := fmt.Sprintf("server=%s;user id=%s;password=%s;port=%s;database=%s",
				viper.GetString("database.host"),
				viper.GetString("database.user"),
				viper.GetString("database.password"),
				viper.GetString("database.port"),
				viper.GetString("database.database"))
			// Tambahkan parameter multiStatements dengan nilai true
			connection += ";multiStatements=true"
		
			dbConn, err := sql.Open("sqlserver", connection)
			if err != nil {
				log.Fatal(err)
			}
		
			defer func() {
				err := dbConn.Close()
				if err != nil {
					log.Fatal(err)
				}
			}()
		
			migrationStatus := viper.GetBool("database.status")
			if migrationStatus {
				// Migrate database if any new schema
				driver, err := sqlserver.WithInstance(dbConn, &sqlserver.Config{})
				if err == nil {
					mig, err := migrate.NewWithDatabaseInstance(viper.GetString("database.path_migrate"), viper.GetString("mysql.database"), driver)
					log.Info(viper.GetString("database.path_migrate"))
					if err == nil {
						err = mig.Up()
						if err != nil {
							if err == migrate.ErrNoChange {
								log.Debug("No database migration")
							} else {
								log.Error(err)
							}
						} else {
							log.Info("Migrate database success")
						}
						version, dirty, err := mig.Version()
						if err != nil && err != migrate.ErrNilVersion {
							log.Error(err)
						}
						log.Debug("Current DB version: " + strconv.FormatUint(uint64(version), 10) + "; Dirty: " + strconv.FormatBool(dirty))
					} else {
						log.Warn(err)
					}
				} else {
					log.Warn(err)
				}
			} else {
				log.Info("Migration status is off.....")
			}
			// Initialize Redis
			var dbRedis *redis.Client
			ctx := context.Background()
			if viper.GetBool("redis.tls_config") {
				// Jika redis.tls_config bernilai true
				dbRedis = redis.NewClient(&redis.Options{
					Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
					Username: viper.GetString("redis.username"),
					Password: viper.GetString("redis.password"),
					DB:       viper.GetInt("redis.database"),
					PoolSize: viper.GetInt("redis.max_connection"),
					TLSConfig: &tls.Config{
						InsecureSkipVerify: true,
					},
				})
			} else {
				// Jika redis.tls_config bernilai false atau tidak ada
				dbRedis = redis.NewClient(&redis.Options{
					Addr:     viper.GetString("redis.host") + ":" + viper.GetString("redis.port"),
					Username: viper.GetString("redis.username"),
					Password: viper.GetString("redis.password"),
					DB:       viper.GetInt("redis.database"),
					PoolSize: viper.GetInt("redis.max_connection"),
				})
			}
		
			log.Info("Redis TLS ", viper.GetBool("redis.tls_config"))
			_, err = dbRedis.Ping(ctx).Result()
			if err != nil {
				log.Fatal(err)
			}
			log.Info("Redis connection established")
		
			// Register repository & usecase ` + initName + `
		
			repoMySQL` + initNameCapitalize + ` := _RepoMySQL` + initNameCapitalize + `.NewSQL` + initNameCapitalize + `Repository(dbConn)
			repoRedis` + initNameCapitalize + ` := _RepoRedis` + initNameCapitalize + `.NewRedis` + initNameCapitalize + `Repository(dbRedis)
		
			usecase` + initNameCapitalize + ` := _Usecase` + initNameCapitalize + `.New` + initNameCapitalize + `Usecase(repoMySQL` + initNameCapitalize + `, repoRedis` + initNameCapitalize + `)
			// server` + initNameCapitalize + ` := _RepoGRPC` + initNameCapitalize + `Object.NewGRPC` + initNameCapitalize + `(usecase` + initNameCapitalize + `)
		
			// Initialize gRPC server
			go func() {
				listen, err := net.Listen("tcp", ":"+viper.GetString("server.grpc_port"))
				if err != nil {
					log.Fatalf("[ERROR] Failed to listen tcp: %v", err)
				}
		
				grpcServer := grpc.NewServer()
				// _RepoGRPC` + initNameCapitalize + `Server.Register` + initNameCapitalize + `orizationServiceServer(grpcServer, server` + initNameCapitalize + `)
				log.Println("gRPC server is running in port", viper.GetString("server.grpc_port"))
				if err := grpcServer.Serve(listen); err != nil {
					log.Fatalf("Failed to serve: %v", err)
				}
			}()
		
			// Initialize HTTP web framework
			app := fiber.New(fiber.Config{
				Prefork:       viper.GetBool("server.prefork"),
				StrictRouting: viper.GetBool("server.strict_routing"),
				CaseSensitive: viper.GetBool("server.case_sensitive"),
				BodyLimit:     viper.GetInt("server.body_limit"),
			})
			app.Use(recover.New())
			app.Use(cors.New(cors.Config{
				AllowOrigins: viper.GetString("middleware.allows_origin"),
			}))
		
			// HTTP routing
			app.Get(viper.GetString("server.base_path")+"/", func(c *fiber.Ctx) error {
				return c.SendString("Hello, World!")
			})

			// Route untuk Swagger
			app.Get("/swagger/*", swagger.HandlerDefault) // Endpoint Swagger
		
			_DeliveryHTTP.RouterAPI(app, usecase` + initNameCapitalize + `)
		
			// Start Fiber HTTP server
			if err := app.Listen(":" + viper.GetString("server.port")); err != nil {
				log.Fatal(err)
			}
		}
		
`

	filePath := fmt.Sprintf("%s/app/main.go", projectFile)
	return os.WriteFile(filePath, []byte(domainContent), 0644)
}
