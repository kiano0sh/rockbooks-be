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

func (r *bookResolver) Author(ctx context.Context, obj *books.Book) (*books.Author, error) {
	authorResult, err := obj.GetBookAuthor()
	if err != nil {
		return nil, err
	}
	return authorResult, nil
}

func (r *bookResolver) Publisher(ctx context.Context, obj *books.Book) (*books.Publisher, error) {
	publisherResult, err := obj.GetBookPublisher()
	if err != nil {
		return nil, err
	}
	return publisherResult, nil
}

func (r *bookResolver) CreatedAt(ctx context.Context, obj *books.Book) (string, error) {
	return obj.CreatedAt.String(), nil
}

func (r *bookAudioResolver) CreatedBy(ctx context.Context, obj *books.BookAudio) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookAudioResolver) Book(ctx context.Context, obj *books.BookAudio) (*books.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *bookAudioResolver) CreatedAt(ctx context.Context, obj *books.BookAudio) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

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

func (r *mutationResolver) CreateBook(ctx context.Context, input model.CreateBookInput) (*books.Book, error) {
	var book books.Book
	book.Name = input.Name
	book.BookFile = input.BookFile
	book.CoverFile = input.CoverFile
	book.WallpaperFile = input.WallpaperFile
	book.AuthorID = input.AuthorID
	book.PublisherID = input.PublisherID
	createdBook, err := book.CreateBook()
	if err != nil {
		return nil, err
	}
	return createdBook, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input model.UpdateBookInput) (*books.Book, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBook(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBookAudio(ctx context.Context, input model.CreateBookAudioInput) (*books.BookAudio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateBookAudio(ctx context.Context, input model.UpdateBookAudioInput) (*books.BookAudio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteBookAudio(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateAuthor(ctx context.Context, input model.CreateAuthorInput) (*books.Author, error) {
	var author books.Author
	author.Name = input.Name
	createdAuthor, err := author.CreateAuthor()
	if err != nil {
		return nil, err
	}
	return createdAuthor, nil
}

func (r *mutationResolver) UpdateAuthor(ctx context.Context, input model.UpdateAuthorInput) (*books.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteAuthor(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePublisher(ctx context.Context, input model.CreatePublisherInput) (*books.Publisher, error) {
	var publisher books.Publisher
	publisher.Name = input.Name
	createdPublisher, err := publisher.CreatePublisher()
	if err != nil {
		return nil, err
	}
	return createdPublisher, nil
}

func (r *mutationResolver) UpdatePublisher(ctx context.Context, input model.UpdatePublisherInput) (*books.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePublisher(ctx context.Context, id int64) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Self(ctx context.Context) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Authors(ctx context.Context) ([]*books.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Author(ctx context.Context, id int64) (*books.Author, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Publishers(ctx context.Context) ([]*books.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Pages(ctx context.Context, id int64, pagination *model.PaginationInput) (*model.BookPagesWithPagination, error) {
	var bookPage books.BookPage
	bookPage.PaginationInput = mainPagination.CreatePaginationInput(&mainPagination.PaginationOutput{Limit: *pagination.Limit, Page: *pagination.Page, SortBy: pagination.SortBy.String(), SortOrder: pagination.SortOrder.String()})
	bookPage.BookID = id
	pagesResult, paginationValues, err := bookPage.GetBookPages()
	if err != nil {
		return nil, err
	}
	createdPaginationOutput := mainPagination.CreatePaginationResult(paginationValues)
	graphPaginationOutput := &model.PaginationType{Limit: createdPaginationOutput.Limit, Page: createdPaginationOutput.Page, Total: createdPaginationOutput.Total}

	return &model.BookPagesWithPagination{Pagination: graphPaginationOutput, BookPages: pagesResult}, nil
}

func (r *queryResolver) Audios(ctx context.Context, id int64) ([]*books.BookAudio, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Publisher(ctx context.Context, id int64) (*books.Publisher, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Books(ctx context.Context, pagination *model.PaginationInput) (*model.BooksWithPagination, error) {
	var booksInstance books.Book
	booksInstance.PaginationInput = mainPagination.CreatePaginationInput(&mainPagination.PaginationOutput{Limit: *pagination.Limit, Page: *pagination.Page, SortBy: pagination.SortBy.String(), SortOrder: pagination.SortOrder.String()})
	booksResult, paginationValues, err := booksInstance.GetBooks()
	if err != nil {
		return nil, err
	}
	createdPaginationOutput := mainPagination.CreatePaginationResult(paginationValues)
	graphPaginationOutput := &model.PaginationType{Limit: createdPaginationOutput.Limit, Page: createdPaginationOutput.Page, Total: createdPaginationOutput.Total}

	return &model.BooksWithPagination{Pagination: graphPaginationOutput, Books: booksResult}, nil
}

func (r *queryResolver) Book(ctx context.Context, id int64) (*books.Book, error) {
	var bookInstance books.Book
	bookInstance.ID = id
	bookResult, err := bookInstance.GetBook()
	if err != nil {
		return nil, err
	}
	return bookResult, nil
}

// Book returns generated.BookResolver implementation.
func (r *Resolver) Book() generated.BookResolver { return &bookResolver{r} }

// BookAudio returns generated.BookAudioResolver implementation.
func (r *Resolver) BookAudio() generated.BookAudioResolver { return &bookAudioResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type bookResolver struct{ *Resolver }
type bookAudioResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
