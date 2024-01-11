package blog

import fmt "fmt"

const (
	ModuleName = "blog"
	StoreKey   = ModuleName

	PostKey    = "post"
	CommentKey = "comment"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

func KeyCommentPrefix(postSlug string) []byte {
	commentPrefix := fmt.Sprintf("%s_%s", CommentKey, postSlug)

	return KeyPrefix(commentPrefix)
}
