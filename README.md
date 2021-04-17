Doc: https://gqlgen.com/getting-started/


1. go get github.com/99designs/gqlgen
   go install ....
3. Cria o schema no arquivo "graph/schema.graphqls"   
4. Roda comando: go run github.com/99designs/gqlgen init
5. Separa os models em cada entidade
    - model/category.go
    - model/chapter.go
    - model/course.go
6. Altera arquivo gqlgen.yml para explicitar cada modelo
models:
  Category:
    model: github.com/robertomorel/go-lang-graphql/graph/model/Category 
  Course:
    model: github.com/robertomorel/go-lang-graphql/graph/model/Course 
  Chapter:
    model: github.com/robertomorel/go-lang-graphql/graph/model/Chapter
7. Roda comando: go run github.com/99designs/gqlgen generate        
