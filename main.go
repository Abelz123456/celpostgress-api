package main

import (
	"github.com/celpostgress-api/common"
	"github.com/celpostgress-api/controller"
	"github.com/celpostgress-api/database"
	"github.com/celpostgress-api/docs"
	"github.com/celpostgress-api/middleware"
	"github.com/celpostgress-api/repository"
	"github.com/celpostgress-api/routes"
	"github.com/celpostgress-api/services"
	"github.com/celpostgress-api/utils"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	// err := sentry.Init(sentry.ClientOptions{
	// 	Dsn: "https://64f70a24f5ce416798908acd368f1d5f@o4503903109644288.ingest.sentry.io/4503903111544832",
	// 	// Set TracesSampleRate to 1.0 to capture 100%
	// 	// of transactions for performance monitoring.
	// 	// We recommend adjusting this value in production,
	// 	TracesSampleRate: 1.0,
	// })
	// if err != nil {
	// 	log.Fatalf("sentry.Init: %s", err)
	// }

	config, err := utils.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	docs.SwaggerInfo.Title = "Celestial - Celestial API"
	docs.SwaggerInfo.Description = "Celestial - Celestial App API."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:3100"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"}

	db, err := database.GetConnection(config)
	if err != nil {
		panic(err)
	}

	validate := validator.New()

	bankRepository := repository.NewBankRepository()
	bankService := services.NewBankService(bankRepository, db, validate)
	bankController := controller.NewBankController(bankService)

	// imageFileRepository := repository.NewImageFileRepository()
	// imageFileService := services.NewImageFileService(imageFileRepository, db, validate)
	// imageFileController := controller.NewImageFileController(imageFileService)

	// rajaOngkirService := services.NewRajaOngkirService(db, validate)
	// rajaOngkirController := controller.NewRajaOngkirController(rajaOngkirService)

	// sendEmailService := services.NewSendEmailService(db, validate)
	// sendEmailController := controller.NewSendEmailController(sendEmailService)

	permissionPolicyUserRepository := repository.NewPermissionPolicyUserRepository()
	authService := services.NewAuthService(permissionPolicyUserRepository, db, validate)
	authController := controller.NewAuthController(authService)

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // maximum file 8 MiB
	// router.Use(static.Serve("/public"))
	// router.Use(static.Serve("/", static.LocalFile("/", false)))
	// router.Static("/", "./")
	router.Static("/public", "./public")
	// router.StaticFS("/more_static", http.Dir("my_file_system"))
	// router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	// router.StaticFileFS("/more_favicon.ico", "more_favicon.ico", http.Dir("my_file_system"))

	// router := gin.New()
	makeRoutes(router)
	router.Use(gin.Logger())
	router.Use(gin.CustomRecovery(func(c *gin.Context, recovered interface{}) {

		if common.NotFoundErrors(c, recovered) {
			return
		}

		if common.ValidationErrors(c, recovered) {
			return
		}

		common.InternalServerError(c, recovered)

	}))

	router.SetTrustedProxies([]string{"192.168.1.2"})

	v1 := router.Group(docs.SwaggerInfo.BasePath)
	{
		v1.Use(middleware.JWT())
		bank := v1.Group("/bank")
		routes.BankRouter(bankController, bank)
	}

	v2 := router.Group(docs.SwaggerInfo.BasePath)
	{
		auth := v2.Group("/auth")
		routes.AuthRouter(authController, auth)
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// router.Use(CORSMiddleware())
	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"*"},
	// 	AllowMethods:     []string{"PUT", "GET", "OPTIONS", "POST", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	// router.Static("/public", "./public")
	// router.StaticFS("/", gin.Dir("dist", false))
	errPort := router.Run(":3100")

	common.PanicIfError(errPort)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func makeRoutes(r *gin.Engine) {
	cors := func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "*")
		c.Writer.Header().Set("Content-Type", "application/json")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		}
		c.Next()
	}
	r.Use(cors)
}
