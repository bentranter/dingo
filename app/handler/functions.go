package handler

import (
	"github.com/dinever/dingo/app/model"
)

func getAllTags() []*model.Tag {
	tags := new(model.Tags)
	_ = tags.GetAllTags()
	return tags.GetAll()
}

func getRecentPosts() []*model.Post {
	posts := new(model.Posts)
	_, _ = posts.GetPostList(1, 5, false, true, "published_at DESC")
	return *posts
}
