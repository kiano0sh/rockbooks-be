package books

import (
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

// Author

func (author *Author) CreateAuthor() (*Author, error) {
	result := database.DB.Create(&author)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این نویسنده قبلا ثبت شده است", result.Error)
	}
	return author, nil
}

func (author *Author) GetAuthor() (*Author, error) {
	var authorResult *Author
	result := database.DB.Where("id = ?", author.ID).First(&authorResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت نویسنده پیش آمده است!", result.Error)
	}
	// bookResult.auth
	return authorResult, nil
}
