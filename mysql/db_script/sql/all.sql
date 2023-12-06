CREATE SCHEMA `recharge_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci ;

USE `recharge_db`;

--
-- Table structure for table `order_info`
--

CREATE TABLE `recharge_db`.`recharge_info` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `trade_number` VARCHAR(255) NOT NULL COMMENT 'Delivery service order number',
  `shop_order_number` VARCHAR(255) NOT NULL COMMENT 'Store Order Number',
  `shop_type` INT NOT NULL DEFAULT 1 COMMENT 'Store Type 1-Taobao',
  `shop_goods_id` VARCHAR(255) NOT NULL COMMENT 'Store - Product ID',
  `shop_goods_name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Store - Product Name',
  `shop_account_id` VARCHAR(255) NOT NULL COMMENT 'Store - Recharge Account',
  `shop_buyer_phone_number` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Store Buyer Phone',
  `shop_buyer_id` VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Store - Buyer ID',
  `shop_order_create_at` DATETIME NOT NULL COMMENT 'The timestamp of the store',
  `amount` DECIMAL(10,2) NOT NULL COMMENT 'Amount',
  `single_amount` DECIMAL(10,2) NOT NULL COMMENT 'Actual paid unit price of goods',
  `total_amount` DECIMAL(10,2) NOT NULL COMMENT 'Total paid price of goods',
  `count` INT NOT NULL DEFAULT 0 COMMENT 'Quantity to be shipped',
  `real_recharge_count` INT NOT NULL DEFAULT 0 COMMENT 'Actual recharge quantity',
  `trade_create_at` DATETIME NOT NULL ,
  `trade_update_at` DATETIME NOT NULL ,
  `try_recharge_count` INT NOT NULL DEFAULT 0,
  `status` INT NOT NULL DEFAULT 0 COMMENT 'Status: 0- Not processed yet; 2- Partial shipment of goods; 3- Delivery completed;',
  `game_money` INT NOT NULL DEFAULT 0 COMMENT 'Number of game coins corresponding to each transaction',
  `game_type`  INT NOT NULL DEFAULT 0 COMMENT 'game type 1-inter or 2-fee',
  PRIMARY KEY (`id`),
  INDEX `idx_trade` (`trade_number` ASC),
  INDEX `idx_status` (`status` ASC),
  INDEX `idx_order_number` (`shop_order_number` ASC, `shop_type` ASC))
ENGINE=InnoDB COMMENT = 'Direct charging data sheet';

CREATE TABLE `recharge_db`.`simple_recharge_info` (
  `id` BIGINT NOT NULL AUTO_INCREMENT COMMENT 'ID',
  `trade_number` VARCHAR(255) NOT NULL COMMENT 'Delivery service order number',
  `count` INT NOT NULL DEFAULT 0 COMMENT 'Quantity to be shipped',
  PRIMARY KEY (`id`),
  INDEX `idx_trade` (`trade_number` ASC)
)
ENGINE=InnoDB COMMENT = 'Simple charging data sheet';
