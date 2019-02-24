package payments

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
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

func (r *Repository) Get(id string) (*Payment, error) {
	exist, err := r.Exists(id)
	if err != nil {
		return nil, errors.Wrapf(err, "cannot get payment id %q", id)
	}
	if !exist {
		return nil, errors.Errorf("cannot get payment id %q", id)
	}

	path := r.filePath(id)
	f, err := os.OpenFile(path, os.O_RDONLY, 0600)
	defer f.Close()
	if err != nil {
		return nil, errors.Wrap(err, "cannot open file")
	}

	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "cannot read file")
	}

	var payment Payment
	err = json.Unmarshal(bytes, &payment)
	if err != nil {
		return nil, errors.Wrap(err, "cannot unmarshal payment")
	}

	return &payment, nil
}

func (r *Repository) Delete(id string) error {
	exist, err := r.Exists(id)
	if err != nil {
		return errors.Wrapf(err, "cannot delete payment id %q", id)
	}
	if !exist {
		return errors.Errorf("cannot delete payment id %q", id)
	}

	path := r.filePath(id)
	err = os.Remove(path)
	if err != nil {
		return errors.Wrap(err, "cannot delete file")
	}

	return nil
}

func (r *Repository) GetAll() ([]*Payment, error) {
	files, err := ioutil.ReadDir(r.rootPath)
	if err != nil {
		return nil, errors.Wrap(err, "cannot list directory")
	}

	result := []*Payment{}
	for _, file := range files {
		payment, err := r.Get(file.Name())
		if err != nil {
			return nil, errors.Wrap(err, "cannot get one payment")
		}
		result = append(result, payment)
	}

	return result, nil
}

func (r *Repository) filePath(id string) string {
	return filepath.Join(r.rootPath, id)
}
