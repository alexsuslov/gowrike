# GOWRIKE

## Cli

### build
```
go build -o bin/gowrike cmd/cli/main.go

```

### Help

gowrike -help

```
Usage of gowrike:
  -account
        gowrike -account
  -comments
        gowrike -comments
  -contact_timelogs
        gowrike -contact_timelogs -query {contactId}
  -contacts
        gowrike -contacts -query {contactId},{contactId},...
  -create
        cat ticket.json | gowrike -create -query {folderId}
  -dependencies
        gowrike -dependencies -query {dependencyId}
  -env string
        gowrike -config default.env (default ".env")
  -groups
        gowrike -groups -query {groupId}
  -invitations
        gowrike -invitations
  -query string
        gowrike -query {query}
  -task_comments
        gowrike -task_comments -query {taskId}
  -task_dependencies
        gowrike -task_dependencies -query {taskId}
  -tasks
        gowrike -tasks -query {taskId},{taskId},...
  -timelogs
        gowrike -timelogs
  -users
        gowrike -users -query {userId}
  -workflows
        gowrike -workflows

```

### Contacts

```
gowrike -contacts -query KUAKDR2J

{
  "kind": "contacts",
  "data": [
    {
      "id": "KUAKDR2J",
      "firstName": "Евгения",
...

```

### Users

```
gowrike -users -query KUAKSQGA

{
  "kind": "contacts",
  "data": [
    {
      "id": "KUAKSQGA",
      "firstName":
...
```
