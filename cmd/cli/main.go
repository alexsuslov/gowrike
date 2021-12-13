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

var DEBUG = false

type Cli struct {
	env          *string
	query        *string
	contacts     *bool
	users        *bool
	groups       *bool
	invitations  *bool
	create       *bool
	account      *bool
	workflows    *bool
	tasks        *bool
	comments     *bool
	taskComments *bool
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

		flag.Bool("workflows", false,
			"gowrike -workflows"),

		flag.Bool("tasks", false,
			"gowrike -tasks -query {taskId},{taskId},..."),

		flag.Bool("comments", false,
			"gowrike -comments"),

		flag.Bool("task_comments", false,
			"gowrike -task_comments -query {taskId}"),
	}
	flag.Parse()

	if err := godotenv.Load(*cli.env); err != nil {
		logrus.Warningf("no %s file", *cli.env)
	}
	//gowrike.DEBUG = true
	cli.run()
}

func (cli Cli) run() {
	cli.
		Contacts().
		Groups().
		Users().
		Invitations().
		Account().
		Workflows().
		Tasks().
		TaskComments().
		Comments().
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

func (cli *Cli) TaskComments() *Cli {
	if *cli.taskComments {
		if DEBUG {
			log.Println("task comments")
		}
		return cli.Done(gowrike.TaskCommentsRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Comments() *Cli {
	if *cli.comments {
		if DEBUG {
			log.Println("Comments")
		}
		return cli.Done(gowrike.CommentsRaw(context.Background()))
	}
	return cli
}

func (cli *Cli) Tasks() *Cli {
	if *cli.tasks {
		if DEBUG {
			log.Println("Tasks")
		}
		return cli.Done(gowrike.TasksRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Workflows() *Cli {
	if *cli.workflows {
		if DEBUG {
			log.Println("Workflows")
		}
		return cli.Done(gowrike.WorkflowsRaw(context.Background()))
	}
	return cli
}

func (cli *Cli) Account() *Cli {
	if *cli.account {
		if DEBUG {
			log.Println("Account")
		}
		return cli.Done(gowrike.AccountRaw(context.Background()))
	}
	return cli
}

func (cli *Cli) Invitations() *Cli {
	if *cli.invitations {
		if DEBUG {
			log.Println("Invitations")
		}
		return cli.Done(gowrike.InvitationsRaw(context.Background()))
	}
	return cli
}

func (cli *Cli) Groups() *Cli {
	if *cli.groups {
		if DEBUG {
			log.Println("Groups")
		}
		return cli.Done(gowrike.GroupsRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Users() *Cli {
	if *cli.users {
		if DEBUG {
			log.Println("Users")
		}
		return cli.Done(gowrike.UsersRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Contacts() *Cli {
	if *cli.contacts {
		if DEBUG {
			log.Println("Contacts")
		}
		cli.Done(gowrike.ContactsRaw(context.Background(), cli.query))
	}
	return cli
}

func (cli *Cli) Create() *Cli {
	if *cli.create {
		if DEBUG {
			log.Println("Create")
		}
		return cli.Done(gowrike.CreateRaw(context.Background(), cli.query, os.Stdin))
	}
	return cli
}
