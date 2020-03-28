package model

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	if len(alist) != len(articleList) {
		t.Fail()
	}

	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.ID != articleList[i].ID ||
			v.Title != articleList[i].Title {
			t.Fail()
			break
		}
	}
}

func TestGetArticle(t *testing.T) {
	article, err := GetArticle(1)

	if err != nil {
		t.Errorf("Get known article throws error")
	} else if article == nil {
		t.Errorf("No article returned")
	} else if article.ID != 1 {
		t.Errorf("GetArticle(1) = %d; Want 1", article.ID)
	}
}
