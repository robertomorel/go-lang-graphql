# Golang for GraphQL

## Troublesome
Many types of Clients (Desktop, TV, Web, Mobile) consuming the same backend API, which sends back the same information to all these frontend applications. Even though they don't need all available data. 

> This behavior could slow the system. 

That's why BFF (Backend for Frontend) exists. To prepare one backend specifically for a mobile application, another one for Desktop, another for Browser, etc. 
Each of these clients only need a piece of specific information from the entire payload. 

### Solution
If we had a way to give the Client an absolute autonomy to choose which data it would like to receive?

## About 
GraphQL is a query language for APIs and a runtime for fulfilling those queries with your existing data. GraphQL provides a complete and understandable description of the data in your API, gives clients the power to ask for exactly what they need and nothing more, makes it easier to evolve APIs over time, and enables powerful developer tools.

> Created by Facebook

### Advantages
- Unique endpoint
- Unique request
- Server presents all available resources
- Client requests only the needed information
- Data inserts by Mutations
- Work with HTTP
- Return JSON as a response

### Dynamic
It works by a GraphQL Server, which search for entities or another APIs, integrates the data and returns accordingly to the Client's request.

### "N+1" Problematic
When the Server creates many queries and SQL, one for each node that it brings. 
It´s a classic error in GraphQL.

To know more, click [here](https://gqlgen.com/reference/dataloaders/)

### Docs
Click [here](https://gqlgen.com/getting-started/) to getting started 

## How to iniciate a project 
1. Create the project folder
2. Iniciate Go project with `go mod init`
3. Run the following commands:
```shell
go get github.com/99designs/gqlgen
go install github.com/99designs/gqlgen
```
4. Create the schema in the file "graph/schema.graphqls"   
5. Run the command `go run github.com/99designs/gqlgen init`
6. Split the models for each entity
    - model/category.go
    - model/chapter.go
    - model/course.go
7. Update the file "gqlgen.yml" to take down each model. In the end, it should look like this:
```bash
models:
  Category:
    model: github.com/robertomorel/go-lang-graphql/graph/model.Category 
  Course:
    model: github.com/robertomorel/go-lang-graphql/graph/model.Course 
  Chapter:
    model: github.com/robertomorel/go-lang-graphql/graph/model.Chapter
```
8. Run the command `go run github.com/99designs/gqlgen generate`        

## Run the project
You must have Golang environment in your machine, and than run:
```bash
# Clone project
git clone https://github.com/robertomorel/go-lang-graphql.git

# Enter in the project folder
cd ./go-lang-graphql

# Run Server
go run server.go
```

------

## Lets talk
[LinkedIn](https://www.linkedin.com/in/roberto-morel-6b9065193/)
