package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

func LoadGrades(path string) ([]GradeRow, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open grades %s: %w", path, err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1

	if _, err := r.Read(); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}

	var rows []GradeRow
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("read row: %w", err)
		}
		grade, err := strconv.Atoi(rec[0])
		if err != nil {
			return nil, fmt.Errorf("parse grade %q: %w", rec[0], err)
		}
		amount, err := strconv.Atoi(rec[2])
		if err != nil {
			return nil, fmt.Errorf("parse amount %q: %w", rec[2], err)
		}
		row := GradeRow{
			Grade:      grade,
			KenpoGrade: optInt(rec[1]),
			Amount:     amount,
			Lower:      optInt(rec[3]),
			Upper:      optInt(rec[4]),
		}
		rows = append(rows, row)
	}
	return rows, nil
}

func optInt(s string) *int {
	if s == "" {
		return nil
	}
	v, err := strconv.Atoi(s)
	if err != nil {
		return nil
	}
	return &v
}
