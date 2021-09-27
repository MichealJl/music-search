package driver

import (
	"context"
	"fmt"
	"github.com/MichealJl/music-search/proto"
	qqProto "github.com/MichealJl/music-search/proto/qq"
	"github.com/MichealJl/music-search/request"
	"github.com/fatih/structs"
	"math/rand"
	"net/url"
	"time"
)

const (
	cqqDomain        = "https://c.y.qq.com"                  // 获取音乐信息所需域名
	uqqDomain        = "https://u.y.qq.com"                  // 获取播放地址信息所需域名
	qqSearchUri      = "soso/fcgi-bin/client_search_cp"      // 搜索音乐地址
	qqGetPlayInfoUri = "v8/fcg-bin/fcg_play_single_song.fcg" // 获取音乐播放信息
	qqGetPlayUrlUri  = "cgi-bin/musicu.fcg"                  // 获取音乐播放地址
	qqSourceName     = "tencent"
	coverDomain      = "https://y.gtimg.cn"                     // qq音乐封面 域名
	lyricUri         = "lyric/fcgi-bin/fcg_query_lyric_new.fcg" // qq歌词
)

type qqDriver struct{}

func GetQQDriver() *qqDriver {
	return &qqDriver{}
}

func (qq *qqDriver) Search(ctx context.Context, params *proto.SearchParams) (*proto.SearchRet, error) {
	reqUrl := fmt.Sprintf("%s/%s", cqqDomain, qqSearchUri)
	req := request.NewHttpClient()
	req.SetMethod(request.Get)
	queryParams := url.Values{}
	queryParams.Set("format", "json")
	queryParams.Set("p", fmt.Sprintf("%d", params.Page))
	queryParams.Set("n", fmt.Sprintf("%d", params.PageSize))
	queryParams.Set("w", params.Keyword)
	queryParams.Set("aggr", "1")
	queryParams.Set("lossless", "1")
	queryParams.Set("cr", "1")
	queryParams.Set("new_json", "1")
	reqUrl = fmt.Sprintf("%s?%s", reqUrl, queryParams.Encode())
	rsp := &qqProto.QQSearchRsp{}
	err := req.Cal(ctx, reqUrl, rsp)
	if err != nil {
		return nil, err
	}

	return qq.format(rsp), nil
}

// GetPlayUrl 获取播放地址信息
func (qq *qqDriver) GetPlayUrl(ctx context.Context, mid string) (*proto.GetPlayUrlRet, error) {
	playInfoRsp, err := qq.getPlayInfo(ctx, mid)
	if err != nil {
		return nil, err
	}
	req := request.NewHttpClient()
	req.SetMethod(request.Get)
	reqUrl := fmt.Sprintf("%s/%s", uqqDomain, qqGetPlayUrlUri)
	queryParams := url.Values{}
	queryParams.Set("needNewCode", "0")
	queryParams.Set("platform", "yqq.json")
	queryParams.Set("format", "json")
	queryParams.Set("data", buildDataParams(playInfoRsp))
	reqUrl = fmt.Sprintf("%s?%s", reqUrl, queryParams.Encode())
	rsp := &qqProto.GetPlayURLRsp{}
	if er := req.Cal(ctx, reqUrl, rsp); er != nil {
		return nil, er
	}
	vKeys := rsp.Req0.Data.Midurlinfo
	fileTypes := qqProto.GetFileTypeList()
	ret := &proto.GetPlayUrlRet{}
	for i, item := range fileTypes {
		if i <= len(vKeys)-1 && vKeys[i].Vkey != "" {
			ret.Url = rsp.Req0.Data.SIP[0] + vKeys[i].Purl
			s := structs.New(playInfoRsp.Data[0].File)
			val, ok := s.FieldOk(item.Key)
			if !ok {
				continue
			}
			ret.Size = val.Value().(int64)
			break
		}
	}

	return ret, nil
}

// getPlayInfo 获取播放歌曲信息
func (qq *qqDriver) getPlayInfo(ctx context.Context, mid string) (*qqProto.GetPlayInfoRsp, error) {
	req := request.NewHttpClient()
	req.SetMethod(request.Get)
	reqUrl := fmt.Sprintf("%s/%s", cqqDomain, qqGetPlayInfoUri)
	queryParams := url.Values{}
	queryParams.Set("songmid", mid)
	queryParams.Set("platform", "yqq")
	queryParams.Set("format", "json")
	reqUrl = fmt.Sprintf("%s?%s", reqUrl, queryParams.Encode())
	rsp := &qqProto.GetPlayInfoRsp{}
	if err := req.Cal(ctx, reqUrl, rsp); err != nil {
		return nil, err
	}

	return rsp, nil
}

// 获取封面图
func (qq *qqDriver) getCoverUrl(mid string) (cover string) {
	size := 300
	cover = fmt.Sprintf("%s/%s%dx%dM000%s.jpg?max_age=2592000", coverDomain, "music/photo_new/T002R", size, size, mid)
	return
}

// GetLyric 获取qq音乐歌词
func (qq *qqDriver) GetLyric(ctx context.Context, mid string) (string, error) {
	rsp := &qqProto.LyricRsp{}
	client := request.NewHttpClient()
	client.SetMethod(request.Get)
	client.Headers = map[string]string{
		"Referer": "https://y.qq.com/",
		"Host":    "c.y.qq.com",
	}
	queryParams := url.Values{}
	queryParams.Set("songmid", mid)
	queryParams.Set("format", "json")
	reqUrl := fmt.Sprintf("%s/%s?%s", cqqDomain, lyricUri, queryParams.Encode())
	if err := client.Cal(ctx, reqUrl, rsp); err != nil {
		return "", err
	}

	return rsp.Lyric, nil
}

func (qq *qqDriver) format(rsp *qqProto.QQSearchRsp) *proto.SearchRet {
	ret := &proto.SearchRet{
		List: make([]proto.SearchRsItem, 0),
	}
	for _, song := range rsp.Data.Song.List {
		artistName := ""
		if len(song.Singer) > 0 {
			artistName = song.Singer[0].Name
		}
		ret.List = append(ret.List, proto.SearchRsItem{
			ID:       song.Mid,
			Name:     song.Name,
			Artist:   artistName,
			Album:    song.Album.Title,
			PicID:    song.Album.Mid,
			URLID:    song.Mid,
			LyricID:  song.Mid,
			Duration: song.Interval,
			Source:   qqSourceName,
			CoverUrl: qq.getCoverUrl(song.Album.Mid),
		})
	}

	return ret
}

func buildDataParams(playInfoRsp *qqProto.GetPlayInfoRsp) string {
	dataParams := qqProto.GetPlayURLDataParam{}
	dataParams.Req0.Module = "vkey.GetVkeyServer"
	dataParams.Req0.Method = "CgiGetVkey"
	dataParams.Req0.Param.GUID = fmt.Sprintf("%d", getGuid())
	dataParams.Req0.Param.Uin = "0"
	dataParams.Req0.Param.Loginflag = 1
	dataParams.Req0.Param.Platform = "20"
	fileTypes := qqProto.GetFileTypeList()
	for _, item := range fileTypes {
		dataParams.Req0.Param.Songmid = append(dataParams.Req0.Param.Songmid, playInfoRsp.Data[0].Mid)
		dataParams.Req0.Param.Filename = append(dataParams.Req0.Param.Filename, fmt.Sprintf("%s%s.%s", item.Pre, playInfoRsp.Data[0].File.MediaMid, item.Ext))
		dataParams.Req0.Param.Songtype = append(dataParams.Req0.Param.Songtype, playInfoRsp.Data[0].Type)
	}
	data, _ := dataParams.Marshal()

	return string(data)
}

func getGuid() int64 {
	max := int64(9999999999)
	min := int64(1000000000)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min+1) + min
}
