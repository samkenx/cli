package cnfjoin

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Joined(dir, prefix, ext string) ([]byte, error) {
	primaryFileName := fmt.Sprintf("%s.%s", prefix, ext)
	primaryFilePath := filepath.Join(dir, primaryFileName)
	primaryData, err := ioutil.ReadFile(primaryFilePath)
	if err != nil {
		return nil, err
	}

	addonGlobName := fmt.Sprintf("%s.*.%s", prefix, ext)
	addonGlobPath := filepath.Join(dir, addonGlobName)
	addonFileNames, err := filepath.Glob(addonGlobPath)
	if err != nil {
		return nil, err
	}

	var sections []Section

	for _, aoFilePath := range addonFileNames {
		aoFileName := filepath.Base(aoFilePath)
		aoData, err := ioutil.ReadFile(aoFilePath)
		if err != nil {
			if errors.Is(os.ErrNotExist, err) {
				continue
			}

			return nil, err
		}

		section := Section{
			Key:  strings.SplitN(aoFileName, ".", 3)[1],
			Data: aoData,
		}

		sections = append(sections, section)
	}

	join := CnfJoin{
		Primary:  primaryData,
		Sections: sections,
	}

	joinData, err := join.Data()
	if err != nil {
		return nil, err
	}

	return joinData, nil
}
