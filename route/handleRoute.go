package route

import (
	"testcasethree-residentAPI/connection"
	"testcasethree-residentAPI/controllers"
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

	// Auth user
	app.POST("/auth/login", gorm.LoginUser)

	// Province CRUD Route
	app.POST("/province", gorm.PostCreateProvince)
	app.GET("/province", gorm.GetReadProvince)
	app.PATCH("/province", gorm.PatchUpdateProvince)
	app.DELETE("/province", gorm.DeleteRemoveProvince)

	// District CRUD Route
	app.POST("/district", gorm.PostCreateDistrict)
	app.GET("/district", gorm.GetReadDistrict)
	app.PATCH("/district", gorm.PatchUpdateDistrict)
	app.DELETE("/district", gorm.DeleteRemoveDistrict)

	// SubDistrict CRUD Route
	app.POST("/subdistrict", gorm.PostCreateSubDistrict)
	app.GET("/subdistrict", gorm.GetReadSubDistrict)
	app.PATCH("/subdistrict", gorm.PatchUpdateSubDistrict)
	app.DELETE("/subdistrict", gorm.DeleteRemoveSubDistrict)

	// person CRUD Route
	app.POST("/person", gorm.PostCreatePerson)
	app.GET("/person", gorm.GetReadPerson)
	app.PATCH("/person", gorm.PatchUpdatePerson)
	app.DELETE("/person", gorm.DeleteRemovePerson)

	// office CRUD Route
	app.POST("/office", gorm.PostCreateOffice)
	app.GET("/office", gorm.GetReadOffice)
	app.PATCH("/office", gorm.PatchUpdateOffice)
	app.DELETE("/office", gorm.DeleteRemoveOffice)

	return app
}
