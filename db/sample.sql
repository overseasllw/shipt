INSERT INTO
   `category` (`id`, `name`, `description`, `sort_order`, `parent_id`, `created_at`, `deleted_at`)
VALUES
   (
      1, 'Fresh', 'Fruit', 1, NULL, '2018-08-11 12:01:53', NULL
   )
,
   (
      2, 'Meat', 'meat', 2, NULL, '2018-08-11 12:01:53', NULL
   )
,
   (
      3, 'Office', 'office', 3, NULL, '2018-08-11 12:01:53', NULL
   )
;
INSERT INTO
   `category_product` (`id`, `category_id`, `product_id`, `created_at`, `deleted_at`)
VALUES
   (
      1, 1, 1, '2018-08-11 12:03:19', NULL
   )
,
   (
      2, 2, 4, '2018-08-11 12:03:23', NULL
   )
,
   (
      3, 3, 2, '2018-08-11 12:03:24', NULL
   )
,
   (
      4, 1, 3, '2018-08-11 12:03:24', NULL
   )
;
INSERT INTO
   `product_entity` (`id`, `product_name`, `sku`, `price`, `special_price`, `product_type`, `description`, `unit_weight`, `bar_code`, `created_at`, `deleted_at`)
VALUES
   (
      1, 'Apple', '87389', 1.5000, NULL, 'food', 'oganic', 1.0000, '32323878', '2018-08-11 12:03:51', NULL
   )
,
   (
      2, 'note book', '33234', 2.0000, NULL, 'office', 'note', NULL, '88329884', '2018-08-11 12:03:51', NULL
   )
,
   (
      3, 'orange juice', '32355', 4.5000, NULL, 'food', NULL, NULL, '45452452', '2018-08-11 12:03:51', NULL
   )
,
   (
      4, 'beef', '32343', 5.0000, NULL, 'food', NULL, 2.5000, '32423423', '2018-08-11 12:03:51', NULL
   )
;
INSERT INTO
   `sale_order` (`id`, `confirmation_id`, `order_status`, `user_id`, `base_subtotal`, `subtotal`, `discount_amount`, `tax_amount`, `grand_total`, `shipping_amount`, `shipping_address_id`, `billing_address_id`, `coupon_rule_id`, `created_at`, `deleted_at`, `updated_at`)
VALUES
   (
      1, 10000001, 'received', 1, 1.5000, 1.5000, NULL, NULL, 1.5000, NULL, 1, 1, NULL, '2018-08-11 12:04:57', NULL, '2018-08-11 12:07:51'
   )
,
   (
      2, 10000002, 'received', 2, 4.5000, 4.5000, NULL, NULL, 4.5000, NULL, 2, 2, NULL, '2018-08-10 12:04:57', NULL, '2018-08-11 12:07:51'
   )
,
   (
      3, 10000003, 'received', 3, 4.0000, 4.0000, NULL, NULL, 4.0000, NULL, 3, 3, NULL, '2018-08-11 12:04:57', NULL, '2018-08-11 12:07:51'
   )
,
   (
      4, 10000004, 'received', 1, 4.0000, 4.0000, NULL, NULL, 4.0000, NULL, 2, 2, NULL, '2018-08-10 12:04:57', NULL, '2018-08-11 12:07:51'
   )
;
INSERT INTO
   `sale_order_item` (`id`, `sale_order_id`, `product_id`, `qty`, `weight`, `price`, `base_subtotal`, `subtotal`, `discount_amount`, `tax_amount`, `created_at`, `deleted_at`)
VALUES
   (
      1, 1, 1, NULL, 1.5000, 1.5000, 2.2500, 2.2500, NULL, NULL, '2018-08-11 12:05:57', NULL
   )
,
   (
      2, 2, 3, 1, NULL, 4.5000, 4.5000, 4.5000, NULL, NULL, '2018-08-10 12:05:57', NULL
   )
,
   (
      3, 3, 2, 2, NULL, 4.0000, 8.0000, 8.0000, NULL, NULL, '2018-08-11 12:05:57', NULL
   )
,
   (
      4, 4, 2, 1, NULL, 4.0000, 4.0000, 4.0000, NULL, NULL, '2018-08-10 12:05:57', NULL
   )
;
INSERT INTO
   `user_` (`id`, `username`, `email`, `password`, `firstname`, `lastname`, `middlename`, `created_at`, `deleted_at`)
VALUES
   (
      1, 'liwei', 'hi@hello.com', NULL, 'liwei', 'lin', NULL, '2018-08-11 12:06:51', NULL
   )
,
   (
      2, 'john', 'john@world.com', NULL, 'john', 'hhi', NULL, '2018-08-11 12:06:51', NULL
   )
,
   (
      3, 'testing', 'test@me.com', NULL, 'test', 'test', NULL, '2018-08-11 12:06:51', NULL
   )
,
   (
      4, 'liwei2', 'test2@me.com', NULL, 'test2', 'test2', NULL, '2018-08-09 15:47:51', NULL
   )
;
