package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"gofr.dev/pkg/gofr"
)

type Customer struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
}

func main() {
	dbName := "bhavyadb"
	user := "root"
	password := "12345"
	host := "localhost"
	port := 3306

	// Connect to database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", user, password, host, port, dbName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Initialize gofr object
	app := gofr.New()

	// GET /greet
	app.GET("/greet", func(ctx *gofr.Context) (interface{}, error) {
		// Get the value using the redis instance
		value, err := ctx.Redis.Get(ctx.Context, "greeting").Result()
		return value, err
	})

	// POST /customer/{name}
	app.POST("/customer/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")
		email := ctx.Request().FormValue("email")              // Access email from form data
		phoneNumber := ctx.Request().FormValue("phone_number") // Access phone number from form data

		// Inserting a customer row in database
		_, err := db.ExecContext(ctx, "INSERT INTO customers (name, email, phone_number) VALUES (?, ?, ?)", name, email, phoneNumber)
		if err != nil {
			return nil, err
		}

		return nil, err
	})

	// GET /customer
	app.GET("/customer", func(ctx *gofr.Context) (interface{}, error) {
		var customers []Customer

		// Getting customers from database
		rows, err := db.QueryContext(ctx, "SELECT id, name, email, phone_number FROM customers")
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		// Open file for writing
		file, err := os.Create("customers.csv")
		if err != nil {
			return nil, err
		}
		defer file.Close()

		writer := csv.NewWriter(file)
		writer.Flush()

		for rows.Next() {
			var customer Customer
			err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.PhoneNumber)
			if err != nil {
				return nil, err
			}

			customers = append(customers, customer)
			writer.Write([]string{fmt.Sprintf("%d", customer.ID), customer.Name, customer.Email, customer.PhoneNumber})
		}
		writer.Flush()
		return customers, nil
	})

	// GET /customer/search/{name}
	app.GET("/customer/search/{name}", func(ctx *gofr.Context) (interface{}, error) {
		name := ctx.PathParam("name")

		// Searching for customers by name in the database
		rows, err := db.QueryContext(ctx, "SELECT id, name, email, phone_number FROM customers WHERE name = ?", name)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		var customers []Customer

		for rows.Next() {
			var customer Customer
			err := rows.Scan(&customer.ID, &customer.Name, &customer.Email, &customer.PhoneNumber)
			if err != nil {
				return nil, err
			}
			customers = append(customers, customer)
		}

		return customers, nil
	})

	// Start server
	app.Start()
}
