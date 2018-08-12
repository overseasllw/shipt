SELECT u.id                                         custeomr_id,
       u.firstname                                  customer_first_name,
       cp.category_id,
       c.name                                       category_name,
       Sum(Ifnull(qty, 0)) + Sum(Ifnull(weight, 0)) number_purchased
FROM   sale_order_item soi
       JOIN category_product cp ON cp.product_id = soi.product_id
       JOIN sale_order so ON so.id = soi.sale_order_id
       JOIN user_ u ON u.id = so.user_id
       JOIN category c ON c.id = cp.category_id
GROUP  BY cp.category_id,
          so.user_id;
