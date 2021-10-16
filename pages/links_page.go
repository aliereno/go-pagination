package pages

import (
	"net/url"
	"strconv"
)

// Page is used as response struct.
type LinksPage struct {
	NextPage     *string     `json:"nextPage"`
	PreviousPage *string     `json:"previousPage"`
	TotalPage    int         `json:"totalPage"`
	Items        interface{} `json:"items"`
}

func (l LinksPage) Response(page int, pageSize int, totalPage int, items interface{}, query string) interface{} {
	var nextPage, previousPage *string
	if page+1 <= totalPage {
		params := url.Values{}
		params.Add("page", strconv.Itoa(page+1))
		params.Add("page_size", strconv.Itoa(pageSize))
		link := query + "?" + params.Encode()
		nextPage = &link
	}
	if page-1 > 0 {
		params := url.Values{}
		params.Add("page", strconv.Itoa(page-1))
		params.Add("page_size", strconv.Itoa(pageSize))
		link := query + "?" + params.Encode()
		previousPage = &link
	}
	return LinksPage{
		NextPage:     nextPage,
		PreviousPage: previousPage,
		TotalPage:    totalPage,
		Items:        items,
	}
}
