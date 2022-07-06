package utils

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"strings"
)

const (
	HeadLine int = 0
	Text     int = 1
)

type Item struct {
	Type  int
	Value string
}

func ConvertHtml(htmlText string) (*[]Item, error) {
	if htmlText == "" {
		return nil, fmt.Errorf("htmlText is required")
	}

	var items []Item
	htmlTokenizer := html.NewTokenizer(strings.NewReader(htmlText))

	for {
		switch htmlTokenizer.Next() {
		case html.ErrorToken:
			if htmlTokenizer.Err() == io.EOF {
				return &items, nil
			}
			return nil, htmlTokenizer.Err()
		case html.TextToken:
			if len(items) > 0 {
				items = append(items, Item{
					Type:  Text,
					Value: htmlTokenizer.Token().Data,
				})
			} else {
				items = append(items, Item{
					Type:  HeadLine,
					Value: htmlTokenizer.Token().Data,
				})
			}
		}
	}
}
