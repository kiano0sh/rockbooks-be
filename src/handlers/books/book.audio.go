package books

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	gonanoid "github.com/matoous/go-nanoid/v2"
	database "gitlab.com/kian00sh/rockbooks-be/src/database/postgresql"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/consts"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
	"gitlab.com/kian00sh/rockbooks-be/src/utils/strings"
)

// Books

func (bookAudio *BookAudio) CreateBookAudio() (*BookAudio, error) {

	mainFilePath := consts.AUDIOS_PATH + fmt.Sprint(bookAudio.BookID)
	fileID, err := gonanoid.New()
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
	err = os.MkdirAll(mainFilePath, os.ModePerm)
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
	// Handle Book Cover
	audioPath, err := strings.ConcatExtensionToEnd(filepath.Join(mainFilePath, fileID), bookAudio.BookAudioFile.ContentType)
	if err != nil {
		return nil, grapherrors.ReturnGQLError("مشکلی در ثبت صوت کتاب پیش آمده است", err)
	}
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
