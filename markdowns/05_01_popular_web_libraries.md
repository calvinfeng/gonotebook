# Popular Web Libraries

As our application gets larger and more complex, we need help from some open source libraries to make
our life easier. There are many good web frameworks out there but I am going to show you the couple
familiar libraries I've been using at work.

## [Gorilla Mux](https://github.com/gorilla/mux)

Mux stands for HTTP request multiplexer. It is essentially a router that routes requests to the
appropriate request handler in your application. Mux library offers URL pattern matching, query
params patter matching, URL host matching and the list goes on.

For example:

```go
r := mux.NewRouter()
r.HandleFunc("/robots/{name}", RobotHandler)
r.HandleFunc("/maps/destinations/{id:[0-9]+}", DestinationHandler)
```

Notice that key and category will become available as a variable through mux router pattern matching.
If I were to send a request to the `/robots/bender/` endpoint, then `name` will hold the value
`bender`. I can extract this value using `mux.Vars(r)`.

```go
func RobotHandler(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  fmt.Println(vars["name"])
}
```

## [Logger](https://github.com/sirupsen/logrus)

Logging is crucial for debugging applications in production. Go provides a default `log` package.
However, sometimes you'd want more. `logrus` provides colorful logging and additional log fields to
identify the source of errors. More importantly, `logrus` provides integration with [Sentry](https://sentry.io/welcome/).

## [Command-line Interface](https://github.com/spf13/cobra) (CLI)

So far we've been building applications that runs a single command, i.e. whatever you put in your
main function. What if we want to have more command like the following.

    go install
    myapp runmigration
    myapp runserver
    myapp reset
    ...

We can accomplish that using the `cobra` package. We define a list of commands and each command is
mapping to an execution function, much like `main()`.

```go
var root = &cobra.Command{
  Use: "myapp",
  Short: "my toy application"
}

var runmigrationCmd = &cobra.Command{
  Use:   "runmigration",
  Short: "run migration on database",
  RunE:  runmigration,
}

var runserverCmd = &cobra.Command{
  Use:   "runserver",
  Short: "run user authentication server",
  RunE:  runserver,
}

// Execute configures command and executes them.
func Execute() {
  // Run it!
  rootCmd.AddCommand(runserverCmd, runmigrationCmd)
  if err := rootCmd.Execute(); err != nil {
    os.Exit(1)
  }
}
```

In our main function, we just need to run the `Execute` function.

```go
func main() {
  Execute()
}
```

## [Golang Object Relational Mapping](https://github.com/jinzhu/gorm) (GORM)

Every modern day application needs an ORM. Every major framework provides it, e.g. Django & Rails.
However, I personally advise against using ORM, unless time is a limiting resource to you, which is
the case here. Currently `gorm` is the most popular ORM for Golang, at least according to GitHub stars.

`gorm` is still lacking in features compared to **ActiveRecord**, but for most cases it is good enough.
If performance is an issue, try to write your own raw SQL query. `gorm` provides the following key beneifts.

* Associations
* Eager Loading
* Hooks
* Transactions
* SQL Builder
* Auto Migrations

Here's a simple example on how to use `gorm`. More detailed usage will be discussed in the user
authentication project.

```go
type Dog struct {
  gorm.Model
  Name string `gorm:"type:varchar(255); column:name"`
  Age  int    `gorm:"type:integer;      column:age"`
}

func main() {
  db, err := gorm.Open("postgres",
    fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, sslMode),
  )

  db.AutoMigrate(&Dog{})

  d := Dog{
    Name: "Loki",
    Age: 6,
  }

  if err := db.Create(&d).Error; err != nil {
    fmt.Println(err)
  }
}
```

## [Migration](https://github.com/golang-migrate/migrate)

Although `gorm` provides auto migration, I still prefer writing the migrations manually so that I can
up or down migrate to any version I want. Also, writing raw SQL isn't so bad after all. SQL itself is
a pretty high level language and extremely human readable.

Running migration is pretty easy with `golang-migrate/migrate`. First, you need to create your SQL
files, e.g.

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