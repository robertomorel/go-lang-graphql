package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/robertomorel/go-lang-graphql/graph/generated"
	"github.com/robertomorel/go-lang-graphql/graph/model"
)

// Gerado depois de rodar "gqlgen generate" ------------
func (r *categoryResolver) Courses(ctx context.Context, obj *model.Category) ([]*model.Course, error) {
	//panic(fmt.Errorf("not implemented"))

	// Listar os cursos de uma determinada categoria
	var courses []*model.Course

	// Percorre todos os cursos
	for _, v := range r.Resolver.Courses {
		// Aquele curso que tiver a CategoryId igual ao ID passado...
		if v.Category.ID == obj.ID {
			courses = append(courses, v)
		}
	}

	return courses, nil
}

// Gerado depois de rodar "gqlgen generate" ------------
func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	//panic(fmt.Errorf("not implemented"))

	// Listar os chapters de um determinado curso
	var chapters []*model.Chapter

	// Percorre todos os chapters
	for _, v := range r.Resolver.Chapters {
		// Aquele chapter que tiver a CourseId igual ao ID passado...
		if v.Course.ID == obj.ID {
			chapters = append(chapters, v)
		}
	}

	return chapters, nil
}

/*
	* goes in front of a variable that holds a memory address and resolves it
	(it is therefore the counterpart to the & operator). It goes and gets the
	thing that the pointer was pointing at, e.g
	How to test
	----------------------------
	mutation createCategory {
		createCategory(input: {
			name: "PHP",
			description: "PHP is awesome!"
		}) {
			id
			name
			description
		}
	}
	----------------------------
*/
func (r *mutationResolver) CreateCategory(ctx context.Context, input model.NewCategory) (*model.Category, error) {
	//panic(fmt.Errorf("not implemented"))
	// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
	// The format 'verbs' are derived from C's but are simpler
	// fmt.Sprintf formata I/O
	category := model.Category{
		ID:          fmt.Sprintf("T%d", rand.Int()), // Transforma em dígito o rand.Int()
		Name:        input.Name,
		Description: &input.Description,
	}
	// Append em r.Categories, passando a nova categoria
	// & goes in front of a variable when you want to get that variable's memory address
	r.Categories = append(r.Categories, &category)

	return &category, nil
}

/*
	----------------------------
	mutation createCourse {
		createCourse(input: {
			name: "Evolving with PHP"
			description: "Mega PHP is awesome!"
			categoryId: "T4037200794235010051"
		}) {
			id
			name
			description
			category {
				id
				name
			}
		}
	}
	----------------------------
*/
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	// Criando variável do tipo Category (*model está em memória)
	var category *model.Category

	// Fazendo um laço em todas as r.Categories
	for _, v := range r.Categories {
		// Se o ID da category for igual ao categoryId passado...
		if v.ID == input.CategoryID {
			category = v
		}
	}

	//panic(fmt.Errorf("not implemented"))
	course := model.Course{
		ID:          fmt.Sprintf("T%d", rand.Int()), // Transforma em dígito o rand.Int()
		Name:        input.Name,
		Description: &input.Description,
		Category:    category, // Não podemos informar apenas o ID, mas a entidade inteira com o ID correto
	}

	r.Courses = append(r.Courses, &course)

	return &course, nil
}

/*
	----------------------------
	mutation createChapter {
		createChapter(input: {
			name: "Evolving with PHP - Chapter 1"
			courseId: "T4037200794235010051"
		}) {
			id
			name
			course {
				name
				description
			}
		}
	}
	----------------------------
*/
func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapter) (*model.Chapter, error) {
	// Criando variável do tipo Category (*model está em memória)
	var course *model.Course

	// Fazendo um laço em todas as r.Categories
	for _, v := range r.Courses {
		// Se o ID da category for igual ao categoryId passado...
		if v.ID == input.CourseID {
			course = v
		}
	}

	//panic(fmt.Errorf("not implemented"))
	chapter := model.Chapter{
		ID:     fmt.Sprintf("T%d", rand.Int()), // Transforma em dígito o rand.Int()
		Name:   input.Name,
		Course: course, // Não podemos informar apenas o ID, mas a entidade inteira com o ID correto
	}

	r.Chapters = append(r.Chapters, &chapter)

	return &chapter, nil
}

/*
	How to test
	----------------------------
	query findCategories {
		categories {
			id
			name
			description
			courses {
				name
			}
		}
	}
	----------------------------
*/
func (r *queryResolver) Categories(ctx context.Context) ([]*model.Category, error) {
	//panic(fmt.Errorf("not implemented"))
	// Retornando um a listagem e um erro em branco (sem erro)
	return r.Resolver.Categories, nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	//panic(fmt.Errorf("not implemented"))
	// Retornando um a listagem e um erro em branco (sem erro)
	return r.Resolver.Courses, nil
}

func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	//panic(fmt.Errorf("not implemented"))
	// Retornando um a listagem e um erro em branco (sem erro)
	return r.Resolver.Chapters, nil
}

// Category returns generated.CategoryResolver implementation.
func (r *Resolver) Category() generated.CategoryResolver { return &categoryResolver{r} }

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Gerado depois de rodar "gqlgen generate" ------------
type categoryResolver struct{ *Resolver }
type courseResolver struct{ *Resolver }

// -----------------------------------------------------
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
