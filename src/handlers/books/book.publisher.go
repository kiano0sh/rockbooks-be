package books

import (
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

// Publisher

func (publisher *Publisher) CreatePublisher() (*Publisher, error) {
	result := database.DB.Create(&publisher)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("این ناشر قبلا ثبت شده است", result.Error)
	}
	return publisher, nil
}

func (book *Publisher) GetPublisher() (*Publisher, error) {
	var publisherResult *Publisher
	result := database.DB.Where("id = ?", book.ID).First(&publisherResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در دریافت ناشر پیش آمده است!", result.Error)
	}
	// publisherResult.auth
	return publisherResult, nil
}
