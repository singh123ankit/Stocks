DROP DATABASE IF EXISTS stocksdb;
CREATE DATABASE stocksdb;

\c stocksdb;

CREATE TABLE stocks(
  stockid SERIAL PRIMARY KEY,
  name TEXT,
  price INT,
  company TEXT
);
