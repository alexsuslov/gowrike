# GOWRIKE

## Cli

gowrike -help

```
Usage of gowrike:
  -contacts
        gowrike -contacts -query {contact_id},{contact_id},...
  -create
        cat ticket.json | gowrike -create -folder_id {folder_id}
  -env string
        gowrike -config default.env (default ".env")
  -folder_id string
        gowrike -folder_id {folder_id}
  -query string
        gowrike -query {query}
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