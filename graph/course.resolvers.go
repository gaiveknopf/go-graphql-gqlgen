package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gaiveknopf/graphql/graph/generated"
	"gaiveknopf/graphql/graph/model"
	"math/rand"
)

// Chapters is the resolver for the chapters field.
func (r *courseResolver) Chapters(ctx context.Context, obj *model.Course) ([]*model.Chapter, error) {
	var chapters []*model.Chapter

	for _, v := range r.Resolver.Chapters {
		if v.Course.ID == obj.ID {
			chapters = append(chapters, v)
		}
	}

	return chapters, nil
}

// CreateCourse is the resolver for the createCourse field.
func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourseInput) (*model.Course, error) {
	fmt.Println("Chegou aqui")

	var category *model.Category

	for _, v := range r.Categories {
		if v.ID == input.CategoryID {
			category = v
		}
	}

	course := &model.Course{
		ID:          fmt.Sprintf("T%d", rand.Int()),
		Name:        input.Name,
		Description: &input.Description,
		Category:    category,
	}

	r.Courses = append(r.Courses, course)
	return course, nil
}

// Courses is the resolver for the courses field.
func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return r.Resolver.Courses, nil
}

// Course returns generated.CourseResolver implementation.
func (r *Resolver) Course() generated.CourseResolver { return &courseResolver{r} }

type courseResolver struct{ *Resolver }
