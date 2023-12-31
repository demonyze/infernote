package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/demonyze/infernote/db/chords"
	"github.com/jackc/pgx/v5"

	"github.com/labstack/echo/v4"
)

func main() {

	url := "postgres://user:secret@localhost:5432/infernote"
	conn, err := pgx.Connect(context.Background(), url)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	queries := chords.New(conn)

	ch, err := queries.CreateChord(context.Background(), chords.CreateChordParams{
		Name: "Cmajor",
		Root: "C",
		Type: chords.TypeMajor,
		// Guitar: {},
	})
	if err != nil {
		fmt.Println(err)
	}
	log.Println(ch)
	// list all authors
	authors, err := queries.ListChords(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	log.Println(authors)

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
