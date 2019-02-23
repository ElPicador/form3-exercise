package payments

import (
	"encoding/json"
	"github.com/pkg/errors"
	"os"
	"path/filepath"
)

type Repository struct {
	rootPath string
}

func NewRepository(path string) *Repository {
	return &Repository{
		rootPath: path,
	}
}

func (r *Repository) Exists(id string) (bool, error) {
	path := r.filePath(id)
	_, err := os.Stat(path)

	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return true, err
	}

	return true, nil
}

func (r *Repository) Save(id string, payment *Payment) error {
	path := r.filePath(id)
	bytes, err := json.Marshal(payment)
	if err != nil {
		return errors.Wrap(err, "cannot json serialize")
	}

	// open file in read-write and creates it if needeed
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0600)
	defer f.Close()
	if err != nil {
		return errors.Wrap(err, "cannot open file")
	}

	_, err = f.Write(bytes)
	if err != nil {
		return errors.Wrap(err, "cannot write data to file")
	}

	// force fsync to disk
	err = f.Sync()
	if err != nil {
		return errors.Wrap(err, "cannot fsync the file")
	}

	return nil
}

func (r *Repository) filePath(id string) string {
	return filepath.Join(r.rootPath, id)
}
