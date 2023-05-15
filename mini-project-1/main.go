package main

import (
	"fmt"
	"mini-project/modules/products"
	"mini-project/modules/transactions"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/ma_project?parseTime=true"))
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&products.Product{})

	productRepo := products.Repository{DB: db}
	productUsecase := products.Usecase{Repo: productRepo}
	productHandler := products.Handler{Usecase: productUsecase}

	transactionRepo := transactions.Repository{DB: db}
	transactionUsecase := transactions.Usecase{TransacationRepo: transactionRepo, ProductRepo: productRepo}
	transactionHandler := transactions.Handler{Usecase: transactionUsecase}

	router := mux.NewRouter()

	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")
	router.HandleFunc("/products", productHandler.AddProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.EditProduct).Methods("PUT")
	router.HandleFunc("/products/{id}/status", productHandler.SoftDelete).Methods("PATCH")

	router.HandleFunc("/transactions", transactionHandler.GetAll).Methods("GET")
	router.HandleFunc("/transactions/{id}", transactionHandler.GetById).Methods("GET")
	router.HandleFunc("/transactions", transactionHandler.Create).Methods("POST")

	PORT := ":8080"
	fmt.Println("Starting server at localhost", PORT)
	http.ListenAndServe(PORT, router)
}
