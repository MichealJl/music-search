package music_search

import (
	"context"
	"github.com/MichealJl/music-search/proto"
	"testing"
	"time"
)

func TestMusic(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	driver := GetMusicDriver(TencentPlatform)
	rsp, err := driver.Search(ctx, &proto.SearchParams{
		Keyword:  "不负人间",
		Page:     1,
		PageSize: 20,
	})
	if err != nil {
		t.Fatal("Search music fatal", err)
		return
	}
	if len(rsp.List) <= 0 {
		t.Log("search result null", rsp)
		return
	}
	// 获取播放连接
	urlRsp, urlErr := driver.GetPlayUrl(ctx, rsp.List[0].ID)
	if urlErr != nil {
		t.Fatal("GetPlayUrl id :", rsp.List[0].ID, "GetPlayUrl id err", urlErr)
	}
	t.Log("播放链接返回地址：", urlRsp.Url)
	// 获取歌词
	lyricRsp, lyricErr := driver.GetLyric(ctx, rsp.List[0].ID)
	if lyricErr != nil {
		t.Fatal("GetLyric id :", rsp.List[0].ID, "GetLyric id err", lyricErr)
	}
	t.Log("返回地址base64：", lyricRsp)
}
