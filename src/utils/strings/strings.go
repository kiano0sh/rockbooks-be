package strings

import (
	"mime"
	"strings"

	"gitlab.com/kian00sh/rockbooks-be/src/utils/grapherrors"
)

func NormalizeMediaName(name string) string {
	result := strings.ReplaceAll(name, " ", "-")
	return strings.ToLower(result)
}

func ConcatExtensionToEnd(name string, contentType string) (string, error) {
	coverExtension, err := mime.ExtensionsByType(contentType)
	if err != nil {
		return "", grapherrors.ReturnGQLError("فرمت فایل مشکل دارد", err)
	}
	return name + coverExtension[len(coverExtension)-1], nil
}
