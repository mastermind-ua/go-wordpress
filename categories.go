package wordpress

import (
	"fmt"
	"net/http"
)

type Category struct {
	ID          int    `json:"id,omitempty"`
	Count       int    `json:"count,omitempty"`
	Description string `json:"description,omitempty"`
	Link        string `json:"link,omitempty"`
	Name        string `json:"name,omitempty"`
	Slug        string `json:"slug,omitempty"`
	Taxonomy    string `json:"taxonomy,omitempty"`
	Parent      int    `json:"parent,omitempty"`
}

type CategoriesCollection struct {
	client *Client
	url    string
}

func (col *CategoriesCollection) List(params interface{}) ([]Category, *http.Response, []byte, error) {
	var categories []Category
	resp, body, err := col.client.List(col.url, params, &categories)
	return categories, resp, body, err
}

func (col *CategoriesCollection) Get(id int, params interface{}) (*Category, *http.Response, []byte, error) {
	var entity Category
	entityURL := fmt.Sprintf("%v/%v", col.url, id)
	resp, body, err := col.client.Get(entityURL, params, &entity)
	return &entity, resp, body, err
}
