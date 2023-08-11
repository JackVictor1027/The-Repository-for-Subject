package model

type LikesModel struct {
	ActionType string `json:"action_type"` // 1-点赞，2-取消点赞
	Token      string `json:"token"`       // 用户鉴权token
	VideoID    string `json:"video_id"`    // 视频id
}
type VideoInfo struct {
	Id     int `json:"id"`
	Author struct {
		Id              int    `json:"id"`
		Name            string `json:"name"`
		FollowCount     int    `json:"follow_count"`
		FollowerCount   int    `json:"follower_count"`
		IsFollow        bool   `json:"is_follow"`
		Avatar          string `json:"avatar"`
		BackgroundImage string `json:"background_image"`
		Signature       string `json:"signature"`
		TotalFavorited  string `json:"total_favorited"`
		WorkCount       int    `json:"work_count"`
		FavoriteCount   int    `json:"favorite_count"`
	} `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int    `json:"favorite_count"`
	CommentCount  int    `json:"comment_count"`
	IsFavorite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}
