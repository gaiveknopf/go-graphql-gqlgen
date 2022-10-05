package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"gaiveknopf/graphql/graph/model"
	"math/rand"
)

// CreateChapter is the resolver for the createChapter field.
func (r *mutationResolver) CreateChapter(ctx context.Context, input model.NewChapterInput) (*model.Chapter, error) {
	var course *model.Course

	for _, v := range r.Courses {
		if v.ID == input.CourseID {
			course = v
		}
	}

	chapter := &model.Chapter{
		ID:     fmt.Sprintf("T%d", rand.Int()),
		Name:   input.Name,
		Course: course,
	}
	r.Chapters = append(r.Chapters, chapter)

	return chapter, nil
}

// Chapters is the resolver for the chapters field.
func (r *queryResolver) Chapters(ctx context.Context) ([]*model.Chapter, error) {
	return r.Resolver.Chapters, nil
}
