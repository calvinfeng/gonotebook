# Raw SQL in Go

Let's write SQL in Go. Although it may seem a lot of work at first, it is actually the ultimate
route to go if you want to write highly performant applications.

## Create Database

Use `psql` to create a database.

    $ psql
    create database rawsql with owner = cfeng
