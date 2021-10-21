package main

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

const (
	FREQUENCY = "%f"
	CALL      = "%c"
	DATE      = "%d"
	TIME      = "%t"
	BAND      = "%b"
	MODE      = "%m"
)

func readInputFiles(files StringArray) (cs []Contact, err error) {
	cs = make([]Contact, 0)
	for _, fileName := range files {
		f, err := os.Open(fileName)
		if err != nil {
			return cs, err
		}
		reader := bufio.NewReader(f)
		setters, err := readStructure(reader)
		if err != nil {
			return nil, err
		}
		contacts, err := readContacts(reader, setters)
		if err != nil {
			return nil, err
		}
		cs = append(cs, contacts...)
	}
	return cs, nil
}

func readStructure(reader *bufio.Reader) ([]FieldSetter, error) {
	line := ""
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		trimmedLine := strings.TrimSpace(string(l))
		if trimmedLine != "" {
			line = strings.TrimSpace(trimmedLine[1:])
			break
		}
	}

	return parseReadingTemplate(line), nil
}

func readContacts(reader *bufio.Reader, setters []FieldSetter) (contacts []Contact, err error) {
	for {
		l, err := reader.ReadString('\n')
		if err == io.EOF {
			return contacts, nil
		}
		if err != nil {
			return nil, err
		}
		trimmedLine := strings.TrimSpace(string(l))
		trimmedLine = stripComment(trimmedLine)
		var fields []string
		if trimmedLine != "" {
			fields = strings.Split(trimmedLine, "\t")
		}
		if len(fields) != len(setters) {
			return nil, errors.New("The number of fields in a line doesn't match the template")
		}
		contact := Contact{}
		for i, f := range fields {
			setters[i](&contact, f)
		}
		contacts = append(contacts, contact)
	}
}

func stripComment(line string) string {
	p := strings.Index(line, "\"")
	if p == -1 {
		return line
	}
	return strings.TrimSpace(line[0:p])
}

type LineReader struct {
	reader *bufio.Reader
}

func NewLineReader(f io.Reader) *LineReader {
	return &LineReader{bufio.NewReader(f)}
}

func (lr *LineReader) readLine() (line string, comment string, err error) {
	l, err := lr.reader.ReadString('\n')
	if err != nil {
		return "", "", err
	}
	p := strings.Index(l, "\"")
	if p == -1 {
		return strings.TrimSpace(line), "", nil
	}
	return strings.TrimSpace(l[0:p]), strings.TrimSpace(l[p+1 : len(l)]), nil
}
