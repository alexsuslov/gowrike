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
  -contacts
        gowrike -contacts -query {contactId},{contactId},...
  -create
        cat ticket.json | gowrike -create -query {folderId}
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
  -tasks
        gowrike -tasks -query {taskId},{taskId},...
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
