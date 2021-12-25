package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"gitlab.com/kian00sh/rockbooks-be/graph/generated"
	"gitlab.com/kian00sh/rockbooks-be/graph/model"
	"gitlab.com/kian00sh/rockbooks-be/src/handlers/books"
	"gitlab.com/kian00sh/rockbooks-be/src/handlers/users"
	"gitlab.com/kian00sh/rockbooks-be/src/jwt"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	mainPagination "gitlab.com/kian00sh/rockbooks-be/src/utils/pagination"
)

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (string, error) {
	var user users.User
	user.DisplayName = input.DisplayName
	user.Email = input.Email
	user.Password = input.Password
	err := user.Register()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (string, error) {
	var user users.User
	user.Email = input.Email
	user.Password = input.Password
	err := user.Authenticate()
	if err != nil {
		return "", err
	}
	token, err := jwt.GenerateToken(user.Email)
	if err != nil {
		return "", err
	}
	return token, nil
}

func (r *mutationResolver) RefreshToken(ctx context.Context, input model.RefreshTokenInput) (string, error) {
	email, err := jwt.ParseToken(input.Token)
	if err != nil {
		return "", grapherrors.ReturnGQLError("Token is invalid!", err)
	}
	token, err := jwt.GenerateToken(email)
	if err != nil {
		return "", grapherrors.ReturnGQLError("Something went wrong with token creation!", err)
	}
	return token, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, input model.CreateBookInput) (*model.Book, error) {
	var book books.Book
	book.Name = input.Name
	book.BookFile = input.BookFile
	book.AuthorID = input.AuthorID
	book.PublisherID = input.PublisherID
	createdBook, err := book.CreateBook()
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input model.UpdateBookInput) (*model.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBookAudio(ctx context.Context, input model.CreateBookAudioInput) (*model.BookAudio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBookAudio(ctx context.Context, input model.UpdateBookAudioInput) (*model.BookAudio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBookAudio(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.CreateAuthorInput) (*model.Author, error) {
	var author books.Author
	author.Name = input.Name
	createdAuthor, err := author.CreateAuthor()
	if err != nil {
		return nil, err
	}
	return createdAuthor, nil
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, input model.UpdateAuthorInput) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePublisher(ctx context.Context, input model.CreatePublisherInput) (*model.Publisher, error) {
	var publisher books.Publisher
	publisher.Name = input.Name
	createdPublisher, err := publisher.CreatePublisher()
	if err != nil {
		return nil, err
	}
	return createdPublisher, nil
}

func (r *mutationResolver) UpdatePublisher(ctx context.Context, input model.UpdatePublisherInput) (*model.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePublisher(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Self(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Authors(ctx context.Context) ([]*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Author(ctx context.Context, id int64) (*model.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Publishers(ctx context.Context) ([]*model.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pages(ctx context.Context, id int64, pagination *model.PaginationInput) (*model.BookPagesWithPagination, error) {
	var bookPage books.BookPage
	bookPage.PaginationInput = mainPagination.CreatePaginationInput(pagination)
	bookPage.BookID = id
	pages, paginationValues, err := bookPage.GetBookPages()
	if err != nil {
		return nil, err
	}
	return &model.BookPagesWithPagination{Pagination: mainPagination.CreatePaginationResult(paginationValues), BookPages: pages}, nil
}

func (r *queryResolver) Audios(ctx context.Context, id int64) ([]*model.BookAudio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Publisher(ctx context.Context, id int64) (*model.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context, pagination *model.PaginationInput) (*model.BooksWithPagination, error) {
	var booksInstance books.Book
	booksInstance.PaginationInput = mainPagination.CreatePaginationInput(pagination)
	booksResult, paginationValues, err := booksInstance.GetBooks()
	if err != nil {
		return nil, err
	}
	return &model.BooksWithPagination{Pagination: mainPagination.CreatePaginationResult(paginationValues), Books: booksResult}, nil
}

func (r *queryResolver) Book(ctx context.Context, id int64) (*model.Book, error) {
	var bookInstance books.Book
	bookInstance.ID = id
	bookResult, err := bookInstance.GetBook()
	if err != nil {
		return nil, err
	}
	return bookResult, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
