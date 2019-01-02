# User Authentication

## Postgres Setup

I am going to use PostgreSQL for this project, so let's create one. The superuser on my computer is
`cfeng` so I will use that to create a database named `go_user_auth`. If you don't have a role or
wish to create a separate role for this project, then just do the following

    $ psql postgres
    postgres=# create role <name> superuser login;

Create a database named `go_user_auth` with owner pointing to whichever role you like. I am using
cfeng on my computer.

    $ psql postgres
    postgres=# create database go_user_auth with owner=cfeng;

Actually just in case you don't remember the password to your `ROLE`, do the following

    postgres=# alter user <your_username> with password <whatever you like>

I did mine with

    postgres=# alter user cfeng with password 'cfeng';

## Project Dependencies

I am going to introduce couple new open source libraries to you for this project:

* `sirupsen/logrus`
* `gorilla/mux`
* `jinzhu/gorm` -  Object Relational Mapping for Go

You should take a look at their Github page and see what they are for before you start working on
this project.

## Project User Authentication

* [Lesson 5 User Authentication](https://www.youtube.com/channel/UCoKwJSadNdeJkpfBpI-f5Ow)

## Bonus & Additional Resource

If you want to learn more about session storage, security, encryption, and many other topics
relating to web applications, take a look at this eBook [astaxie][1]

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/userauth)

[1]:https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/