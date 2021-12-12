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
  -contacts
        gowrike -contacts -query {contact_id},{contact_id},...
  -create
        cat ticket.json | gowrike -create
  -env string
        gowrike -config default.env (default ".env")
  -folder_id string
        gowrike -folder_id {folder_id}
  -query string
        gowrike -query {query}
  -users
        gowrike -users -query {user_id}
```

## Contacts

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