# GO GraphQL - Hackernews demo
GOlang GraphQL training from HowToGraphQL.com

Following the tutorial from [How To GraphQL for GO](https://www.howtographql.com/graphql-go/0-introduction/)


# Usage

## MySQL

MySQL is provided with a `docker-compose.yaml` file on the repository root.
Start the DB with the following command:

    docker-compose up -d

## Service 

Navigate to `cmd/server` and start he server

    cd cmd/server
    go run main-go

Then navigate to http://localhost:8080/ to open the GraphQL playground.

## Use the GraphQL Server

### Create User

**Request**
```graphql
mutation {
    createUser(input: {username: "user1", password: "123"})
}
```

**Response**
```json
{
  "data": {
    "createUser": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMDI0NDMsInVzZXJuYW1lIjoiam9obiJ9.SyjeM8lzdIpM5HS6l68tR0yb744gXoSdSrhKCAHRlNo"
  }
}
```

The `createUser` response contains a JWT token, that should be used in the `Authorization` header like this:

```json
{
  "Authorization": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMDI0NDMsInVzZXJuYW1lIjoiam9obiJ9.SyjeM8lzdIpM5HS6l68tR0yb744gXoSdSrhKCAHRlNo"
}
```
## Login User

**Request**
```graphql
mutation {
    login(input: {username: "user1", password: "123"})
}
```

**Response**
```json
{
  "data": {
    "login": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMTA1NjIsInVzZXJuYW1lIjoidXNlcjEifQ.9cNmB5Nzyj5-rH6YcOXOt6wzw8QmEsNAJxGnViDcK9E"
  }
}
```

This response holds the new JWT token to be used in the `Authorization` header.

## Refresh User Token

**Request**
```graphql
mutation {
    refreshToken(input: {token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMTA1NjIsInVzZXJuYW1lIjoidXNlcjEifQ.9cNmB5Nzyj5-rH6YcOXOt6wzw8QmEsNAJxGnViDcK9E"})
}
```

**Response**
```json
{
  "data": {
    "refreshToken": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODAwMTA2ODksInVzZXJuYW1lIjoidXNlcjEifQ.9IULnunhHm9lwoo03Qri4SFMDJGC7HyjHH_gOpaDPNE"
  }
}
```

This response holds the new JWT token to be used in the `Authorization` header.

## Create Links

**Request**
```graphql
mutation {
  createLink(input: {title: "real link!", address: "www.graphql.org"}){
    title
    address
    user{
      name
    }
  }
}
```

**Response**
```json
{
  "data": {
    "createLink": {
      "title": "real link!",
      "address": "www.graphql.org",
      "user": {
        "name": "user1"
      }
    }
  }
}
```

If you do not have a valid token, then you will get an `Access Denied` error

## Get Links

**Request**
```graphql
query {
  links {
    title
    address
  }
}
```

**Response**
```json
{
  "data": {
    "links": [
      {
        "title": "real link!",
        "address": "www.graphql.org"
      }
    ]
  }
}
```