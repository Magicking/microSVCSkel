# µSVC


-------

The goal of this project is to build a µSVC that read/write in mongoDB with a
basic API

### Part 0

 - Write the Dockerfile & docker-compose, see examples in this repository

### Part 1

Get a quick look at thoses 
 - Workspace hierarchy (https://golang.org/doc/code.html#Workspaces)
 - Vendor (https://github.com/kovetskiy/manul)
 - Documentation (https://godoc.org)

### Part 2

Write a simple OpenAPI specification with a /list that list collections and
/push in the mongodb that push some data into mongodb

### Part 3

 - Generate the boilerplate using ```swagger generate server -f docs/microsvc.yml```
 - Get dependancies (Use https://golanglibs.com or https://github.com search engine to find a MongoDB driver and PostgreSQL driver)
 - Find a library that suits our need

### Part 4

 - Write code to insert into mongoDB, for a complete tutorial see https://goswagger.io/tutorial/todo-list.html

### Part 5

 - Switch from MondoDB to PostgreSQL
 1. Search a new database driver or ORM to abstract database access (GORM will be used, see http://jinzhu.me/gorm/)
 2. Search for a Postgresql Docker image
 3. Change docker-compose file to use the Postgresql image and set environments variables
 4. Write code to handle argument (see go-flags documentation)
 5. Create an ```internal``` directory to put some database logic, see https://docs.google.com/document/d/1e8kOo3r51b2BWtTs_1uADIA5djfXhPT36s6eHVRIvaU/edit)
 6. Create a file ```database.go``` inside the internal directory
 7. This file should have 3 methods and an object (see GORM models definition) structure, InitDatabase, Insert, ListAll, that should respectively:
  - Make a connection to PostgreSQL and initialize the database with the defined object
  - Insert an object in the database
  - Retreive all objects in the database
 8. Change code in restapi/configure_... .go to setup the database and use Insert and ListAll to setup logic

### Don't panic

 - Update dependencies and update vendor
 - Use documentation, issue trackers, Stackoverflow and Google (in that order) extensively

### Usage

Start and build containers using ```docker-compose up --build``` while being in the base directory.
Containers can be interrupted either by typing Ctrl+C or ```docker-compose down```.

Manul use git submodules to track vendors dependencies therefor when cloning this
repository use ```git clone --recurse-submodules https://github.com/Magicking/microSVCSkel```
to clone a repository **WITH** vendors.

To fetch dependencies of your project in your GOPATH, use ```go get -d -v ./...```, to
copy dependencies in your vendor directory, use ```manul -I -r```.

Most used command sorted by usage:
 - docker-compose up --build
 - docker-compose down --volume
 - swagger generate server -f docs/microsvc.yml
 - go get -d -v ./...
 - manul -I -r
 - git clone --recurse-submodules https://github.com/Magicking/microSVCSkel

### Ressources

Frameworks:
- Docker & docker-compose (https://docker.github.io/engine/installation/ & https://docs.docker.com/compose/install/)
- MongoDB (github.com/globalsign/mgo)
- Swagger (https://goswagger.io)

Documentation:
 - Dockerfile specification reference: https://docs.docker.com/engine/reference/builder/
 - Compose-file specification reference: https://docs.docker.com/compose/compose-file/
 - Swagger/OpenAPI: https://github.com/OAI/OpenAPI-Specification/blob/master/versions/2.0.md
 - Vendors search engine: https://golanglibs.com
Golang ressources:
 - tl;dr: https://tour.golang.org/welcome/1
 - Idioms: https://gobyexample.com
 - How to write Go code: https://golang.org/doc/code.html
 - Effective Golang: https://golang.org/doc/effective_go.html
 - Search engine Golang library: https://golanglibs.com/
 - Book: The Go Programming Language (Alan A. A. Donovan, Brian W. Kernighan)
 - Gorm: http://jinzhu.me/gorm/
 - Go-flags: https://godoc.org/github.com/jessevdk/go-flags
