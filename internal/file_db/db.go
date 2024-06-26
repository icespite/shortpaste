package file_db

import (
	"github.com/timeforaninja/shortpaste/internal/types"
	"io"
	"net/http"
	"os"
	"path"
)

// LocalFileDB struct containing the storage path.
type localFileDB struct {
	storagePath string
}

func NewLocalFileDB(storagePath string) types.FileDB {
	return &localFileDB{
		storagePath: storagePath,
	}
}

func (db *localFileDB) Write(subPath, content string) error {
	filePath := path.Join(db.storagePath, "data", subPath)

	// make sure path exists
	err := os.MkdirAll(path.Dir(filePath), 0700)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, []byte(content), 0600)
	if err != nil {
		return err
	}

	return nil
}

func (db *localFileDB) WriteStream(subPath string, content io.Reader) error {
	filePath := path.Join(db.storagePath, "data", subPath)

	// make sure path exists
	err := os.MkdirAll(path.Dir(filePath), 0700)
	if err != nil {
		return err
	}

	dst, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, content)
	if err != nil {
		return err
	}

	return nil
}

func (db *localFileDB) ServeFile(w http.ResponseWriter, r *http.Request, subPath string) {
	filePath := path.Join(db.storagePath, "data", subPath)
	http.ServeFile(w, r, filePath)
}

func (db *localFileDB) Read(subPath string) ([]byte, error) {
	// "texts" , fileID+"."+fileExt
	filePath := path.Join(db.storagePath, "data", subPath)
	return os.ReadFile(filePath)
}

func (db *localFileDB) Stat(subPath string) (os.FileInfo, error) {
	filePath := path.Join(db.storagePath, "data", subPath)
	return os.Stat(filePath)
}
