package hw10

import (
	"errors"
	"fmt"
	"io"
	"os"
)

func Copy(fromPath string, toParh string, offset, limit int64) error {

	if offset < 0 {
		return fmt.Errorf("offset is less than zero")
	}

	if limit < 0 {
		return fmt.Errorf("limit is less than zero")
	}

	fromInfo, err := os.Stat(fromPath)
	if err != nil {
		return fmt.Errorf("from file %s : %w", fromPath, errors.Unwrap(err))
	}

	if !fromInfo.Mode().IsRegular() {
		return fmt.Errorf("from file '%s': unnsupported file", fromPath)
	}

	fromfile, err := os.Open(fromPath)
	if err != nil {
		return fmt.Errorf("open file frompath is error")
	}
	defer fromfile.Close()
	fromSize := fromInfo.Size()
	if offset > fromSize {
		return fmt.Errorf(("offset is more than size of file"))
	} else {
		_, err := fromfile.Seek(offset, 0)
		if err != nil {
			return fmt.Errorf("error seek")
		}

	}

	newfile, err := os.Create(toParh)
	if err != nil {
		return fmt.Errorf("error by create to path")
	}
	defer newfile.Close()

	_, err = io.CopyN(newfile, fromfile, limit)
	if err != nil {
		return fmt.Errorf("error copy")
	}
	return nil
}
