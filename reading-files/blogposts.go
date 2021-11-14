package blogposts

import "io/fs"

type Post struct {
}

// Package fs defines basic interfaces to a file system.

// The FS interface is the minimum implementation required of the file system.
// A file system may implement additional interfaces,
func NewPostsFromFS(fileSystem fs.FS) []Post {
	dir, _ := fs.ReadDir(fileSystem, ".")
	var posts []Post
	for range dir {
		posts = append(posts, Post{})
	}
	return posts
}
