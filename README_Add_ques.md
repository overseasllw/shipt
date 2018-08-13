# Additional questions

#### We want to give customers the ability to create lists of products for a one-click ordering of bulk items. How would you design the tables, what are the pros and cons of your approach?

*create two tables `wishlist` which is bulk list   and `wishlist_product`  which is bulk items*
```
wishlist(
  id         int
  user_id    int
  name       text
  ...
  created_at timestamp
  deleted_at timestamp
)
```

```
wishlist_product(
  id               int
  product_id       int
  wishlist_id      int
  qty              int
  weight           float
  ...
  created_at       timestamp
  deleted_at       timestamp
)
```
- pros: customers can have several different bulk list for one click purchased.
      easy to check how many lists not group by or distinct need to be used.
      normalized list table and list item table, in list table don't have a lot
      of duplicate record

- cons: two many table joins when onClick to place order it will slow down the speed.
      customer may have empty list, when deleted list need to cascade deleted from
      item table or item table will have redundant data


#### If Shipt knew the exact inventory of stores, and when facing a high traffic and limited supply of a particular item, how do you distribute the inventory among customers checking out?

1. During checkout, we can make see how many customer are checkout at the same time, if after checkout our it will
out of stock, then we can use a optimize algorithm make as many as customer can place order, and we may can make sure the return customer checkout first if we want get more retention, if we want to gain more new customers we let new customer checkout first. and we can set a order limit(how many single items you can order in one order), this will also can make more people can order.

2. We can set up a buffer inventory, if the inventory less than buffer inventory stop sell it, make sure not over sale. And if inventory is enough, when customer add to cart, we can hold that inventory for few minutes, if customer not checked out after the time we set up, and put the inventory back.

3. Because we know the inventory, so the other way can use cache to deduct inventory cache can handle high concurrency better than database, the use a message queue to update database inventory.
