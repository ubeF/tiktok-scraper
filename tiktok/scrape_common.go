package tiktok

type Author struct {
	AvatarThumb string
	Nickname    string
	UniqueId    string
	SecUId      string
	Signature   string
}

type Challenge struct {
	Title string
}

type Video struct {
	Cover string
}

type Stats struct {
	CollectCount int
	CommentCount int
	DiggCount    int
	PlayCount    int
	ShareCount   int
}
