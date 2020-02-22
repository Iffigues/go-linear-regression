package csvs

import (
	"encoding/csv"
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

// GetFileContent return content file
func GetFileContent(file string) (text string, err error) {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		return "", err
	}
	text = string(content)
	return text, nil
}

func getVal(g []string) (a, b float64, err error) {
	a, err = strconv.ParseFloat(g[0], 64)
	if err != nil {
		return
	}
	b, err = strconv.ParseFloat(g[1], 64)
	return
}

func getIntMap(rec [][]string) (c, g []float64, err error) {
	for _, val := range rec {
		if len(val) > 1 {
			a, b, err := getVal(val)
			if err == nil {
				c = append(c, a)
				g = append(g, b)
			}
		}
	}
	return
}

func isGood(g interface{}) (err error) {
	switch g.(type) {
	case [][]string:
		return nil
	default:
		return errors.New("pls dataset have to be [][]string")
	}
}

// GetCsv return float array in csv
func GetCsv(file string) (a, b []float64, err error) {
	in, err := GetFileContent(file)
	if err != nil {
		return nil, nil, err
	}
	r := csv.NewReader(strings.NewReader(in))
	records, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	if len(records) < 2 {
		return nil, nil, errors.New("empty data")
	}
	a, b, err = getIntMap(records[1:])
	return
}
