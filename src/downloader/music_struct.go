package downloader

// DJradio 主播电台类型
type DJradio struct {
	Code     int           `json:"code"`
	Count    int           `json:"count"`
	More     bool          `json:"more"`
	Programs []ProgramInfo `json:"programs"`
}

// Song 歌曲类型
type Song struct {
	Code       int           `json:"code"`
	Equalizers []interface{} `json:"equalizers"`
	Songs      []SongInfo    `json:"songs"`
}

// Program 单条主播类型
type Program struct {
	Code    int         `json:"code"`
	Program ProgramInfo `json:"program"`
}

// Artist 歌手类型
type Artist struct {
	Code     int        `json:"code"`
	Artist   ArtistInfo `json:"artist"`
	HotSongs []SongInfo `json:"hotSongs"`
	More     bool       `json:"more"`
}

// ArtistAlbum 歌手专辑类型
type ArtistAlbum struct {
	Code      int         `json:"code"`
	Artist    ArtistInfo  `json:"artist"`
	HotAlbums []AlbumInfo `json:"hotAlbums"`
	More      bool        `json:"more"`
}

// Album 专辑类型
type Album struct {
	Code  int       `json:"code"`
	Album AlbumInfo `json:"album"`
}

// Playlist 歌单类型
type Playlist struct {
	Code   int        `json:"code"`
	Result Collection `json:"result"`
}

//============================

// Collection 类型
type Collection struct {
	AdType                int          `json:"adType"`
	Artists               []ArtistInfo `json:"artists"`
	CloudTrackCount       int          `json:"cloudTrackCount"`
	CommentCount          int          `json:"commentCount"`
	CommentThreadID       string       `json:"commentThreadId"`
	CoverImgID            int          `json:"coverImgId"`
	CoverImgIDStr         string       `json:"coverImgId_str"`
	CoverImgURL           string       `json:"coverImgUrl"`
	CreateTime            int64        `json:"createTime"`
	Creator               UserInfo     `json:"creator"`
	Description           string       `json:"description"`
	HighQuality           bool         `json:"highQuality"`
	ID                    int          `json:"id"`
	Name                  string       `json:"name"`
	NewImported           bool         `json:"newImported"`
	Ordered               bool         `json:"ordered"`
	PlayCount             int          `json:"playCount"`
	Privacy               int          `json:"privacy"`
	ShareCount            int          `json:"shareCount"`
	SpecialType           int          `json:"specialType"`
	Status                int          `json:"status"`
	Subscribed            bool         `json:"subscribed"`
	SubscribedCount       int          `json:"subscribedCount"`
	Subscribers           []string     `json:"subscribers"`
	Tags                  []string     `json:"tags"`
	TotalDuration         int          `json:"totalDuration"`
	TrackCount            int          `json:"trackCount"`
	TrackNumberUpdateTime int64        `json:"trackNumberUpdateTime"`
	TrackUpdateTime       int64        `json:"trackUpdateTime"`
	Tracks                []SongInfo   `json:"tracks"`
	UpdateTime            int64        `json:"updateTime"`
	UserID                int          `json:"userId"`
}

// ProgramInfo 类型
type ProgramInfo struct {
	BdAuditStatus   int        `json:"bdAuditStatus"`
	BlurCoverURL    string     `json:"blurCoverUrl"`
	Buyed           bool       `json:"buyed"`
	Channels        []string   `json:"channels"`
	CommentCount    int        `json:"commentCount"`
	CommentThreadID string     `json:"commentThreadId"`
	CoverURL        string     `json:"coverUrl"`
	CreateTime      int64      `json:"createTime"`
	Description     string     `json:"description"`
	DJ              UserInfo   `json:"dj"`
	Duration        int        `json:"duration"`
	H5Links         string     `json:"h5Links"`
	ID              int        `json:"id"`
	IsPublish       bool       `json:"isPublish"`
	LikedCount      int        `json:"likedCount"`
	ListenerCount   int        `json:"listenerCount"`
	MainSong        SongInfo   `json:"mainSong"`
	MainTrackID     int        `json:"mainTrackId"`
	Name            string     `json:"name"`
	ProgramDesc     string     `json:"programDesc"`
	ProgramFeeType  int        `json:"programFeeType"`
	PubStatus       int        `json:"pubStatus"`
	Radio           RadioInfo  `json:"radio"`
	Reward          bool       `json:"reward"`
	SerialNum       int        `json:"serialNum"`
	ShareCount      int        `json:"shareCount"`
	Songs           []SongInfo `json:"songs"`
	Subscribed      bool       `json:"subscribed"`
	SubscribedCount int        `json:"subscribedCount"`
	TitbitImages    string     `json:"titbitImages"`
	Titbits         string     `json:"titbits"`
	TrackCount      int        `json:"trackCount"`
}

// UserInfo 类型
type UserInfo struct {
	AccountStatus      int      `json:"accountStatus"`
	AuthStatus         int      `json:"authStatus"`
	Authority          int      `json:"authority"`
	AvatarImgID        int      `json:"avatarImgId"`
	AvatarImgIDStr     string   `json:"avatarImgIdStr"`
	AvatarURL          string   `json:"avatarUrl"`
	BackgroundImgID    int      `json:"backgroundImgId"`
	BackgroundImgIDStr string   `json:"backgroundImgIdStr"`
	BackgroundURL      string   `json:"backgroundUrl"`
	Birthday           int64    `json:"birthday"`
	Brand              string   `json:"brand"`
	City               int      `json:"city"`
	DefaultAvatar      bool     `json:"defaultAvatar"`
	Description        string   `json:"description"`
	DetailDescription  string   `json:"detailDescription"`
	DJStatus           int      `json:"djStatus"`
	ExpertTags         []string `json:"expertTags"`
	Followed           bool     `json:"followed"`
	Gender             int      `json:"gender"`
	Mutual             bool     `json:"mutual"`
	Nickname           string   `json:"nickname"`
	Province           int      `json:"province"`
	RemarkName         string   `json:"remarkName"`
	Signature          string   `json:"signature"`
	UserID             int      `json:"userId"`
	UserType           int      `json:"userType"`
	VipType            int      `json:"vipType"`
}

// RadioInfo 类型
type RadioInfo struct {
	Buyed                 bool    `json:"buyed"`
	Category              string  `json:"category"`
	CategoryID            int     `json:"categoryId"`
	CreateTime            int64   `json:"createTime"`
	Desc                  string  `json:"desc"`
	DJ                    string  `json:"dj"`
	Finished              bool    `json:"finished"`
	ID                    int     `json:"id"`
	LastProgramCreateTime int64   `json:"lastProgramCreateTime"`
	LastProgramID         int     `json:"lastProgramId"`
	LastProgramName       string  `json:"lastProgramName"`
	Name                  string  `json:"name"`
	PicURL                string  `json:"picUrl"`
	Price                 float64 `json:"price"`
	ProgramCount          int     `json:"programCount"`
	PurchaseCount         int     `json:"purchaseCount"`
	RadioFeeType          int     `json:"radioFeeType"`
	Rcmdtext              string  `json:"rcmdtext"`
	SubCount              int     `json:"subCount"`
	UnderShelf            bool    `json:"underShelf"`
	Videos                string  `json:"videos"`
}

// ArtistInfo 类型
type ArtistInfo struct {
	AlbumSize   int      `json:"albumSize"`
	Alias       []string `json:"alias"`
	BriefDesc   string   `json:"briefDesc"`
	ID          int      `json:"id"`
	Img1v1ID    int      `json:"img1v1Id"`
	Img1v1IDStr string   `json:"img1v1Id_str"`
	Img1v1URL   string   `json:"img1v1Url"`
	MusicSize   int      `json:"musicSize"`
	Name        string   `json:"name"`
	PicID       int      `json:"picId"`
	PicURL      string   `json:"picUrl"`
	TopicPerson int      `json:"topicPerson"`
	Trans       string   `json:"trans"`
}

// MusicType 类型
type MusicType struct {
	Bitrate     int     `json:"bitrate"`
	DfsID       int64   `json:"dfsId"`
	DfsIDStr    string  `json:"dfsId_str"`
	Extension   string  `json:"extension"`
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	PlayTime    int     `json:"playTime"`
	Size        int64   `json:"size"`
	Sr          int     `json:"sr"`
	VolumeDelta float64 `json:"volumeDelta"`
}

// SongInfo 类型
type SongInfo struct {
	Album           AlbumInfo    `json:"album"`
	Alias           []string     `json:"alias"`
	Artists         []ArtistInfo `json:"artists"`
	Audition        string       `json:"audition"`
	BMusic          MusicType    `json:"bMusic"`
	CommentThreadID string       `json:"commentThreadId"`
	CopyFrom        string       `json:"copyFrom"`
	CopyrightID     int          `json:"copyrightId"`
	Crbt            string       `json:"crbt"`
	DayPlays        int          `json:"dayPlays"`
	Disc            string       `json:"disc"`
	Duration        int          `json:"duration"`
	Fee             int          `json:"fee"`
	Ftype           int          `json:"ftype"`
	HMusic          MusicType    `json:"hMusic"`
	HearTime        int          `json:"hearTime"`
	ID              int          `json:"id"`
	LMusic          MusicType    `json:"lMusic"`
	MMusic          MusicType    `json:"mMusic"`
	Mp3URL          string       `json:"mp3Url"`
	MVID            int          `json:"mvid"`
	Name            string       `json:"name"`
	No              int          `json:"no"`
	PlayedNum       int          `json:"playedNum"`
	Popularity      float64      `json:"popularity"`
	Position        int          `json:"position"`
	Ringtone        string       `json:"ringtone"`
	RtURL           string       `json:"rtUrl"`
	RtURLs          []string     `json:"rtUrls"`
	Rtype           int          `json:"rtype"`
	RURL            string       `json:"rurl"`
	Score           int          `json:"score"`
	Starred         bool         `json:"starred"`
	StarredNum      int          `json:"starredNum"`
	Status          int          `json:"status"`
}

// AlbumInfo 类型
type AlbumInfo struct {
	Alias           []string     `json:"alias"`
	Artist          ArtistInfo   `json:"artist"`
	Artists         []ArtistInfo `json:"artists"`
	BlurPicURL      string       `json:"blurPicUrl"`
	BriefDesc       string       `json:"briefDesc"`
	CommentThreadID string       `json:"commentThreadId"`
	Company         string       `json:"company"`
	CompanyID       int          `json:"companyId"`
	CopyrightID     int          `json:"copyrightId"`
	Description     string       `json:"description"`
	ID              int          `json:"id"`
	Name            string       `json:"name"`
	OnSale          bool         `json:"onSale"`
	Paid            bool         `json:"paid"`
	Pic             int          `json:"pic"`
	PicID           int          `json:"picId"`
	PicIDStr        string       `json:"picId_str"`
	PicURL          string       `json:"picUrl"`
	PublishTime     int64        `json:"publishTime"`
	Size            int          `json:"size"`
	Songs           []SongInfo   `json:"songs"`
	Status          int          `json:"status"`
	SubType         string       `json:"subType"`
	Tags            string       `json:"tags"`
	Atype           string       `json:"type"`
}

// Item 下载文件类型
type Item struct {
	Dir, FileName, FileURL string
}
