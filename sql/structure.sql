CREATE DATABASE IF NOT EXISTS pizzeria;
USE pizzeria;

CREATE TABLE pizzas (
  id int(11) NOT NULL AUTO_INCREMENT,
  price int(11) unsigned,
  name text COLLATE utf8_hungarian_ci,
  PRIMARY KEY (id),
  UNIQUE KEY id_UNIQUE (id)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

CREATE TABLE ingredients (
  id int(11) NOT NULL AUTO_INCREMENT,
  name text COLLATE utf8_hungarian_ci,
  available tinyint(2) NOT NULL DEFAULT 1,
  gluten_free tinyint(2) NOT NULL DEFAULT 0,
  PRIMARY KEY (id),
  UNIQUE KEY id_UNIQUE (id)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

CREATE TABLE pizza_ingredients (
  pizza_id int(11) NOT NULL,
  ingredient_id int(11) NOT NULL,
  KEY pizza_id (pizza_id),
  KEY ingredient_id (ingredient_id),
  CONSTRAINT fk_pizza_id FOREIGN KEY (pizza_id) REFERENCES pizzas (id) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT fk_ingredient_id FOREIGN KEY (ingredient_id) REFERENCES ingredients (id) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE customers (
  id int(11) NOT NULL AUTO_INCREMENT,
  name text COLLATE utf8_hungarian_ci,
  email varchar(255) CHARACTER SET utf8 COLLATE utf8_hungarian_ci NOT NULL,
  password text COLLATE utf8_hungarian_ci,
  PRIMARY KEY (id),
  UNIQUE KEY id_UNIQUE (id),
  UNIQUE KEY email_UNIQUE (email)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;

CREATE TABLE orders (
  id int(11) NOT NULL AUTO_INCREMENT,
  customer_id int(11) NOT NULL,
  started_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  completed_at TIMESTAMP NULL DEFAULT NULL,
  price int(11) unsigned,
  address text COLLATE utf8_hungarian_ci,
  PRIMARY KEY (id),
  KEY customer_id (customer_id),
  CONSTRAINT fk_customer_id FOREIGN KEY (customer_id) REFERENCES customers (id) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;


CREATE TABLE pizzas_for_orders (
  order_id int(11) NOT NULL,
  pizza_id int(11) NOT NULL,
  number_of_pizzas int(11) NOT NULL,
  KEY pizza_id (pizza_id),
  KEY order_id (order_id),
  CONSTRAINT fk_pizzas_id FOREIGN KEY (pizza_id) REFERENCES pizzas (id) ON DELETE CASCADE ON UPDATE NO ACTION,
  CONSTRAINT fk_order_id FOREIGN KEY (order_id) REFERENCES orders (id) ON DELETE CASCADE ON UPDATE NO ACTION
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
