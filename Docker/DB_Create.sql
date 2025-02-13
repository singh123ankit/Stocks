CREATE DATABASE IF NOT EXISTS stocksdb;

\c stocksdb;

CREATE TABLE stocks(
  stockid SERIAL PRIMARY KEY,
  name TEXT,
  price INT,
  company TEXT
);
