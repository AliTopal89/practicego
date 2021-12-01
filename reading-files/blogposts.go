package blogposts

import (
	"bufio"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tag         []string
	//remember string is a single string
	// []string is array of strings
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagSeperator         = "Tag: "
)

// Package fs defines basic interfaces to a file system.

// The FS interface is the minimum implementation required of the file system.
// A file system may implement additional interfaces,
func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post
	for _, f := range dir {
		post, err := getPost(fileSystem, f.Name())
		if err != nil {
			return nil, err //todo: needs clarification, should we totally fail if one file fails? or just ignore?
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err
	}
	defer postFile.Close()
	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	title := readLine(titleSeparator)
	description := readLine(descriptionSeparator)
	tag := strings.Split(readLine(tagSeperator), ", ")
	fmt.Println("what is the post", postFile)

	return Post{Title: title, Description: description, Tag: tag}, nil
}
