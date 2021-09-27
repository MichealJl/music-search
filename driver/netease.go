package driver

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/MichealJl/music-search/proto"
	"github.com/MichealJl/music-search/proto/netease"
	"github.com/MichealJl/music-search/request"
	"io"
	"math"
	"net/url"
	"strings"
)

const (
	neteaseDomain     = "http://music.163.com"
	neteasePicDomain  = "https://p3.music.126.net"
	neteaseSearchUri  = "api/search/get/web"
	neteasePlayUri    = "song/media/outer/url"
	neteaseLyricUri   = "api/song/lyric"
	neteaseSourceName = "netease"
)

type neteaseDriver struct{}

func GetNetease() *neteaseDriver {
	return &neteaseDriver{}
}

func (n *neteaseDriver) Search(ctx context.Context, params *proto.SearchParams) (*proto.SearchRet, error) {
	client := request.NewHttpClient()
	client.SetMethod(request.Get)
	offset := (params.Page - 1) * params.PageSize
	queryParams := url.Values{}
	queryParams.Set("s", params.Keyword)
	queryParams.Set("limit", fmt.Sprintf("%d", params.PageSize))
	queryParams.Set("offset", fmt.Sprintf("%d", offset))
	queryParams.Set("total", "true")
	queryParams.Set("type", "1")
	reqUrl := fmt.Sprintf("%s/%s?%s", neteaseDomain, neteaseSearchUri, queryParams.Encode())
	rsp := &netease.SearchNeteaseRsp{}
	if err := client.Cal(ctx, reqUrl, rsp); err != nil {
		return nil, err
	}

	return n.format(rsp), nil
}

func (n *neteaseDriver) format(rsp *netease.SearchNeteaseRsp) *proto.SearchRet {
	ret := &proto.SearchRet{
		List: make([]proto.SearchRsItem, 0),
	}
	for _, song := range rsp.Result.Songs {
		item := proto.SearchRsItem{
			ID:       fmt.Sprintf("%d", song.ID),
			Name:     song.Name,
			Artist:   song.Artists[0].Name,
			Album:    song.Album.Name,
			PicID:    fmt.Sprintf("%d", song.Album.PicID),
			URLID:    fmt.Sprintf("%d", song.ID),
			LyricID:  fmt.Sprintf("%d", song.ID),
			Duration: int64(math.Ceil(float64(song.Duration) / float64(1000))),
			CoverUrl: fmt.Sprintf("%s/%s/%d.jpg?param=300y300", neteasePicDomain, n.encryptId(fmt.Sprintf("%d", song.Album.PicID)), song.Album.PicID),
			Source:   neteaseSourceName,
		}
		ret.List = append(ret.List, item)
	}

	return ret
}

func (n *neteaseDriver) encryptId(id string) string {
	magic := "3go8&$8*3*3h0k(2)2"
	songId := ""
	for i := 0; i < len(id); i++ {
		songId += string(magic[i%len(magic)] ^ id[i])
	}
	h := md5.New()
	io.WriteString(h, songId)
	result := base64.StdEncoding.EncodeToString(h.Sum(nil))
	result = strings.Replace(result, "/", "_", -1)
	result = strings.Replace(result, "+", "-", -1)

	return result
}

func (n *neteaseDriver) GetPlayUrl(ctx context.Context, id string) (*proto.GetPlayUrlRet, error) {
	return &proto.GetPlayUrlRet{
		Url:  fmt.Sprintf("%s/%s?id=%s.mp3", neteaseDomain, neteasePlayUri, id),
		Size: 0,
	}, nil
}

func (n *neteaseDriver) GetLyric(ctx context.Context, id string) (string, error) {
	client := request.NewHttpClient()
	client.Headers = map[string]string{
		"User-Agent": "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.82 Safari/537.36",
	}
	client.SetMethod(request.Get)
	queryParams := url.Values{}
	queryParams.Set("id", id)
	queryParams.Set("lv", "-1")
	reqUrl := fmt.Sprintf("%s/%s?%s", neteaseDomain, neteaseLyricUri, queryParams.Encode())
	rsp := &netease.LyricRsp{}
	if err := client.Cal(ctx, reqUrl, rsp); err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString([]byte(rsp.Lrc.Lyric)), nil
}
