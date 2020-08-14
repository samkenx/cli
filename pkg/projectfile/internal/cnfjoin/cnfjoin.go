package cnfjoin

import (
	"bufio"
	"bytes"
	"errors"
)

type Section struct {
	Key  string
	Data []byte
}

type CnfJoin struct {
	Primary  []byte
	Sections []Section
}

func (j *CnfJoin) Data() ([]byte, error) {
	return join(j.Primary, j.Sections)
}

func join(primary []byte, sections []Section) ([]byte, error) {
	sects := makeSectionTrack(sections)

	var out []byte
	buf := bytes.NewBuffer(primary)

	sc := bufio.NewScanner(buf)

	for sc.Scan() {
		bs := sc.Bytes()
		out = append(out, bs...)
		out = append(out, '\n')

		key, err := yamlKey(bs)
		if err != nil {
			return nil, err
		}
		if key == "" {
			continue
		}

		section, ok := sects[key]
		if !ok {
			continue
		}

		out = append(out, section.data...)
		sects[key].done = true
	}

	for sectName, sectItem := range sects {
		if sectItem.done {
			continue
		}

		out = append(out, []byte(sectName)...)
		out = append(out, ':', '\n')
		out = append(out, sectItem.data...)
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return out, nil
}

func yamlKey(line []byte) (string, error) {
	if len(line) == 0 {
		return "", nil
	}

	if line[0] == '?' {
		return "", errors.New("complex keys are not supported")
	}

	split := bytes.SplitN(line, []byte(":"), 2)
	if len(split) < 2 || len(split[1]) != 0 {
		return "", nil
	}

	return string(split[0]), nil
}

type sectionItem struct {
	data []byte
	done bool
}

type sectionTrack map[string]*sectionItem

func makeSectionTrack(sections []Section) sectionTrack {
	st := make(sectionTrack)

	for _, s := range sections {
		st[s.Key] = &sectionItem{
			data: s.Data,
		}
	}

	return st
}
