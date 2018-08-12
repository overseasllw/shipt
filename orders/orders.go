package orders

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"../common"
	m "../models"
	"github.com/labstack/echo"
)

//return sales numbers by user and category
func GetCategorySalesHandler(c echo.Context) (err error) {
	rows, err := common.DB.Query(`select u.id custeomr_id,u.firstname customer_first_name,
    cp.category_id,c.name category_name,sum(ifnull(qty,0))+sum(ifnull(weight,0)) number_purchased
    from sale_order_item soi
    join category_product cp on cp.product_id = soi.product_id
    join sale_order so on so.id = soi.sale_order_id
    join user_ u on u.id =  so.user_id
    join category c on c.id = cp.category_id
    group by category_id,so.user_id`)

	categorySales := []m.CategorySales{}
	if err != nil {
		if err != sql.ErrNoRows {
			return c.JSON(200, categorySales)
		}
		return c.JSON(500, nil)
	}
	defer rows.Close()

	for rows.Next() {
		cs := m.CategorySales{}
		rows.Scan(&cs.CustomerId, &cs.FirstName, &cs.CategoryId, &cs.CategoryName, &cs.NumberPurchased)
		categorySales = append(categorySales, cs)
	}

	return c.JSON(200, categorySales)
}

//Get Product sale By date range and {day/week/month}
func GetSaleByDateRange(c echo.Context) (err error) {
	from := c.QueryParam("from")
	to := c.QueryParam("to")
	queryType := c.QueryParam("type")
	export := c.QueryParam("export")

	dateFormat := ""
	if queryType == "day" {
		dateFormat = `,date_format(soi.created_at,'%Y-%m-%d')`
	} else if queryType == "month" {
		dateFormat = `,date_format(soi.created_at,'%Y-%m')`
	} else if queryType == "year" {
		dateFormat = `,date_format(soi.created_at,'%Y')`
	} else if queryType == "week" {
		dateFormat = `,concat(date_format(soi.created_at,'%Y'),"/",week(soi.created_at))`
	}

	query := `select soi.product_id` + dateFormat + ` dt,
		pe.product_name,sum(ifnull(soi.qty,0))+sum(ifnull(soi.weight,0)) sold
		from sale_order_item soi
		join product_entity pe on pe.id = soi.product_id `

	if strings.TrimSpace(from) != "" {
		query += `where soi.created_at >= '` + from + " 00:00:00'"
	}
	if strings.TrimSpace(to) != "" {
		if strings.TrimSpace(from) != "" {
			query += " and "
		} else {
			query += " where "
		}
		query += " soi.created_at <= '" + to + " 23:59:59'"
	}
	query += ` group by soi.product_id` + dateFormat + ` order by dt asc`

	pss := []m.ProductSaleByTimeRange{}
	queryRows, err := common.DB.Query(query)
	if err != nil {
		if err == sql.ErrNoRows {
			return c.JSON(200, pss)
		}
	}
	defer queryRows.Close()

	for queryRows.Next() {
		ps := m.ProductSaleByTimeRange{}
		queryRows.Scan(&ps.ProductId, &ps.Date, &ps.ProductName, &ps.QtySold)
		pss = append(pss, ps)
	}
	if export == "true" {
		filename := time.Now().Format(time.RFC3339)
		exportCSV(pss, c, filename+"_by_"+queryType)

		return c.JSON(204, nil)
	}
	return c.JSON(200, pss)
}

//export CSV file
func exportCSV(pss []m.ProductSaleByTimeRange, c echo.Context, filename string) {
	file := &bytes.Buffer{}
	//	writer := csv.NewWriter(file)
	//defer writer.Flush()

	//csv := [][]string{}
	//csv = append(csv, m.CSV_TITLE)
	file.WriteString(strings.Join(m.CSV_TITLE, ",") + "\n")
	for _, ps := range pss {
		line := []string{ps.Date, fmt.Sprintf("%d", ps.ProductId), ps.ProductName,
			fmt.Sprintf("%.2f", ps.QtySold)}
		//csv = append(csv, line)
		//writer.Write(line)
		file.WriteString(strings.Join(line, ",") + "\n")
	}

	//writer.WriteAll(csv)
	c.Response().Header().Set("Content-Type", "text/csv")
	c.Response().Header().Set("Content-Disposition", "attachment;filename="+filename+".csv")
	defer c.Response().Flush()
	//log.Print(string(file.Bytes()))
	c.Response().Write(file.Bytes())
}

//Get Customer Orders
func GetOrderHistoryForCustomer(c echo.Context) (err error) {
	customerId := c.Param("customer_id")

	orderRows, err := common.DB.Query(`select id,confirmation_id,order_status,user_id,
		base_subtotal,subtotal,discount_amount,tax_amount,grand_total,shipping_amount,
		created_at from sale_order where user_id=?`, customerId)
	orders := []m.Order{}
	if err != nil {
		if err != sql.ErrNoRows {
			return c.JSON(200, orders)
		}
		return c.JSON(500, nil)
	}
	defer orderRows.Close()
	for orderRows.Next() {
		o := m.Order{}
		err = orderRows.Scan(&o.Id, &o.ConfirmationId, &o.Status, &o.UserId, &o.BaseSubtotal,
			&o.Subtotal, &o.DiscountAmount, &o.TaxAmount, &o.GrandTotal, &o.ShippingAmount,
			&o.CreatedAt)
		o.OrderItems, err = getOrderItems(o.Id)
		orders = append(orders, o)
	}

	return c.JSON(200, orders)
}

//retrieve Order item Details
func getOrderItems(orderId int64) (items []m.OrderItem, err error) {
	itemRows, err := common.DB.Query(`select id,sale_order_id,product_id,qty,weight,
		price,base_subtotal,subtotal,discount_amount,tax_amount,created_at from sale_order_item
		where sale_order_id =?`, orderId)
	if err != nil {
		if err != sql.ErrNoRows {
			return items, err
		}
		return items, err
	}
	defer itemRows.Close()
	for itemRows.Next() {
		oi := m.OrderItem{}
		itemRows.Scan(&oi.Id, &oi.OrderId, &oi.ProductId, &oi.Qty, &oi.Weight, &oi.Price,
			&oi.BaseSubtotal, &oi.Subtotal, &oi.DiscountAmount, &oi.TaxAmount, &oi.CreatedAt)
		items = append(items, oi)
	}
	return items, err
}
