package api

import (
	"net/url"
	"testing"
)

const bilibiliTestUrl = "https://api.bilibili.com/8290927"

func TestBiliBiliLive_GetRoom(t *testing.T) {
	u, _ := url.Parse(bilibiliTestUrl)
	t.Log((&BiliBiliLive{u}).GetRoom())
}

func TestBiliBiliLive_GetUrl(t *testing.T) {
	u, _ := url.Parse(bilibiliTestUrl)
	t.Log((&BiliBiliLive{u}).GetUrls())
}
