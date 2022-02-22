CREATE SCHEMA IF NOT EXISTS test AUTHORIZATION postgres;
DROP TABLE IF EXISTS customers;
CREATE TABLE customers (id serial primary key,name varchar (50));