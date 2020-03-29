package model

import "errors"

// Article contains the ID, title and text of an article
type Article struct {
	ID      uint32 `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Article 1 body"},
	Article{ID: 2, Title: "Article 2", Content: "Article 2 body"},
}

// GetAllArticles retrieves all articles
func GetAllArticles() []Article {
	return articleList
}

// GetArticle takes an article id and returns the article or throws an error if article isn't found
func GetArticle(id uint32) (*Article, error) {
	for _, a := range articleList {
		if a.ID == id {
			return &a, nil
		}
	}

	return nil, errors.New("Can't find article")
}
