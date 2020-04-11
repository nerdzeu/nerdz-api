# NERDZ API

To autogenerate the the `../rest/responses.go` file, you can run:

```sh
rts -server http://localhost:9090/v1 -pkg rest -headers "Authorization: Bearer _uZV-FCsS3-ytssqZC6qLw" -routes routes.txt -out ../rest/responses.go
```

The file, than, must be manually edited.

The command above uses the `routes.txt` file and an access token present in the test database.
`rts` is https://github.com/galeone/rts
