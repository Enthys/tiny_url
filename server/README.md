# TinyUrl
## About
This application allows you to have short urls of any full URI you provide.

## Using the application
Many of the required operations are available as commands in the provided
Makefile.

### Starting the application
Before starting the application there are a few things which are required before
being able to start the application. One is to copy over the `.env.default` file
under the name `.env`. It is used by the `Makefile` and by the application in
case there is no `-db-dsn` flag provided. The other is to download the
`migrate` binary which is used for operations concerning database migrations.
Both of there operations have been streamlined by the `make run/setup` command.
```sh
make run/setup
```

You can use the `make run/dev` command to start the application using the
`docker-compose.yml` file.

__Alternatively__ you can run the application locally using the
`go run ./cmd/api` command
```sh
go run ./cmd/api -port=4000 -db-dsn='postgres://user:pass@addr:port/db_name?sslmode=disable' -cors-trusted-origins=frontend-addr:port
```

> Have in mind that before that before you attempt to use the application you
> will need to run the migrations.
> ```sh
> # Running the database from docker
> make db/migrations/up
> 
> # Running local instance
> make 
> ```

## Endpoints

### Go to Link
**Request**
```
@GET /url/:short_url
```
If there is an existing url that has the provided short url then you will be
redirected to the address sitting behind the short link. Otherwise you will
receive a status code of `404 Not Found` and `{"error": "resource not found"}`.

### Create ShortUrl
**Request:**
```
@POST /v1/short-urls
{
  "url": "https://google.com"
}
```
**Response**
```json
// Status Code: 201 Created
{
  "short-url": {
    "id": 1,
    "url": "https://google.com",
    "short_url": "1K2B3KL1HJB2K3J"
  }
}
```
> The `url` for which a short link has to be provided needs to be a full URI
> path(e.g. `http://example.com`, `https://google.com`). URIs which do not have
> a schema will be flagged as invalid.

> The returned `short_url` value can be used by the `Go to Link` route.

### Get ShortUrls
**Request**
```
@GET /v1/short-urls
```
```json
// Status Code: 200 Ok
{
  "short-urls": [
    {
      "id": 1,
      "url": "https://google.com",
      "short_url": "1K2B3KL1HJB2K3J"
    },
    {
      "id": 2,
      "url": "https://youtube.com",
      "short_url": "ASZKDJB234JKLB2"
    },
    {
      "id": 3,
      "url": "https://dropbox.com",
      "short_url": "23K5JBLKJ1K4JB1"
    },
  ]
}
```
You have the option of providing `page` and `page_size` as query arguments.
- `page`
  - Default value: 1
  - Minimum value: 1
  - Maximum value: 10 000 000
- `page_size`
  - Default value: 20
  - Minimum value: 1
  - Maximum value: 100

### Delete ShortUrl
**Request**
```
@DELETE /v1/short-urls/:id
```

**Response**
```json
// Status Code: 204 No Content
{}
```

If the provided `id` does not belong to any registered short url then a status
code of `404 Not Found` with body `{"error": "resource not found"}` will be
returned.

## Extras

### Quick database connection
The command `make db/connect` allows you to quickly connect to the database. For
this to work you will have to be running the application using `docker`.
```sh
make db/connect
```

### Create new migrations
The command `make db/migrations/new` allows us to more easily create new 
migrations for the application. The command requires us to provide a name for
the migration we want to create.
```
make db/migrations/up name=create_example_table
```
