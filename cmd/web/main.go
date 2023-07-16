package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/CloudyKit/jet/v6"
	"github.com/alexedwards/scs/postgresstore"
	"github.com/alexedwards/scs/v2"
	_ "github.com/lib/pq"
	"github.com/muchiri08/hnews/models"
	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	appName string
	server  server
	debug   bool
	errLog  *log.Logger
	infoLog *log.Logger
	view    *jet.Set
	session *scs.SessionManager
	Models  models.Models
}

type server struct {
	host string
	port string
	url  string
}

func main() {

	migrate := flag.Bool("migrate", false, "should migrate - drop all tables")
	flag.Parse()

	server := server{
		host: "localhost",
		port: "8009",
		url:  "http://localhost:8080",
	}

	//init postgres db
	datasourceName := "postgres://root:KiNuThiaPro$2@localhost/hackernews?sslmode=disable"
	database, err := openDb(datasourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	//init upper/db
	upper, err := postgresql.New(database)
	if err != nil {
		log.Fatal(err)
	}
	defer func(upper db.Session) {
		err := upper.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(upper)

	//run migration scripts
	if *migrate {
		fmt.Println("Running migration...")
		if err := runMigrationScript(upper); err != nil {
			log.Fatal(err)
		}
		fmt.Println("Done running migration.")
	}

	//init our application
	app := application{
		server:  server,
		appName: "Hacker News",
		debug:   true,
		infoLog: log.New(os.Stdout, "INFO\t", log.Ltime|log.Ldate|log.Lshortfile),
		errLog:  log.New(os.Stderr, "ERROR\t", log.Ltime|log.Ldate|log.Llongfile),
		Models:  models.New(upper),
	}

	//init jet template
	if app.debug {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"), jet.InDevelopmentMode())
	} else {
		app.view = jet.NewSet(jet.NewOSFileSystemLoader("./views"))

	}

	//init session
	app.session = scs.New()
	app.session.Lifetime = 24 * time.Hour
	app.session.Cookie.Persist = true
	app.session.Cookie.Name = app.appName
	app.session.Cookie.Domain = app.server.host
	app.session.Cookie.SameSite = http.SameSiteStrictMode
	app.session.Store = postgresstore.New(database)

	//start the server
	if err := app.listenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func openDb(datasourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", datasourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func runMigrationScript(db db.Session) error {
	script, err := os.ReadFile("./migrations/tables.sql")
	if err != nil {
		return err
	}

	_, err = db.SQL().Exec(string(script))

	return err
}
