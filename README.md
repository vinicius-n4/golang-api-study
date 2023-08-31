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

| Key       | Type     | Description                      |
| :-------- | :------- | :------------------------------- |
| `name`    | `string` | **Required**. A name of a person |

#### Update item

```http
  PUT /update/${id}
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `string` | **Required**. Id of item to update |

#### Delete item

```http
  DELETE /delete/${id}
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `string` | **Required**. Id of item to delete |

## Usage/Examples
// TODO: update documentation with docker command

Open a terminal and, in the repository directory, run the following command to start the server that will be listening `http://localhost:8000`.

```sh
$ go run main.go
```

### List items

```sh
$ curl --location --request GET 'http://localhost:8000/list'
```

Return:
```json
{
    "0": {
        "name": "Vinicius"
    },
    "1": {
        "name": "Nogueira"
    },
    "2": {
        "name": "Costa"
    }
}
```

### Create item

```sh
$ curl --location --request POST 'http://localhost:8000/create' \
--form 'name="Vinicius"'
```

Return:
```json
{
    "name": "Vinicius"
}
```

### Update item

```sh
$ curl --location --request PUT 'http://localhost:8000/update/0' \
--form 'name="Diogo"'
```

Return:

`Status Code: 200 OK`
```json
{
    "name": "Diogo"
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
$ curl --location --request DELETE 'http://localhost:8000/delete/0' \
--form 'name="João"'
```

Return:

`Status Code: 200 OK`
```json
{
    "name": ""
}
```

`Status Code: 400 Bad Request`
```json
{
    "message": "ID doesn't exist. Try to list items before delete them."
}
```
