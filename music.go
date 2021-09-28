package music_search

import (
	"context"
	"github.com/MichealJl/music-search/driver"
	"github.com/MichealJl/music-search/proto"
)

type Platform int8

const (
	TencentPlatform Platform = iota // qq音乐
	NeteasePlatform                 //网易云
)

type Driver interface {
	Search(context.Context, *proto.SearchParams) (*proto.SearchRet, error)
	GetPlayUrl(context.Context, string) (*proto.GetPlayUrlRet, error)
	GetLyric(context.Context, string) (string, error)
}

var musicDriverMap = map[Platform]Driver{
	TencentPlatform: driver.GetQQDriver(),
	NeteasePlatform: driver.GetNetease(),
}

func GetMusicDriver(platform Platform) Driver {
	return musicDriverMap[platform]
}
