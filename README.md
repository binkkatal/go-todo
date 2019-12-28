# GO-TODO setup

### Requirement 
  requires https://github.com/golang-migrate/migrate cli to be installed on the system

## Create table and run migrations

Run `make test` to create a fresh table and run migrations.
this also runs the test cases.

## Running the api

Prior to running , execute the env.sh file to set the env variables
Run `make server` to build the server.
This will build the binary and place it in `/bin/http`.
After the binary is built, execute it and the server will start running.

## API Endpoints

### Create

```
 URL: http://localhost:9000/api/v1/todo
 METHOD: POST
 params:
  {
      "title": "Start working",
      "note": "The tasks should be finished by midnight",
      "due_date": "2000-01-01T00:00:00Z"
  }
```

### Update

```
 URL: http://localhost:9000/api/v1/todo/{ID}
 METHOD: PUT
 params:
  {
      "title": "Do more work",
      "note": "The tasks will be delayed",
      "due_date": "2000-01-01T00:00:00Z"
  }
```

### GET

```
 URL: http://localhost:9000/api/v1/todo/{ID}
 METHOD: GET
```

### DELETE

```
 URL: http://localhost:9000/api/v1/todo/{ID}
 METHOD: DELETE
```
