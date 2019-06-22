# Popular Web Libraries

As our application gets larger and more complex, we need help from some open source libraries to make our life easier. There are many good web frameworks out there but I am going to show you the couple popular libraries I've enjoyed using.

## [Echo](https://github.com/labstack/echo)

Echo is a minimalist web framework. I am using it primarily for routing and HTTP request/response handling. The setup is very easy; I create a server and then configure handlers for it.

```go
srv := echo.New()
srv.File("/", "public/index.html")
srv.GET("/robots/{name}", RetrieveRobotHandler)
srv.POST("/robots", CreateRobotHandler)
```

Each handler looks like the following.

```go
func RetrieverobotHandler (ctx echo.Context) error {
  robot := &Robot{}

  if err := db.Where("name = ?", ctx.Param("name")).Find(robot).Error; err != nil {
    return echo.NewHTTPError(http.StatusNotFound, err)
  }

  return ctx.JSON(http.StatusOK, users)
}
```

## [Logger](https://github.com/sirupsen/logrus)

Logging is crucial for debugging applications in production. Go provides a default `log` package. However, sometimes you'd want more. `logrus` provides colorful logging and additional log fields to identify the source of errors. More importantly, `logrus` provides integration with [Sentry](https://sentry.io/welcome/).

```go
package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	setLogOutput()
}

func main() {
	logrus.Info("some information")
	logrus.Warn("some warning")
	logrus.Error("some errors")
	logrus.Fatal("big problem")
}

var logToFile = true

func setLogOutput() {
	if !logToFile {
		logrus.SetOutput(os.Stdout)
		return
	}

	f, err := os.OpenFile("example.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}

	logrus.SetOutput(io.MultiWriter(f, os.Stdout))
}

```

## [Command-line Interface](https://github.com/spf13/cobra) \(CLI\)

So far we've been building applications that runs a single command, i.e. whatever you put in your main function. What if we want to have more command like the following.

```bash
go install
myapp runmigration
myapp runserver
myapp reset
...
```

We can accomplish that using the `cobra` package. We define a list of commands and each command is mapping to an execution function, much like `main()`.

```go
var RunMigrationsCmd = &cobra.Command{
  Use:   "runmigrations",
  Short: "run migrations on database",
  RunE:  runMigrations,
}

var RunServerCmd = &cobra.Command{
  Use:   "runserver",
  Short: "run user authentication server",
  RunE:  runServer,
}
```

In our main function, we just need to run the `Execute` function on the root command.

```go
var root = &cobra.Command{
  Use: "myapp",
  Short: "my toy application"
}

func main() {
  root.AddCommand(RunServerCmd, RunMigrationsCmd)
  if err := root.Execute(); err != nil {
    log.Fatal(err)
  }
}
```

## [Golang Object Relational Mapping](https://github.com/jinzhu/gorm) \(GORM\)

Every modern day application needs an ORM. Every major framework provides it, e.g. Django & Rails. However, I personally advise against using ORM, unless time is a limiting resource to you, which is the case here. Currently `gorm` is the most popular ORM for Golang, at least according to GitHub stars.

`gorm` is still lacking in features compared to **ActiveRecord**, but for most cases it is good enough. If performance is an issue, try to write your own raw SQL query. `gorm` provides the following key benefits.

* Associations
* Eager Loading
* Hooks
* Transactions
* SQL Builder
* Auto Migrations

Here's a simple example on how to use `gorm`. More detailed usage will be discussed in the user authentication project.

```go
type Dog struct {
  gorm.Model
  Name string `gorm:"column:name"`
  Age  int    `gorm:"column:age"`
}

var databaseAddr = fmt.Sprintf(
  "postgresql://%s:%s@%s:%s/%s?%s", 
  user,
  password,
  host,
  port,
  database,
  sslMode,
)

func main() {
  db, err := gorm.Open("postgres", databaseAddr)
  
  db.AutoMigrate(&Dog{})

  d := Dog{
    Name: "Loki",
    Age: 6,
  }

  if err := db.Create(&d).Error; err != nil {
    log.Fatal(err)
  }
}
```

## [Migration](https://github.com/golang-migrate/migrate)

Although `gorm` provides auto migration, I still prefer writing the migrations manually so that I can up or down migrate to any version I want. Sometimes you gotta appreciate raw SQL a bit. SQL itself is a domain specific language and fairly human readable.

Running migration is pretty easy with `golang-migrate/migrate`. First, you need to create your SQL files, e.g.

```sql
CREATE TABLE dogs(
  id PRIMARY SERIAL,
  created_at TIMESTAMP WITH TIME ZONE,
  name VARCHAR(255),
  age INTEGER
);
```

Put all the files into a migrations directory and then run migration in Go.

```go
func main() {
  dir := "file://./migrations"
  psql := "postgresql://cfeng:cfeng@localhost:5432/myapp?sslmode=disable"

  migration, err := migrate.New(dir, psql)
  if err != nil {
    fmt.Println(err)
    return
  }

  if err := migration.Up(); err != nil {
    fmt.Println(err)
    return
  }
}
```

