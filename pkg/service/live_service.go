package service

import (
	"context"
	"errors"
	"pql/pkg/constant"
	"pql/pkg/parse/live"
	"pql/pkg/vo"
	"slices"
	"strings"
	"sync"

	"github.com/wailsapp/wails/v3/pkg/application"
)

type LiveService struct {
	*ServiceContext
	remote *live.Live
}

func NewLiveService(sc *ServiceContext) *LiveService {
	return &LiveService{
		ServiceContext: sc,
	}
}

// Initialisation
func (ls *LiveService) ServiceStartup(ctx context.Context, options application.ServiceOptions) error {
	ls.remote = live.New(ls.Http)
	ttwid := ls.Store.GetStore(constant.Ttwid)
	ls.remote.SetTtwid(ttwid)
	return nil
}

func (ls *LiveService) ServiceShutdown() error {
	return nil
}

func (ls *LiveService) OpenWeb(roomId, typee string) {
	url := ls.remote.GetRoomWebUri(roomId, typee)
	if url != "" {
		ls.App.Browser.OpenURL(url)
	}
}

func (ls *LiveService) GetLiveList() []vo.LiveVo {
	lives := ls.Live.GetLiveList()
	if len(lives) <= 0 {
		return lives
	}
	var wg sync.WaitGroup
	for i, v := range lives {
		wg.Go(func() {
			info, err := ls.remote.GetInfo(v.RoomId, v.Type)
			if err != nil {
				return
			}
			lives[i].Rid = info.RoomId
			lives[i].RoomInfo = *info
		})
	}
	wg.Wait()
	slices.SortFunc(lives, liveSort)
	return lives
}

func (ls *LiveService) GetPlayInfo(roomId, roomtype string, extra map[string]string) (*live.LiveInfo, error) {
	return ls.remote.GetPlayUrl(roomId, roomtype, extra)
}

func (ls *LiveService) Search(params vo.LiveSearchVo) (*live.LiveResponse[live.RoomInfo], error) {
	return ls.remote.Search(live.LiveParams{
		Keyword: params.Keyword,
		Page:    params.Page,
		Type:    params.Type,
	})
}

func (ls *LiveService) GetLiveInfo(roomId, roomtype string) *vo.LiveVo {
	return ls.Live.GetLiveInfo(roomId, roomtype)
}

func (ls *LiveService) GetLiveRemoteInfo(roomId, roomtype string) (*live.RoomInfo, error) {
	return ls.remote.GetInfo(roomId, roomtype)
}

func (ls *LiveService) CreateLive(live vo.CreateLiveVo) error {
	if err := live.Valid(); err != nil {
		return err
	}

	_, err := ls.GetLiveRemoteInfo(live.RoomId, live.Type)
	if err != nil {
		return err
	}
	return ls.Live.CreateLive(live)
}

func (ls *LiveService) UpdateLive(live vo.UpdateLiveVo) error {
	return ls.Live.UpdateLive(live)
}

func (ls *LiveService) DeleteLive(ids []int) error {
	return ls.Live.DeleteLive(ids)
}

func (ls *LiveService) UpdateTtwid(ttwid string) error {
	if strings.TrimSpace(ttwid) == "" {
		return errors.New("ttwid不能为空！")
	}
	ls.remote.SetTtwid(ttwid)
	return ls.Store.SetStore(constant.Ttwid, ttwid)
}

func liveSort(a, b vo.LiveVo) int {
	if a.IsLive != b.IsLive {
		if a.IsLive {
			return -1
		}
		return 1
	}
	if a.Sort == nil && b.Sort == nil {
		return 0
	}
	if a.Sort != nil && b.Sort == nil {
		return -1
	}
	if a.Sort == nil && b.Sort != nil {
		return 1
	}
	if *a.Sort > *b.Sort {
		return -1
	}
	if a.CreatedAt.After(b.CreatedAt.Time) {
		return -1
	} else if a.CreatedAt.Before(b.CreatedAt.Time) {
		return 1
	}
	return 0
}
