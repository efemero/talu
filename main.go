package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	books, err := parse_books("input/books.txt")
	if err != nil {
		panic(err)
	}
	for _, book := range books {
		fmt.Println(book)
	}
}

type Book struct {
	Authors  []string
	Title    string
	Isbn13   string
	Note     int
	Datetime time.Time
}

func parse_books(filename string) (books []Book, e error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	books = []Book{}
	var current_month time.Time
	scanner := bufio.NewScanner(f) // f is the *os.File
	i := 0
	for scanner.Scan() {
		i++
		text := scanner.Text()
		if len(strings.TrimSpace(text)) == 0 {
			// we don't care about empty lines
			continue
		} else if strings.Contains(text, " : ") {
			authors_rest := strings.Split(text, " : ")
			authors_str, rest := authors_rest[0], authors_rest[1]
			title := rest
			note := 0
			switch {
			case strings.HasSuffix(title, "(+)"):
				note = 1
				title = strings.TrimSuffix(title, "(+)")
			case strings.HasSuffix(title, "(++)"):
				note = 2
				title = strings.TrimSuffix(title, "(++)")
			case strings.HasSuffix(title, "(+++)"):
				note = 3
				title = strings.TrimSuffix(title, "(+++)")
			case strings.HasSuffix(title, "(💙)"):
				note = 4
				title = strings.TrimSuffix(title, "(💙)")
			case strings.HasSuffix(title, "(❤️)"):
				note = 5
				title = strings.TrimSuffix(title, "(❤️)")
			}
			authors := strings.Split(authors_str, " & ")
			// this is a book definition
			books = append(books, Book{
				Authors:  authors,
				Note:     note,
				Datetime: current_month,
			})
		} else {
			// this is a month definition
			month_year := strings.Split(text, " ")
			month_fr, year_str := month_year[0], month_year[1]
			var month time.Month
			switch month_fr {
			case "Janvier":
				month = 1
			case "Février":
				month = 2
			case "Mars":
				month = 3
			case "Avril":
				month = 4
			case "Mai":
				month = 5
			case "Juin":
				month = 6
			case "Juillet":
				month = 7
			case "Août":
				month = 8
			case "Septembre":
				month = 9
			case "Octobre":
				month = 10
			case "Novembre":
				month = 11
			case "Décembre":
				month = 12
			default:
				return nil, errors.New(fmt.Sprintf("month '%s' is not valid at line %d", month_fr, i))
			}
			year, err := strconv.Atoi(year_str)
			if err != nil {
				return nil, fmt.Errorf("cannot read year at line %d: %w", i, err)
			}
			current_month = time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)
		}
	}
	if err := scanner.Err(); err != nil {
		// handle error
	}
	return books, nil
}
