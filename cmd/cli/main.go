package main

import (
	"context"
	"flag"
	"github.com/alexsuslov/godotenv"
	"github.com/alexsuslov/gowrike"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

var config string
var create *bool

type Cli struct {
	env      *string
	query    *string
	contacts *bool
	users    *bool
	folderID *string
	create   *bool
}

func main() {
	cli := Cli{
		flag.String("env", ".env",
			"gowrike -config default.env"),

		flag.String("query", "",
			"gowrike -query {query}"),

		flag.Bool("contacts", false,
			"gowrike -contacts -query {contact_id},{contact_id},..."),

		flag.Bool("users", false,
			"gowrike -users -query {user_id}"),

		flag.String("folder_id", "",
			"gowrike -folder_id {folder_id}"),

		flag.Bool("create", false,
			"cat ticket.json | gowrike -create"),
	}
	flag.Parse()

	if err := godotenv.Load(*cli.env); err != nil {
		logrus.Warningf("no %s file", *cli.env)
	}

	cli.run()
}

func (cli Cli) run() {
	cli.
		Contacts().
		Users().
		Create()
}

func (cli *Cli) Done(body io.ReadCloser, err error) *Cli {
	if err != nil {
		panic(err)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err != nil {
		panic(err)
	}
	return &Cli{}
}

func (cli *Cli) Users() *Cli {

	if cli.users == nil {
		return cli
	}

	if cli.query == nil {
		return cli.Done(gowrike.UsersRaw(context.Background()))
	}

	ids := strings.Split(*cli.query, ",")
	return cli.Done(gowrike.ContactsRaw(context.Background(), ids...))

}

func (cli *Cli) Contacts() *Cli {

	if cli.contacts == nil {
		return cli
	}

	if cli.query == nil {
		return cli.Done(gowrike.ContactsRaw(context.Background()))
	}

	ids := strings.Split(*cli.query, ",")
	return cli.Done(gowrike.ContactsRaw(context.Background(), ids...))

}

func (cli *Cli) Create() *Cli {
	if cli.create == nil || cli.folderID == nil {
		return cli
	}

	if !*cli.create {
		return cli
	}

	return cli.Done(gowrike.CreateRaw(context.Background(), *cli.folderID, os.Stdin))
}
