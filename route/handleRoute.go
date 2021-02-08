package route

import (
	"testcasethree-residentAPI/connection"
	"testcasethree-residentAPI/controllers"
	"testcasethree-residentAPI/middlewares"
	"testcasethree-residentAPI/models"
	"testcasethree-residentAPI/seeders"

	"github.com/gin-gonic/gin"
)

// RouteHandler func
func RouteHandler(app *gin.Engine) *gin.Engine {

	// connect to postgre db
	pgDB := connection.Connect()
	gorm := controllers.Gorm{DB: pgDB}

	// Migrations
	models.Migrations(pgDB)

	// Seeder
	seeders.SeedProvince(pgDB)
	seeders.SeedDistrict(pgDB)
	seeders.SeedSubDistrict(pgDB)
	seeders.SeedPerson(pgDB)
	seeders.SeedOffice(pgDB)
	seeders.SeedOfficePersonLocation(pgDB)
	seeders.SeedUser(pgDB)

	// Middleware
	authMiddleware := middlewares.Auth

	// Auth user
	app.POST("/auth/login", gorm.LoginUser)

	// Upload
	app.PATCH("/person/photo", authMiddleware, gorm.UploadSingle)

	// Province CRUD Route
	app.POST("/province", authMiddleware, gorm.PostCreateProvince)
	app.GET("/province", gorm.GetReadProvince) // no auth | public route
	app.PATCH("/province", authMiddleware, gorm.PatchUpdateProvince)
	app.DELETE("/province", authMiddleware, gorm.DeleteRemoveProvince)

	// District CRUD Route
	app.POST("/district", authMiddleware, gorm.PostCreateDistrict)
	app.GET("/district", gorm.GetReadDistrict) // no auth | public route
	app.PATCH("/district", authMiddleware, gorm.PatchUpdateDistrict)
	app.DELETE("/district", authMiddleware, gorm.DeleteRemoveDistrict)

	// SubDistrict CRUD Route
	app.POST("/subdistrict", authMiddleware, gorm.PostCreateSubDistrict)
	app.GET("/subdistrict", gorm.GetReadSubDistrict) // no auth | public route
	app.PATCH("/subdistrict", authMiddleware, gorm.PatchUpdateSubDistrict)
	app.DELETE("/subdistrict", authMiddleware, gorm.DeleteRemoveSubDistrict)

	// person CRUD Route
	app.POST("/person", authMiddleware, gorm.PostCreatePerson)
	app.GET("/person", authMiddleware, gorm.GetReadPerson)
	app.PATCH("/person", authMiddleware, gorm.PatchUpdatePerson)
	app.DELETE("/person", authMiddleware, gorm.DeleteRemovePerson)

	// office CRUD Route
	app.POST("/office", authMiddleware, gorm.PostCreateOffice)
	app.GET("/office", authMiddleware, gorm.GetReadOffice)
	app.PATCH("/office", authMiddleware, gorm.PatchUpdateOffice)
	app.DELETE("/office", authMiddleware, gorm.DeleteRemoveOffice)

	// report get route
	app.GET("/report/person/count", authMiddleware, gorm.ReportPersonByGender)
	app.GET("/report/person/office", gorm.ReportPersonOffice)
	app.GET("/report/person/office/count", gorm.ReportCountPersonOfficeByGender)

	return app
}
