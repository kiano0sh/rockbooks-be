package books

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	gonanoid "github.com/matoous/go-nanoid/v2"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/handlers/users"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/consts"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

// Books

func (bookAudio *BookAudio) CreateBookAudio() (*BookAudio, error) {

	mainFilePath := consts.AUDIOS_PATH + fmt.Sprint(bookAudio.BookPageID)
	fileID, err := gonanoid.New()
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
	err = os.MkdirAll(mainFilePath, os.ModePerm)
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
	// Handle Book Cover
	audioPath := filepath.Join(mainFilePath, fileID) + ".wav"
	touchedAudioFile, err := os.OpenFile(audioPath, consts.CREATE_FILE_FLAG, consts.CREATE_FILE_PERMISSION)
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
	defer touchedAudioFile.Close()
	bytes, err := io.Copy(touchedAudioFile, bookAudio.BookAudioFile.File)
	_ = bytes
	// If error is not nil then panics
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
	// Store where we have placed the audio file
	bookAudio.Audio = audioPath
	// Create the bookAudio
	result := database.DB.Create(&bookAudio)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", result.Error)
	}
	return bookAudio, nil
}

func (bookAudio *BookAudio) GetBookAudioUser() (*users.User, error) {
	var bookAudioUserResult *users.User
	result := database.DB.Where("id = ?", bookAudio.UserID).Find(&bookAudioUserResult)
	if result.Error != nil {
		return nil, grapherrors.ReturnGQLError("در دریافت سازنده صوت این صفحه مشکلی پیش آمده!", result.Error)
	}
	return bookAudioUserResult, nil
}
