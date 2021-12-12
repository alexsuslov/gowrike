package main

import (
	"context"
	"flag"
	"github.com/alexsuslov/godotenv"
	"github.com/alexsuslov/gowrike"
	"github.com/sirupsen/logrus"
	"io"
	"log"
	"os"
)

var config string
var create *bool

type Cli struct {
	env         *string
	query       *string
	contacts    *bool
	users       *bool
	groups      *bool
	invitations *bool
	create      *bool
	account     *bool
}

func main() {
	cli := Cli{
		flag.String("env", ".env",
			"gowrike -config default.env"),

		flag.String("query", "",
			"gowrike -query {query}"),

		flag.Bool("contacts", false,
			"gowrike -contacts -query {contactId},{contactId},..."),

		flag.Bool("users", false,
			"gowrike -users -query {userId}"),

		flag.Bool("groups", false,
			"gowrike -groups -query {groupId}"),

		flag.Bool("invitations", false,
			"gowrike -invitations"),

		flag.Bool("create", false,
			"cat ticket.json | gowrike -create -query {folderId}"),

		flag.Bool("account", false,
			"gowrike -account"),
	}
	flag.Parse()

	if err := godotenv.Load(*cli.env); err != nil {
		logrus.Warningf("no %s file", *cli.env)
	}
	gowrike.DEBUG = true
	cli.run()
}

func (cli Cli) run() {
	cli.
		Contacts().
		Groups().
		Users().
		Invitations().
		Account().
		Create()
	os.Exit(0)
}

func (cli *Cli) Done(body io.ReadCloser, err error) *Cli {
	if err != nil {
		logrus.Error(err)
		os.Exit(0)
	}
	defer body.Close()
	if _, err := io.Copy(os.Stdout, body); err != nil {
		panic(err)
	}
	return cli
}

func (cli *Cli) Account() *Cli {
	if *cli.account {
		log.Println("Account")
		return cli.Done(gowrike.AccountRaw(context.Background()))
	}
	return cli
}

func (cli *Cli) Invitations() *Cli {
	if *cli.invitations {
		log.Println("Invitations")
		return cli.Done(gowrike.InvitationsRaw(context.Background()))
	}
	return cli
}

func (cli *Cli) Groups() *Cli {
	if *cli.groups {
		log.Println("Groups")
		return cli.Done(gowrike.GroupsRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Users() *Cli {
	if *cli.users {
		log.Println("Users")
		return cli.Done(gowrike.UsersRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Contacts() *Cli {
	if *cli.contacts {
		log.Println("Contacts")
		cli.Done(gowrike.ContactsRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Create() *Cli {
	if *cli.create {
		log.Println("Create")
		return cli.Done(gowrike.CreateRaw(context.Background(), cli.query, os.Stdin))
	}
	return cli
}
