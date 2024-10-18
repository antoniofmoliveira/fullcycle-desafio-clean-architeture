-- execute as root before first time use of database
CREATE DATABASE IF NOT EXISTS ordersystem;
CREATE USER 'ordersystem'@'%' IDENTIFIED BY 'ordersystem';
GRANT ALL  ON  ordersystem.* TO 'ordersystem'@'%';
