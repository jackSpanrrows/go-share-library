package downloader

import (
	"github.com/parnurzeal/gorequest"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
)

// Action for DJradio
func (in *DJradio) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	var item Item
	for i := 0; i < in.Count; i++ {
		all := in.Count
		item.Dir = in.Programs[i].DJ.Brand + "/"
		if temp := in.Programs[i].MainSong.HMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Programs[i].MainSong.Name + "(" + strconv.Itoa(in.Programs[i].MainSong.HMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Programs[i].MainSong.HMusic.DfsID, 10))
		} else if temp := in.Programs[i].MainSong.MMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Programs[i].MainSong.Name + "(" + strconv.Itoa(in.Programs[i].MainSong.MMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Programs[i].MainSong.MMusic.DfsID, 10))
		} else {
			item.FileName = iton(i, all) + "." + in.Programs[i].MainSong.Name + "(" + strconv.Itoa(in.Programs[i].MainSong.LMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Programs[i].MainSong.LMusic.DfsID, 10))
		}
		items = append(items, item)
	}
	return
}

// Action for Song
func (in *Song) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	var item Item
	for i := 0; i < len(in.Songs); i++ {
		all := len(in.Songs)
		item.Dir = in.Songs[i].Album.Name + "/"
		if temp := in.Songs[i].HMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Songs[i].Name + "(" + strconv.Itoa(in.Songs[i].HMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Songs[i].HMusic.DfsID, 10))
		} else if temp := in.Songs[i].MMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Songs[i].Name + "(" + strconv.Itoa(in.Songs[i].MMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Songs[i].MMusic.DfsID, 10))
		} else {
			item.FileName = iton(i, all) + "." + in.Songs[i].Name + "(" + strconv.Itoa(in.Songs[i].LMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Songs[i].LMusic.DfsID, 10))
		}
		items = append(items, item)
	}
	return
}

// Action for Program
func (in *Program) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	var item Item
	item.Dir = in.Program.Name + "/"
	if temp := in.Program.MainSong.HMusic.DfsID; temp != 0 {
		item.FileName = in.Program.MainSong.Name + "(" + strconv.Itoa(in.Program.MainSong.HMusic.Bitrate/1000) + "k).mp3"
		item.FileURL = encryptedID(strconv.FormatInt(in.Program.MainSong.HMusic.DfsID, 10))
	} else if temp := in.Program.MainSong.MMusic.DfsID; temp != 0 {
		item.FileName = in.Program.MainSong.Name + "(" + strconv.Itoa(in.Program.MainSong.MMusic.Bitrate/1000) + "k).mp3"
		item.FileURL = encryptedID(strconv.FormatInt(in.Program.MainSong.MMusic.DfsID, 10))
	} else {
		item.FileName = in.Program.MainSong.Name + "(" + strconv.Itoa(in.Program.MainSong.LMusic.Bitrate/1000) + "k).mp3"
		item.FileURL = encryptedID(strconv.FormatInt(in.Program.MainSong.LMusic.DfsID, 10))
	}
	items = append(items, item)
	return
}

// Action for Artist
func (in *Artist) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	var item Item
	for i := 0; i < len(in.HotSongs); i++ {
		all := len(in.HotSongs)
		item.Dir = in.Artist.Name + "/"
		if temp := in.HotSongs[i].HMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.HotSongs[i].Name + "(" + strconv.Itoa(in.HotSongs[i].HMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.HotSongs[i].HMusic.DfsID, 10))
		} else if temp := in.HotSongs[i].MMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.HotSongs[i].Name + "(" + strconv.Itoa(in.HotSongs[i].MMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.HotSongs[i].MMusic.DfsID, 10))
		} else {
			item.FileName = iton(i, all) + "." + in.HotSongs[i].Name + "(" + strconv.Itoa(in.HotSongs[i].LMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.HotSongs[i].LMusic.DfsID, 10))
		}
		items = append(items, item)
	}
	return
}

// Action for ArtistAlbum
func (in *ArtistAlbum) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	for i := 0; i < len(in.HotAlbums); i++ {
		tempURL := "http://music.163.com/#/album?id=" + strconv.Itoa(in.HotAlbums[i].ID)
		a, _, _ := parseURL(tempURL)
		var in Album
		tempList := in.Action(a)
		for i := 0; i < len(tempList); i++ {
			items = append(items, tempList[i])
		}
	}
	return
}

// Action for Album
func (in *Album) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	var item Item
	for i := 0; i < in.Album.Size; i++ {
		all := in.Album.Size
		item.Dir = in.Album.Name + "/"
		if temp := in.Album.Songs[i].HMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Album.Songs[i].Name + "(" + strconv.Itoa(in.Album.Songs[i].HMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Album.Songs[i].HMusic.DfsID, 10))
		} else if temp := in.Album.Songs[i].MMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Album.Songs[i].Name + "(" + strconv.Itoa(in.Album.Songs[i].MMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Album.Songs[i].MMusic.DfsID, 10))
		} else {
			item.FileName = iton(i, all) + "." + in.Album.Songs[i].Name + "(" + strconv.Itoa(in.Album.Songs[i].LMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Album.Songs[i].LMusic.DfsID, 10))
		}
		items = append(items, item)
	}
	return
}

// Action for Playlist
func (in *Playlist) Action(URL string) (items []Item) {
	gorequest.New().Get(URL).
		Set("cookie", "appver=1.5.2").
		Set("referer", "http://music.163.com/").
		EndStruct(&in)
	var item Item
	for i := 0; i < in.Result.TrackCount; i++ {
		all := in.Result.TrackCount
		item.Dir = in.Result.Name + "/"
		if temp := in.Result.Tracks[i].HMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Result.Tracks[i].Name + "(" + strconv.Itoa(in.Result.Tracks[i].HMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Result.Tracks[i].HMusic.DfsID, 10))
		} else if temp := in.Result.Tracks[i].MMusic.DfsID; temp != 0 {
			item.FileName = iton(i, all) + "." + in.Result.Tracks[i].Name + "(" + strconv.Itoa(in.Result.Tracks[i].MMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Result.Tracks[i].MMusic.DfsID, 10))
		} else {
			item.FileName = iton(i, all) + "." + in.Result.Tracks[i].Name + "(" + strconv.Itoa(in.Result.Tracks[i].LMusic.Bitrate/1000) + "k).mp3"
			item.FileURL = encryptedID(strconv.FormatInt(in.Result.Tracks[i].LMusic.DfsID, 10))
		}
		items = append(items, item)
	}
	return
}

func parseURL(URL string) (sURL string, action string, err error) {
	URL = strings.Replace(URL, "/#", "", -1)
	u, err := url.Parse(URL)
	if err != nil {
		return
	}
	if u.Host != "music.163.com" {
		err = errors.New("URL is not from music.163.com")
		return
	}
	m, err := url.ParseQuery(u.RawQuery)
	if err != nil {
		return
	}
	var prefix, id string
	prefix = "http://music.163.com/api"
	id = m["id"][0]

	switch u.Path {
	case "/album":
		sURL = prefix + "/album/" + id
		action = "Album"
	case "/song":
		sURL = prefix + "/song/detail/?id=" + id + "&ids=%5B" + id + "%5D"
		action = "Song"
	case "/artist":
		sURL = prefix + "/artist/" + id
		action = "Artist"
	case "/artist/album":
		sURL = prefix + "/artist/albums/" + id + "?offset=0&limit=1000"
		action = "ArtistAlbum"
	case "/playlist":
		sURL = prefix + "/playlist/detail?id=" + id + "&ids=%5B" + id + "%5D"
		action = "Playlist"
	case "/program":
		sURL = prefix + "/dj/program/detail?id=" + id + "&ids=%5B" + id + "%5D"
		action = "Program"
	case "/djradio":
		sURL = prefix + "/dj/program/byradio?radioId=" + id + "&asc=1&limit=1000"
		action = "DJradio"
	}
	return
}

func encryptedID(id string) (URLString string) {
	byte1 := []byte(id)
	byte2 := []byte("3go8&$8*3*3h0k(2)2")
	for i, v := range byte1 {
		byte1[i] = v ^ byte2[i%len(byte2)]
	}
	h := md5.New()
	h.Write(byte1)
	encodeString := base64.StdEncoding.EncodeToString(h.Sum(nil))
	encodeString = strings.Replace(encodeString, "/", "_", -1)
	encodeString = strings.Replace(encodeString, "+", "-", -1)

	URLString = "http://p" + strconv.Itoa(RandInt(1, 3)) + ".music.126.net/" + encodeString + "/" + id + ".mp3"
	return
}

// RandInt 生成min~max范围内随机整数
func RandInt(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return min + rand.Intn(max-min)
}

//Exist 是否存在
func Exist(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || os.IsExist(err)
}

func download(i Item, retry chan Item, ok chan int) {
	saveFile := i.Dir + clearSymbol(i.FileName)
	if Exist(saveFile) {
		log.Printf("%v is exist.\n", saveFile)
		<-ok
		return
	}
	err := os.MkdirAll(i.Dir, 0777)
	if err != nil {
		log.Print("downloadFile[1]:" + i.FileURL + " ====>>> " + err.Error())
		retry <- i
		return
	}
	resp, err := http.Get(i.FileURL)
	if err != nil {
		log.Print("downloadFile[2]:" + i.FileURL + " ====>>> " + err.Error())
		retry <- i
		return
	}
	defer resp.Body.Close()
	fd, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print("downloadFile[3]:" + i.FileURL + " ====>>> " + err.Error())
		retry <- i
		return
	}
	err = ioutil.WriteFile(saveFile, fd, 0666)
	if err != nil {
		log.Print("downloadFile[4]:" + i.FileURL + " ====>>> " + err.Error())
		retry <- i
		return
	}
	log.Println(saveFile + " downloaded.")
	<-ok
}

func clearSymbol(in string) (out string) {
	out = strings.Replace(in, ":", "", -1)
	out = strings.Replace(out, ",", "", -1)
	out = strings.Replace(out, "\"", "", -1)
	return
}

// iton parse the int to NO. with zero
func iton(no, all int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(len(strconv.Itoa(all)))+"d", no+1)
}