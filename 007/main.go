package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/narie-monarie/database"
)

var router *chi.Mux
var db *sql.DB

type Article struct {
	ID      int           `json:"id"`
	Title   string        `json:"title"`
	Content template.HTML `json:"content"`
}

func dbCreateArticle(article *Article) error {
	query, err := db.Prepare("insert into articles(title,content) values (?,?)")
	defer query.Close()

	if err != nil {
		return err
	}

	_, err = query.Exec(article.Title, article.Content)

	if err != nil {
		return err
	}

	return nil

}

func dbGetAllArticles() ([]*Article, error) {
	query, err := db.Prepare("select id, title, content from articles")
	defer query.Close()
	if err != nil {
		return nil, err
	}
	result, err := query.Query()
	if err != nil {
		return nil, err
	}
	articles := make([]*Article, 0)
	for result.Next() {
		data := new(Article)
		err := result.Scan(
			&data.ID,
			&data.Title,
			&data.Content,
		)
		if err != nil {
			return nil, err
		}
		articles = append(articles, data)
	}
	return articles, nil
}

func dbGetArticle(articleID string) (*Article, error) {
	query, err := db.Prepare("select id, title, content from articles where id = ?")
	defer query.Close()
	if err != nil {
		return nil, err
	}
	result := query.QueryRow(articleID)
	data := new(Article)
	err = result.Scan(&data.ID, &data.Title, &data.Content)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func dbUpdateArticle(id string, article *Article) error {
	query, err := db.Prepare("update articles set (title, content) = (?,?) where id=?")
	defer query.Close()
	if err != nil {
		return err
	}
	_, err = query.Exec(article.Title, article.Content, id)
	if err != nil {
		return err
	}
	return nil
}

func dbDeleteArticle(id string) error {
	query, err := db.Prepare("delete from articles where id=?")
	defer query.Close()
	if err != nil {
		return err
	}
	_, err = query.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func catch(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func ChangeMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			switch method := r.PostFormValue("_method"); method {
			case http.MethodPut:
				fallthrough
			case http.MethodPatch:
				fallthrough
			case http.MethodDelete:
				r.Method = method
			default:
			}
		}
		next.ServeHTTP(w, r)
	})
}

func ArticleCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		articleID := chi.URLParam(r, "articleID")
		article, err := dbGetArticle(articleID)
		if err != nil {
			fmt.Println(err)
			http.Error(w, http.StatusText(404), 404)
			return
		}
		ctx := context.WithValue(r.Context(), "article", article)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	router = chi.NewRouter()
	router.Use(middleware.Recoverer)
	var err error
	db, err = database.Connect()
	catch(err)
	router.Use(ChangeMethod)
	router.Get("/", GetAllArticles)
	router.Route("/articles", func(r chi.Router) {
		r.Get("/", NewArticle)
		r.Post("/", CreateArticle)
		r.Route("/{articleID}", func(r chi.Router) {
			r.Use(ArticleCtx)
			r.Get("/", GetArticle)       // GET /articles/1234
			r.Put("/", UpdateArticle)    // PUT /articles/1234
			r.Delete("/", DeleteArticle) // DELETE /articles/1234
			r.Get("/edit", EditArticle)  // GET /articles/1234/edit
		})
	})
	err = http.ListenAndServe(":8005", router)
	catch(err)
}
