# golang-api-study

This is my study repository about a functional Golang API service. This API allows a user to Create, List, Update and Delete items from a in memory database.

## API Reference

#### List items

```http
  GET /list
```

#### Create item

```http
  POST /create
```

| Key        | Type                  | Description                      |
|:-----------|:----------------------|:---------------------------------|
| `name`     | `form-data` | **Required**. A name of a person |
| `document` | `form-data` | **Required**. A document number  |

#### Update item

```http
  PUT /update/${id}
```

| Parameter   | Type               | Description                        |
|:------------|:-------------------|:-----------------------------------|
| `id`        | `endpoint vaiable` | **Required**. Id of item to update |
| `name`      | `form-data`        | **Required**. A name of a person   |
| `document`  | `form-data`        | **Required**. A document number    |

#### Delete item

```http
  DELETE /delete/${id}
```

| Parameter | Type                | Description                        |
| :-------- |:--------------------| :--------------------------------- |
| `id`      | `endpoint variable` | **Required**. Id of item to delete |

## Usage/Examples

Open a terminal and, in the repository directory, run the following command to start the database and the server that will be listening `http://localhost:8000`.

```sh
$ docker compose up
```

### List items

```sh
$ curl --location --request GET 'http://localhost:8000/list'
```

Return:
```json
[
    {   "id": 1,
        "name": "Vinicius",
        "document": "09876543210"
    },
    {   "id": 2,
        "name": "Nogueira",
        "document": "98765432109"
    },
    {   "id": 3,
        "name": "Costa",
        "document": "87654321098"
    }
]
```

### Create item

```sh
$ curl --location --request POST 'http://localhost:8000/create' \
--form 'name="Silva"' \
--form 'document="76543210987"'
```

Return:
```json
{   "id": 4,
    "name": "Silva",
    "document": "76543210987"
}
```

### Update item

```sh
$ curl --location --request PUT 'http://localhost:8000/update/4' \
--form 'name="Silva"' \
--form 'document="12345678901"'
```

Return:

`Status Code: 200 OK`
```json
{
  "id": 4,
  "name": "Silva",
  "document": "12345678901"
}
```

`Status Code: 400 Bad Request`
```json
{
    "message": "ID doesn't exist. Try to list items before update them."
}
```

### Delete item

```sh
$ curl --location --request DELETE 'http://localhost:8000/delete/4'
```

Return:

`Status Code: 200 OK`
```json
{
  "id": 4,
  "name": "Silva",
  "document": "12345678901"
}
```

`Status Code: 400 Bad Request`
```json
{
    "message": "ID doesn't exist. Try to list items before delete them."
}
```
