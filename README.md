# Coding Project instruction
######
More work can be done
1. database design is only several tables, need add trigger for some table,
then when modify table, the updated_at column can update. and need more tables,
like some tables related product, now it only have product_entity(basic info) table.
2. for the current finished APIs, need to consider more to make the data more accurate,
like refund, canceled, returned
3. if this a internal tool admin panel, need to add more analysis function like
customer cohort, product cohort, analysis customer order  etc.
4. add customer list/detail dashboard, order list/detail dashboard, and category,
product add the same
5. in this project i did not user ORM,I can refactor the code use ORM.
6. if it is a admin, add more APIs, like update/delete/post
7. add more common helper function like error handling
######
##### Additional questions
 - find in README_Add_ques.md

## shipt
 - common [all configuration (if run in local change the database username/password)]
 - db [database table script(task 1,2,3), The project user *Mysql* Database]
 - models [all the Object map to database(Order,Product,Category,User) or Some Object used for task]
 - orders [all task functions (api and helper functions)]
 - frontend [a simple frontend created just for show the task result using reactjs and semantic_ui]
 - application.go [the main function]

 - screenshot [the screenshot for some task_4, task_5, task_7]

## How to start
  1. Backend
     *go to shipt folder and run commands*
      > `go get -d ./...`  [install all dependencies] \
      > `go run application.go`

      *before run the it place make sure you have mysql database installed and
      change the username and password in /common/common.go
      in db/sample.sql is some dump data i made up for testing can run it to dump data into local database
      and also there is a ER-Diagram to show the summary*

  2. Frontend
     *go to shipt/frontend folder and run commands*
        > `npm install` [install all dependencies] \
        > `npm start` [start frontend]

#### frontend Url is `http://localhost:3000`

#### In db folder query.sql for task_3
###
*After you start Backend you can use the api from some api client(postman) or from browser
the Url is http://localhost:8080/*

####

For task_3/task_4
```
GET /api/v1/categories/sales/reports/
```
Response body
```
[
  {
  "customer_id": 1,
  "customer_first_name": "liwei",
  "category_id": 1,
  "category_name": "Fresh",
  "number_purchased": 1.5
  },
  {
  "customer_id": 2,
  "customer_first_name": "john",
  "category_id": 1,
  "category_name": "Fresh",
  "number_purchased": 1
  },
]
```

Get data by date range task_5
```
GET /api/v1/sales/products/?from=&to=&type=&export=false
```
Response body
```
[
  {
  "product_id": 3,
  "date": "2018-08-10",
  "product_name": "orange juice",
  "quantity": 1
  }
]
```

Download a csv file for filter data (task_6)
```
GET /api/v1/sales/products/?from=&to=&type=&export=true
```

Get order for customer task_7
```
GET /api/v1/customers/:customer_id/orders/
```
Response body
```
[
  {
  "id": 1,
  "confirmation_id": 10000001,
  "status": "received",
  "customer_id": 1,
  "base_subtotal": 1.5,
  "subtotal": 1.5,
  "discount_amount": null,
  "tax_amount": null,
  "grand_total": 1.5,
  "shipping_amount": null,
  "created_at": "2018-08-11T12:04:57Z",
  "order_items": [
    {
    "id": 1,
    "sale_order_id": 1,
    "product_id": 1,
    "qty": null,
    "price": 1.5,
    "weight": 1.5,
    "base_subtotal": 2.25,
    "subtotal": 2.25,
    "discount_amount": null,
    "tax_amount": null,
    "created_at": "2018-08-11T12:05:57Z"
    }
  ],
  "address": null
  }
]
```
