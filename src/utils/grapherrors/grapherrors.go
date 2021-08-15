package grapherrors

import (
	"log"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func ReturnGQLError(Message string, OriginalError error) error {
	log.Println(OriginalError)
	// TODO send logs to kibana
	return gqlerror.Errorf(Message)
}
