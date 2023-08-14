package file_db

import (
	"io"
	"net/http"
	"os"
	"path"
)

// FileDB struct containing the storage path.
type FileDB struct {
	storagePath string
}

func NewFileDB(storagePath string) *FileDB {
	return &FileDB{
		storagePath: storagePath,
	}
}

func (db *FileDB) Write(subPath, content string) error {
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

func (db *FileDB) WriteStream(subPath string, content io.Reader) error {
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

func (db *FileDB) ServeFile(w http.ResponseWriter, r *http.Request, subPath string) {
	filePath := path.Join(db.storagePath, "data", subPath)
	http.ServeFile(w, r, filePath)
}

func (db *FileDB) Read(subPath string) ([]byte, error) {
	// "texts" , fileID+"."+fileExt
	filePath := path.Join(db.storagePath, "data", subPath)
	return os.ReadFile(filePath)
}

func (db *FileDB) Stat(subPath string) (os.FileInfo, error) {
	filePath := path.Join(db.storagePath, "data", subPath)
	return os.Stat(filePath)
}
