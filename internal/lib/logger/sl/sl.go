package sl

import (
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
)

func Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
