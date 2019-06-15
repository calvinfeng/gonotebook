# Build User Authentication

## Postgres Setup

### Create role & database

I am going to use PostgreSQL for this project, so let's create a database. The superuser on my computer is `cfeng`. If you don't have a role or wish to create a separate role for this project, then just do the following

```text
$ psql postgres
postgres=# create role <name> superuser login;
```

Create a database named `go_academy_userauth` with owner pointing to whichever role you like. I am using cfeng on my computer.

```text
$ psql postgres
postgres=# create database go_academy_userauth with owner=cfeng;
```

Actually just in case you don't remember the password to your `ROLE`, do the following

```text
postgres=# alter user <your_username> with password <whatever you like>
```

### Example

Create `cfeng`

```text
postgres=# create role cfeng superuser login;
```

Create database

```text
postgres=# create database go_academy_userauth with owner=cfeng;
```

Update password

```text
postgres=# alter user cfeng with password 'cfeng';
```

## Project Dependencies

I am going to introduce couple new open source libraries for this project:

* `spf13/cobra`
* `sirupsen/logrus`
* `labstack/echo`
* `jinzhu/gorm`

## Project User Authentication

This project is going to use JWT based authentication instead of the traditional session based authentication. For the detailed comparison of the two, take a look at this [article](https://medium.com/@sherryhsu/session-vs-token-based-authentication-11a6c5ac45e4). 

{% hint style="info" %}
This project is not going to use real JWT for simplicity sake, instead it uses a mock token. A real token can be generated using verified open source libraries, for more information go to [JWT.io](https://jwt.io/).
{% endhint %}

Nevertheless, it is important to show how JWT authentication works.

![](../.gitbook/assets/jwt.png)

### Authentication Endpoints

We expose endpoints on the server for users to sign in \(authenticate\) or register an account. The server is responsible for returning a token to the client side. Once user receives the token, he/she will embed this token into his/her request headers when he/she needs to access resources from the user. In practice, we would have two different services to complete the whole authentication cycle but we will combine the two services into one for this project. 

{% api-method method="post" host="http://localhost:8000" path="/api/authenticate/" %}
{% api-method-summary %}
Authenticate
{% endapi-method-summary %}

{% api-method-description %}
Authenticate a user
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-body-parameters %}
{% api-method-parameter name="email" type="string" required=true %}
email of the account
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=true %}
password of the account
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "name": "Calvin Feng",
    "email": "cfeng@goacademy.com",
    "jwt_token": "qcubXgJQRIHCon0b25HnuSOgaaw="

```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="http://localhost:8000" path="/api/register/" %}
{% api-method-summary %}
Register
{% endapi-method-summary %}

{% api-method-description %}
Register a new user
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-body-parameters %}
{% api-method-parameter name="name" type="string" required=true %}
name of the new user
{% endapi-method-parameter %}

{% api-method-parameter name="email" type="string" required=true %}
email of the new user
{% endapi-method-parameter %}

{% api-method-parameter name="password" type="string" required=true %}
password of the new user
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=201 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "id": 1,
    "name": "Calvin Feng",
    "email": "cfeng@goacademy.com",
    "jwt_token": "qcubXgJQRIHCon0b25HnuSOgaaw="
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

Once front end receives the token, it should put it into local storage. Next time when user opens his/her browser, the client application should use the same token to verify that user is indeed signed in.

### User Resource Endpoints

Since we are not using real JWT token, front end needs to make requests to server to ask for user information. We need to expose some user endpoints for that.

{% api-method method="get" host="http://localhost:8000" path="/api/users/" %}
{% api-method-summary %}
Users
{% endapi-method-summary %}

{% api-method-description %}
Fetch the list of all users
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Token" type="string" required=true %}
token to authenticate with server
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
[
    {
        "id": 1,
        "name": "Alice",
        "email": "alice@goacademy.com"
    },
    {
        "id": 2,
        "name": "Bob",
        "email": "bob@goacademy.com"
    },
    {
        "id": 3,
        "name": "Calvin",
        "email": "calvin@goacademy.com"
    }
]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="http://localhost:8000" path="/api/users/current/" %}
{% api-method-summary %}
Current User
{% endapi-method-summary %}

{% api-method-description %}
Fetch current user
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Token" type="string" required=true %}
token to authenticate with server
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "id": 4,
    "name": "Calvin Feng",
    "email": "cfeng@goacademy.com",
    "jwt_token": "qcubXgJQRIHCon0b25HnuSOgaaw="
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

### Message Resource Endpoints

Users can send messages to each other. When a user signs in to the system, he/she should be able to view all the sent and received messages. 

{% api-method method="post" host="http://localhost:8000" path="/api/messages/" %}
{% api-method-summary %}
Send Message
{% endapi-method-summary %}

{% api-method-description %}
Send a message to another user
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Token" type="string" required=true %}
token to authenticate with server
{% endapi-method-parameter %}
{% endapi-method-headers %}

{% api-method-body-parameters %}
{% api-method-parameter name="receiver\_id" type="integer" required=true %}
integer ID of the message receiver
{% endapi-method-parameter %}

{% api-method-parameter name="body" type="string" required=true %}
message body
{% endapi-method-parameter %}
{% endapi-method-body-parameters %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=201 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
{
    "id": 3,
    "created_at": "2019-06-15T01:03:21.460369322-07:00",
    "body": "Hello there",
    "sender_id": 4,
    "receiver_id": 1
}
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

Users can see their own list of received and sent messages.

{% api-method method="get" host="http:localhost:8000" path="/api/messages/sent/" %}
{% api-method-summary %}
Sent Messages
{% endapi-method-summary %}

{% api-method-description %}
Fetch the list of current user's sent messages
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Token" type="string" required=true %}
token to authenticate with server
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
[
    {
        "id": 3,
        "created_at": "2019-06-15T01:03:21.460369-07:00",
        "body": "Hello World",
        "sender_id": 4,
        "sender": {
            "id": 4,
            "name": "Calvin",
            "email": "cfeng@goacademy.com"
        },
        "receiver_id": 1,
        "receiver": {
            "id": 1,
            "name": "Alice",
            "email": "alice@goacademy.com"
        }
    }
]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

{% api-method method="get" host="http://localhost:8000" path="/api/messages/received/" %}
{% api-method-summary %}
Received Messages
{% endapi-method-summary %}

{% api-method-description %}
Fetch the list of current user's received messages
{% endapi-method-description %}

{% api-method-spec %}
{% api-method-request %}
{% api-method-headers %}
{% api-method-parameter name="Token" type="string" required=true %}
token to authenticate with server
{% endapi-method-parameter %}
{% endapi-method-headers %}
{% endapi-method-request %}

{% api-method-response %}
{% api-method-response-example httpCode=200 %}
{% api-method-response-example-description %}

{% endapi-method-response-example-description %}

```javascript
[
    {
        "id": 1,
        "created_at": "2019-05-30T00:00:25.335622-07:00",
        "body": "Hello Alice",
        "sender_id": 3,
        "sender": {
            "id": 3,
            "name": "Calvin",
            "email": "cfeng@goacademy.com"
        },
        "receiver_id": 1,
        "receiver": {
            "id": 1,
            "name": "Alice",
            "email": "alice@goacademy.com"
        }
    },
    {
        "id": 2F,
        "created_at": "2019-06-15T01:03:21.460369-07:00",
        "body": "Hello World",
        "sender_id": 3,
        "sender": {
            "id": 3,
            "name": "Calvin",
            "email": "cfeng@goacademy.com"
        },
        "receiver_id": 1,
        "receiver": {
            "id": 1,
            "name": "Alice",
            "email": "alice@goacademy.com"
        }
    }
]
```
{% endapi-method-response-example %}
{% endapi-method-response %}
{% endapi-method-spec %}
{% endapi-method %}

### Models

There are two primary resources on our server, i.e. users and messages.

```go
// User is a user model.
type User struct {
	// Both
	ID       uint   `gorm:"column:id"          json:"id"`
	Name     string `gorm:"column:name"        json:"name" `
	Email    string `gorm:"column:email"       json:"email"`
	JWTToken string `gorm:"column:jwt_token"   json:"jwt_token,omitempty"`

	// JSON only
	Password string `sql:"-" json:"password,omitempty"`

	// Database only
	CreatedAt      time.Time `gorm:"column:created_at"      json:"-"`
	UpdatedAt      time.Time `gorm:"column:updated_at"      json:"-"`
	PasswordDigest []byte    `gorm:"column:password_digest" json:"-"`
}
```

```go
// Message is a model for messages.
type Message struct {
	ID        uint      `gorm:"column:id"          json:"id"`
	CreatedAt time.Time `gorm:"column:created_at"  json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"  json:"-"`
	Body      string    `gorm:"column:body"        json:"body"`

	// Foreign keys
	SenderID   uint  `gorm:"column:sender_id"       json:"sender_id"`
	Sender     *User `gorm:"foreignkey:sender_id"   json:"sender,omitempty"`
	ReceiverID uint  `gorm:"column:receiver_id"     json:"receiver_id"`
	Receiver   *User `gorm:"foreignkey:receiver_id" json:"receiver,omitempty"`
}
```

### Migrations

Instead of using auto migration feature of GORM, I prefer to write our own SQL because it gives us greater flexibility and better organization.

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    name VARCHAR(255),
    email VARCHAR(255),
    jwt_token VARCHAR(255),
    password_digest BYTEA
);

CREATE UNIQUE INDEX ON users(name);
CREATE UNIQUE INDEX on users(email);
CREATE UNIQUE INDEX ON users(jwt_token);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    sender_id INTEGER REFERENCES users(id),
    receiver_id INTEGER REFERENCES users(id),
    body TEXT
);

CREATE INDEX ON messages(sender_id);
CREATE INDEX ON messages(receiver_id);
```

### Put everything together

Let's start off with creating a skeleton for the project. Details will be discussed in the video section. Create a project file structure as follows.

```text
go-academy/
    userauth/
        cmd/
            run_migrations.go
            run_server.go
        handler/
        migrations/
        model/
        public/
        main.go
```

Inside the `main.go`, set up logging and cobra commands.

{% code-tabs %}
{% code-tabs-item title="main.go" %}
```go
package main

import (
	"os"

	"github.com/calvinfeng/go-academy/userauth/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func main() {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logrus.SetOutput(os.Stdout)

	logrus.SetLevel(logrus.DebugLevel)

	root := &cobra.Command{
		Use:   "userauth",
		Short: "user authentication service",
	}

	root.AddCommand(cmd.RunMigrationsCmd, cmd.RunServerCmd)
	if err := root.Execute(); err != nil {
		logrus.Fatal(err)
	}
}
```
{% endcode-tabs-item %}
{% endcode-tabs %}

Create two commands inside `cmd` package, one for migration and one for running server.

{% code-tabs %}
{% code-tabs-item title="run\_server.go" %}
```go
package cmd

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/jinzhu/gorm"
	"github.com/spf13/cobra"

	// Driver for Postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// RunServerCmd is a command to run server from terminal.
var RunServerCmd = &cobra.Command{
	Use:   "runserver",
	Short: "run user authentication server",
	RunE:  runServer,
}

func runServer(cmd *cobra.Command, args []string) error {
	conn, err := gorm.Open("postgres", pgAddr)
	if err != nil {
		return err
	}

	srv := echo.New()

	srv.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "HTTP[${time_rfc3339}] ${method} ${path} status=${status} latency=${latency_human}\n",
		Output: io.MultiWriter(os.Stdout),
	}))

	srv.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	srv.File("/", "public/index.html")
	srv.Static("/assets", "public/assets")
	if err := srv.Start(":8080"); err != nil {
		return err
	}

	return nil
}
```
{% endcode-tabs-item %}
{% endcode-tabs %}

```go
package cmd

import (
	"fmt"

	"github.com/golang-migrate/migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	
	_ "github.com/lib/pq" // Driver
	_ "github.com/golang-migrate/migrate/database/postgres" // Driver
	_ "github.com/golang-migrate/migrate/source/file"       // Driver
)

const (
	host         = "localhost"
	port         = "5432"
	user         = "cfeng"
	password     = "cfeng"
	database     = "go_academy_userauth"
	ssl          = "sslmode=disable"
	migrationDir = "file://./migrations/"
)

var log = logrus.WithFields(logrus.Fields{
	"pkg": "cmd",
})

var pgAddr = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?%s", user, password, host, port, database, ssl)

// RunMigrationsCmd is a command to run migration.
var RunMigrationsCmd = &cobra.Command{
	Use:   "runmigrations",
	Short: "run migration on database",
	RunE:  runMigrations,
}

func runMigrations(cmd *cobra.Command, args []string) error {
	migration, err := migrate.New(migrationDir, pgAddr)
	if err != nil {
		return err
	}

	log.Info("performing reset on database")
	if err = migration.Drop(); err != nil {
		return err
	}

	if err := migration.Up(); err != nil {
		return err
	}

	log.Info("migration has been performed successfully")
}
```

### Video

Now your project should be able to compile and you should be able to run migration on the database.

```text
go install && userauth runmigrations
```

In the video, I will discuss how to write each endpoint.

## Bonus & Additional Resource

If you want to learn more about session storage, security, encryption, and many other topics relating to web applications, take a look at this [GitBook](https://astaxie.gitbooks.io/build-web-application-with-golang/content/en/).

## Source

[GitHub](https://github.com/calvinfeng/go-academy/tree/master/userauth)

