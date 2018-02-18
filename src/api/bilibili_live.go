package api

import (
	"github.com/hr3lxphr6j/bililive-go/src/lib/http"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

const (
	biliBiliRoomApiUrl = "https://api.live.bilibili.com/room/v1/Room/get_info"
	biliBiliUserApiUrl = "https://api.live.bilibili.com/live_user/v1/UserInfo/get_anchor_in_room"
	biliBiliLiveApiUrl = "https://api.live.bilibili.com/room/v1/Room/playUrl"
)

type BiliBiliLive struct {
	Url *url.URL
}

func (b *BiliBiliLive) GetRoom() (*Info, error) {
	body, err := http.Get(biliBiliRoomApiUrl, map[string]string{
		"room_id": strings.Split(b.Url.Path, "/")[1],
		"from":    "room",
	})
	if err != nil {
		return nil, err
	}

	if gjson.GetBytes(body, "code").Int() != 0 {
		return nil, &RoomNotExistsError{b.Url}
	}
	info := &Info{
		Url:      b.Url,
		RoomName: gjson.GetBytes(body, "data.title").String(),
		Status:   gjson.GetBytes(body, "data.live_status").Int() == 1,
	}

	body2, err := http.Get(biliBiliUserApiUrl, map[string]string{
		"roomid": strings.Split(b.Url.Path, "/")[1],
	})
	if err != nil {
		return nil, err
	}

	info.HostName = gjson.GetBytes(body2, "data.info.uname").String()
	return info, nil
}

func (b *BiliBiliLive) GetUrls() ([]*url.URL, error) {
	body, err := http.Get(biliBiliLiveApiUrl, map[string]string{
		"cid":      strings.Split(b.Url.Path, "/")[1],
		"quality":  "0",
		"platform": "web",
	})
	if err != nil {
		return nil, err
	}

	urls := gjson.GetBytes(body, "data.durl.#.url").Array()

	us := make([]*url.URL, 0, 4)
	for _, u := range urls {
		url_, _ := url.Parse(u.String())
		us = append(us, url_)
	}
	return us, nil
}
