CREATE TABLE IF NOT EXISTS product_entity
  (
     id            INT(10) NOT NULL auto_increment,
     product_name  VARCHAR(50),
     sku           VARCHAR (20),
     price         DECIMAL(12, 4),
     special_price DECIMAL(12, 4),
     product_type  VARCHAR(50),
     description   TEXT,
     unit_weight   DECIMAL(12, 4),
     bar_code      VARCHAR(50),
     created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY(id)
  );

CREATE TABLE IF NOT EXISTS category
  (
     id          INT(10) NOT NULL auto_increment,
     name        VARCHAR(50),
     description TEXT,
     sort_order  INT(10),
     created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY(id)
  );

CREATE TABLE IF NOT EXISTS category_product
  (
     id          INT(10) NOT NULL auto_increment,
     category_id INT(10),
     product_id  INT(10),
     created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     PRIMARY KEY(id)
  );

CREATE TABLE IF NOT EXISTS sale_order
  (
     id              INT(10) NOT NULL auto_increment,
     confirmation_id BIGINT(10),
     created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     order_status    ENUM('received', 'processing', 'in_transit',
     'out_for_delivery', 'delivered', 'canceled', 'refunded'),
     user_id         INT(10),
     base_subtotal   DECIMAL(12, 4)
  ), subtotal DECIMAL(12, 4), discount_amount DECIMAL(12, 4), tax_amount DECIMAL
(12, 4), grand_total DECIMAL(12, 4), shipping_amount DECIMAL(12, 4),
shipping_address_id INT(10), billing_address_id INT(10), coupon_rule_id INT(10),
PRIMARY KEY(id) );

CREATE TABLE IF NOT EXISTS sale_order_item
  (
     id              INT(10) NOT NULL auto_increment,
     sale_order_id   INT(10),
     created_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at      TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     product_id      INT(10),
     qty             INT(10),
     weight          DECIMAL(12, 4),
     price           DECIMAL(12, 4),
     base_subtotal   DECIMAL(12, 4),
     subtotal        DECIMAL(12, 4),
     discount_amount DECIMAL(12, 4),
     tax_amount      DECIMAL(12, 4),
     PRIMARY KEY(id)
  );

CREATE TABLE IF NOT EXISTS user_
  (
     id         INT(10) NOT NULL auto_increment,
     username   VARCHAR(10),
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     deleted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     email      VARCHAR(50),
     password   VARCHAR(50),
     firstname  VARCHAR(50),
     lastname   VARCHAR(50),
     middlename VARCHAR(50),
     PRIMARY KEY(id)
  );
