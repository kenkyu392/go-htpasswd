package htpasswd

import (
	"encoding/csv"
	"io"
)

// A Reader reads records from a htpasswd file.
type Reader struct {
	r *csv.Reader
}

// NewReader returns a new Reader that reads from r.
func NewReader(r io.Reader) *Reader {
	return &Reader{r: newCSVReader(r)}
}

// ReadAll reads all the remaining records from r.
// Each record is a slice of fields.
// A successful call returns err == nil, not err == io.EOF. Because ReadAll is
// defined to read until EOF, it does not treat end of file as an error to be
// reported.
func (r *Reader) ReadAll() ([][]string, error) {
	out := make([][]string, 0)
	records, err := r.r.ReadAll()
	if err != nil {
		return out, err
	}
	for _, parts := range records {
		if len(parts) != 2 {
			continue
		}
		out = append(out, []string{parts[0], parts[1]})
	}
	return out, nil
}

func newCSVReader(r io.Reader) *csv.Reader {
	cr := csv.NewReader(r)
	cr.Comma = ':'
	cr.Comment = '#'
	cr.TrimLeadingSpace = true
	cr.FieldsPerRecord = -1
	return cr
}
