### Go : Linked Lists

An implementation of singly Linked List in Go.

Consider a feed (as in twitter) of posts, where each post consists of some data and time and a pointer to the next post.

So our `Post` would look like :

```go

type Post struct {
	body         string
	publishDate int64 // Unix timestamp
	next *Post // link to the next Post
}

```

And, so our `Feed` would look like :

```go

type Feed struct {
	length int // we'll use it later
	start  *Post
	end  *Post
}
```

[Inspired from Ilija Eftimov](https://ieftimov.com/golang-datastructures-linked-lists). You can get here a more detailed and a beautiful explaination of this.
