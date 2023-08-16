package types

import (
	"io"
	"net/http"
	"os"
)

type FileDB interface {
	Write(subPath, content string) error
	WriteStream(subPath string, content io.Reader) error
	ServeFile(w http.ResponseWriter, r *http.Request, subPath string)
	Read(subPath string) ([]byte, error)
	Stat(subPath string) (os.FileInfo, error)
}
