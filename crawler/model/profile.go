package model

type User  struct {
	Company               string `json:"company"`
	FollowersCount        int    `json:"followersCount"`
	FolloweesCount        int    `json:"followeesCount"`
	PostedPostsCount      int    `json:"postedPostsCount"`
	PostedEntriesCount    int    `json:"postedEntriesCount"`
	JobTitle              string `json:"jobTitle"`
	TotalCollectionsCount int    `json:"totalCollectionsCount"`
	Username              string `json:"username"`
}
