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

 - Generate the boilerplate
 - Get dependancies
 - Find a library that suits our need

### Part 4

 - Write code, for a complete tutorial see https://goswagger.io/tutorial/todo-list.html

### Usage

Start and build containers using ```docker-compose up --build``` while being in the base directory.
Containers can be interrupted either by typing Ctrl+C or ```docker-compose down```.

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
 - Book: The Go Programming Language (Alan A. A. Donovan (Author), Brian W. Kernighan)
