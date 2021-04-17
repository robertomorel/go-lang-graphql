/*
	É aqui que tem a implementação para trazer ou inserir os dados que queremos.
	Terá todas as regras de negócio necessárias
*/
package graph

import "github.com/robertomorel/go-lang-graphql/graph/model"

/*
	This file will not be regenerated automatically.
	Este arquivo será usado para injeção de dependências para aplicação.

	Devemos adicionar qualquer dependência necessária aqui
*/

type Resolver struct {
	Categories []*model.Category
	Courses    []*model.Course
	Chapters   []*model.Chapter
}
