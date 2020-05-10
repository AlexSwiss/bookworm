package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AlexSwiss/bookworm/graph/generated"
	"github.com/AlexSwiss/bookworm/graph/model"
	"github.com/AlexSwiss/bookworm/graph/models"
)

func (r *mutationResolver) AddBook(ctx context.Context, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	//create book using input struct
	book := models.Book{
		Name:     input.Name,
		Category: input.Category,
	}

	book.Author = make([]*models.Author, len(author))

	for index, item := range author {
		book.Author[index] = &models.Author{Firstname: item.Firstname, Lastname: item.Lastname}
	}

	db.Create(&book)
	return &book, nil
}

func (r *mutationResolver) EditBook(ctx context.Context, id *int, input *model.NewBook, author []*model.NewAuthor) (*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	var book models.Book

	//find book based on ID
	db = db.Preload("Authors").Where("id = ?", *id).First(&book).Update("name", input.Name)
	if input.Category != "" {
		db.Update("category", *&input.Category)
	}

	//update author
	book.Author = make([]*models.Author, len(author))
	for index, item := range author {
		book.Author[index] = &models.Author{Firstname: item.Firstname, Lastname: item.Lastname}
	}

	db.Save(&book)
	return &book, nil

}

func (r *mutationResolver) DeleteBook(ctx context.Context, id *int) ([]*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	var book models.Book

	//fetch based on ID and delete
	db.Where("id = ?", *id).First(&book).Delete(&book)

	//preload and fetch all recipe
	var books []*models.Book
	db.Preload("Author").Find(&books)

	return books, nil

}

func (r *queryResolver) Books(ctx context.Context, search *string) ([]*models.Book, error) {
	db := models.FetchConnection()
	defer db.Close()

	var books []*models.Book

	//preload loads the author relationship into each book
	db.Preload("Author").Find(&books)

	return books, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
