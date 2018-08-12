package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"./orders"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//Requirement 4. Get Category Sale
	e.GET("/api/v1/categories/sales/reports/", orders.GetCategorySalesHandler)

	/*5. An API endpoint that accepts a date range and a day, week, or month and
	returns a breakdown of products sold by quantity per day/week/month.
	/api/v1/sales/products/?from=&to=&type=&export=false
	*/
	/*
		6. Ability to export the results of #5 to CSV.
			/api/v1/sales/products/?from=&to=&type=&export=true
	*/
	e.GET("/api/v1/sales/products/", orders.GetSaleByDateRange)

	// 7. An API endpoint that returns the orders for a customer.
	e.GET("/api/v1/customers/:customer_id/orders/", orders.GetOrderHistoryForCustomer)

	e.Logger.Fatal(e.Start(":8080"))
}
