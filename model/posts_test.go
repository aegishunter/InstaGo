package model

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPostModel_GetShortDescription(t *testing.T) {
	shortDescriptionPost := NewPost(1, 1, 1, "short", time.Now(), time.Now(), "test.jpg")
	longDescriptionPost := NewPost(1, 1, 1, "this is a long one", time.Now(), time.Now(), "test.jpg")
	assert.Equal(t, "short", shortDescriptionPost.GetShortDescription())
	assert.Equal(t, "this is a ...", longDescriptionPost.GetShortDescription())
}
