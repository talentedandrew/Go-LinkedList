package main

import (
	"errors"
	"fmt"
	"time"
)

type Post struct {
	body        string
	publishDate int64 // Unix timestamp
	next        *Post
}

type Feed struct {
	length int // we'll use it later
	start  *Post
	end    *Post
}

func (f *Feed) Append(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		lastPost := f.end
		lastPost.next = newPost
		f.end = newPost
	}
	f.length++
}

func (f *Feed) Prepend(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
		f.end = newPost
	} else {
		firstPost := newPost
		firstPost.next = f.start
		f.start = newPost
	}
	f.length++
}

func (f *Feed) Remove(publishDate int64) {
	if f.length == 0 {
		panic(errors.New("Feed is empty"))
	}

	var previousPost *Post
	currentPost := f.start

	for currentPost.publishDate != publishDate {
		if currentPost.next == nil {
			panic(errors.New("No such Post found."))
		}

		previousPost = currentPost
		currentPost = currentPost.next
	}
	previousPost.next = currentPost.next

	f.length--
}

func (f *Feed) Insert(newPost *Post) {
	if f.length == 0 {
		f.start = newPost
	} else {
		var previousPost *Post
		currentPost := f.start

		for currentPost.publishDate < newPost.publishDate {
			previousPost = currentPost
			currentPost = previousPost.next
		}

		previousPost.next = newPost
		newPost.next = currentPost
	}
	f.length++
}

func (f *Feed) Inspect() {
	if f.length == 0 {
		fmt.Println("Feed is empty")
	}
	fmt.Println("========================")
	fmt.Println("Feed Length: ", f.length)

	currentIndex := 0
	currentPost := f.start

	for currentIndex < f.length {
		fmt.Printf("Item: %v - %v\n", currentIndex, currentPost)
		currentPost = currentPost.next
		currentIndex++
	}
	fmt.Println("========================")
}

func main() {
	rightNow := time.Now().Unix()
	f := &Feed{}
	p1 := &Post{
		body:        "Let's start.",
		publishDate: rightNow,
	}
	p2 := &Post{
		body:        "Oh! I wanna come too.",
		publishDate: rightNow + 10,
	}
	p3 := &Post{
		body:        "I will also join.",
		publishDate: rightNow + 20,
	}
	p4 := &Post{
		body:        "I'm on my way.",
		publishDate: rightNow + 30,
	}
	pN := &Post{
		body:        "I'll board first",
		publishDate: rightNow + 5,
	}
	f.Append(p1)
	f.Append(p2)
	f.Append(p3)
	f.Append(p4)
	f.Prepend(pN)

	f.Inspect()

	newPost := &Post{
		body:        "This is a new post",
		publishDate: rightNow + 15,
	}
	f.Insert(newPost)
	f.Inspect()
}
