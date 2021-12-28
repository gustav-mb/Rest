# Micro Services

## How to run

### With Script

To run the program with ease, run the script ``run.ps1`` with PowerShell.

In the root directory write:

> ``./run.ps1``

This will open two separate terminals; one for the server, one for the client.

### Manually

To run the program manually do the following:

First start the server in the server directory ``./server`` with:

> ``go run .``

Then start the client in the client directory ``./client`` with:

> ``go run .``

For both the server and client, you can specify the address and port the server should run on and client connect to.

For the program to work, the client and the server must have the same ip address!

To specify an ip address write:

> ``go run . -address <address> -port <port>``

---

## End Points

The REST API has the following end points:

1. Students
2. Teachers
3. Courses

All of the endpoints support:

* ``GET`` by ``id`` and ``name``
* ``PUT`` by ``id``
* ``DELETE`` by ``id``

### Students

The students end point supports:

* students
* students/id/\<id\>
* students/name/\<name\>

### Teachers

The teachers end point supports:

* teachers
* teachers/id/\<id\>
* teachers/name/\<name\>

### Courses

The courses end point supports:

* courses
* courses/id/\<id\>
* courses/name/\<name\>

---

## Client Commands

To interact with the server, the client supports a series of commands to send requests to the server.

A command is made up in the following way:

> ``<CRUD method> <url> <body>``

Note: ``<body>`` can be omitted for ``GET`` and ``DELETE`` requests!

## Get

To e.g., ``GET`` all students write:

> ``get students``

To ``GET`` a student by id or name write:

> ``get students/id/<id>``
>
> ``get students/name/<name>``

To e.g, get the student with id 1, or the student with name "Gustav" write:

> ``get students/id/1``
>
> ``get students/name/Gustav``

## Post (create)

To ``POST`` (create) a new student write:

> ``post students <JSON>``

To e.g., create a new student named "Gustav" write:

> ``post students { "id": 0, "name": "Gustav", "enrollment": "Dropped out", "courseworkload": 10 }``

## Put (update)

To ``PUT`` (update) an existing student write:

> ``put student/id/<id> <JSON>``

To e.g., update an existing student write:

> ``put students/id/1 { "id": 0, "name": "Gustav", "enrollment": "Dropped out", "courseworkload": 10 }``

## Delete

To ``DELETE`` an existing student write:

> ``delete students/id/<id>``

To e.g., delete the student with id 1 write:

> ``delete students/id/1``
