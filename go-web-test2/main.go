package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// BlogPage represents the data to be used in the blog homepage template
type BlogPage struct {
	Title    string
	NavLinks []NavLink
	Posts    []Post
	Sidebar  []SidebarLink
}

// NavLink represents a link in the navigation bar
type NavLink struct {
	Title string
	URL   string
}

// Post represents a blog post
type Post struct {
	Title   string
	Content string
}

// SidebarLink represents a link in the sidebar
type SidebarLink struct {
	Title string
	URL   string
}

func main() {
	http.HandleFunc("/", blogHandler)
	http.ListenAndServe(":8080", nil)
}

func blogHandler(w http.ResponseWriter, r *http.Request) {
	// Define the navigation bar links
	navLinks := []NavLink{
		{"Web建站", "/web"},
		{"VPS服务器推荐", "/vps"},
		{"PHP转Go", "/php"},
	}

	// Define the sidebar links
	sidebarLinks := []SidebarLink{
		{"关于我们", "/about"},
		{"联系我们", "/contact"},
	}

	// Define some sample blog posts
	posts := []Post{
		{"Hello, World!", "Welcome to my blog!"},
		{"My First Post", "This is my first blog post."},
		{"Another Post", "This is another blog post."},
	}

	// Create a new BlogPage object with the relevant data
	pageData := BlogPage{
		Title:    "My Blog",
		NavLinks: navLinks,
		Posts:    posts,
		Sidebar:  sidebarLinks,
	}

	// Parse the blog homepage template
	tmpl, err := template.ParseFiles("blog.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		return
	}

	// Execute the blog homepage template with the BlogPage data
	err = tmpl.Execute(w, pageData)
	if err != nil {
		fmt.Println("Error executing template:", err)
		return
	}
}
