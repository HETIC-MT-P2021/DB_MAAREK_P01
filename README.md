# Native SQL & GO api

![Go](https://github.com/HETIC-MT-P2021/DB_MAAREK_P01/workflows/Go/badge.svg)

A small exercise to practice native queries in go with api endpoint to render the data.


## Setup

### Project setup
You can test this project by running `docker-compose up --build`

It will build then launch the API and database.
API running on port `:8080`, database on `:3306`
> Naturally you need to wait for the api to connect to the db.

Once all done you can test the endpoints

Here's a small description and postman collection to test them quickly: https://documenter.getpostman.com/view/6583625/TVRg9AdT

### Data Setup

You can use any SQL execution software to add the data to the db container, 
connect to the database with database name `classicmodels`, user `root`, password `root` 
and then run the sql script `mysqlsampledatabase.sql`