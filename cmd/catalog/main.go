package catalog

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/Alefeoliveira/imersao17/goapi/internal/db"
	"github.com/Alefeoliveira/imersao17/goapi/internal/service"
	"github.com/Alefeoliveira/imersao17/goapi/internal/webserver"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/imersao17")
	if err != nil {
		panic(err.Error())
	}
	defer database.Close()

	categoryDB := db.NewCategoryDB(database)
	categoryService := service.NewCategoryService(*categoryDB)

	productDb := db.NewProductDB(database)
	productService := service.NewProductService(*productDb)

	webCategoryHandler := webserver.NewWebCategoryHandler(*categoryService)
	webProductHandler := webserver.NewWebProductHandler(*productService)

	c := chi.NewRouter()
	c.Use(middleware.Logger)
	c.Use(middleware.Recoverer)
	c.Get("/category/{id}", webCategoryHandler.GetCategory)
	c.Get("/category", webCategoryHandler.GetCategories)
	c.Post("/category", webCategoryHandler.CreateCategory)

	c.Get("/product/{id}", webProductHandler.GetProduct)
	c.Get("/product", webProductHandler.GetProducts)
	c.Get("/product/category/{categoryID}", webProductHandler.GetProductByCategoryID)
	c.Post("/product", webProductHandler.CreateProduct)

	fmt.Println("server is running on port 8080")

	http.ListenAndServe(":8080", c)
}
