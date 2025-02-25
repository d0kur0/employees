package graph

import (
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/99designs/gqlgen/graphql"
)

func MarshalTimestamp(t time.Time) graphql.Marshaler {
	if t.IsZero() {
		return graphql.Null
	}

	return graphql.WriterFunc(func(w io.Writer) {
		_, err := io.WriteString(w, strconv.Quote(t.Format(time.RFC3339Nano)))
		if err != nil {
			panic("This should never have happened")
		}
	})
}

func UnmarshalTimestamp(v interface{}) (time.Time, error) {
	if tmpStr, ok := v.(string); ok {
		return time.Parse(time.RFC3339Nano, tmpStr)
	}
	return time.Time{}, errors.New("time should be RFC3339Nano formatted string")
}
