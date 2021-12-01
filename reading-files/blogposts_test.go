package blogposts_test

import (
	"fmt"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/quii/learn-go-with-tests/reading-files"
)

type Post struct {
	Title       string
	Description string
	Tag         string
}

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tag: tdd, go`
		secondBody = `Title: Post 2
Description: Description 2
Tag: blah, blah`
	)
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tag:         []string{"tdd", "go"},
	})
	fmt.Printf(" this is what I am looking for %v\n", posts[0])
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
