package blogposts_test

import (
	"fmt"
	blogposts "github.com/quii/learn-go-with-tests/reading-files"
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tag: tdd, go
---
Power level over 9000`
		secondBody = `Title: Post 2
Description: Description 2
Tag: blah, blah
---
Gohan Is Stronger`
	)

	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(fs)

	New(t, err)

	assertPostsLength(t, posts, fs)

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tag:         []string{"tdd", "go"},
		Body:        `Power level over 9000`,
	})
}

func New(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}

func assertPostsLength(t *testing.T, posts []blogposts.Post, fs fstest.MapFS) {
	t.Helper()
	if len(posts) != len(fs) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
		Tag:         []string{"tdd", "go"},
		Body:        `Power level over 9000`,
	})
	fmt.Printf(" this is what I am looking for %v\n", posts[0])
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
