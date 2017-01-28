package main

import (
	"encoding/json"
	// "fmt"
	"github.com/russross/blackfriday"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Replace with your path
var path_prefix = "/home/ubuntu/go/src/github.com/pocolip/blog/"

// Main Page
func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	filenames, _ := ioutil.ReadDir(path_prefix + "blogs")
	blogList := make([]*Blog, 0, 100)

	for left, right := 0, len(filenames)-1; left < right; left, right = left+1, right-1 {
		filenames[left], filenames[right] = filenames[right], filenames[left]
	}

	for _, value := range filenames {
		singleBlog, _ := loadBlog(path_prefix + "blogs/" + value.Name())
		blogList = append(blogList, singleBlog)
	}
	t, _ := template.ParseFiles(path_prefix + "templates/index.html")
	t.Execute(w, blogList)
}

// Display selected blog
func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	title := path_prefix + "blogs/" + r.URL.Path[len("/blog/"):len(r.URL.Path)-1] + ".json"
	b, err := loadBlog(title)
	if err != nil {
		b = &Blog{Title: title}
	}
	b.Content = loadContent(path_prefix + "content/" + strconv.Itoa(b.Id) + ".md")
	t, _ := template.ParseFiles(path_prefix + "templates/blog.html")
	t.Execute(w, b)
}

// Create a blog (PUT)
func createHandler(w http.ResponseWriter, r *http.Request) {
	var blog Blog
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &blog); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//b := RepoCreateBlog(blog)
	// w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// w.WriteHeader(http.StatusCreated)
	// if err := json.NewEncoder(w).Encode(b); err != nil {
	// 	panic(err)
	// }
}

//=============================================================================
// Helper Functions
//=============================================================================

// This is a method named loadBlog that takes a title as its param and returns
// a Blog pointer and an error
func loadBlog(title string) (*Blog, error) {
	body, err := ioutil.ReadFile(title)
	if err != nil {
		return nil, err
	}
	result := Blog{}
	json.Unmarshal(body, &result)
	return &result, nil
}

func loadContent(content string) template.HTML {
	//Load the content from the content folder
	body, err := ioutil.ReadFile(content)
	if err != nil {
		return "No content found"
	}
	return template.HTML(string(blackfriday.MarkdownCommon(body)))
}
