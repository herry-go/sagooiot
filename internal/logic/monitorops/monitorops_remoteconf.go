package monitorops

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	"github.com/sagoo-cloud/sagooiot/internal/dao"
	"github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/internal/model/do"
	"github.com/sagoo-cloud/sagooiot/internal/model/entity"
	"github.com/sagoo-cloud/sagooiot/internal/service"
)

type sMonitoropsRemoteconf struct{}

func sMonitoropsRemoteconfNew() *sMonitoropsRemoteconf {
	return &sMonitoropsRemoteconf{}
}
func init() {
	service.RegisterMonitoropsRemoteconf(sMonitoropsRemoteconfNew())
}

// GetRemoteconfList 获取列表数据
func (s *sMonitoropsRemoteconf) GetRemoteconfList(ctx context.Context, in *model.GetRemoteconfListInput) (list []*model.RemoteconfOutput, err error) {
	err = g.Try(ctx, func(ctx context.Context) {
		m := dao.Remoteconf.Ctx(ctx)
		c := dao.Remoteconf.Columns()
		if err != nil {
			err = gerror.New("获取总行数失败")
			return
		}
		if in.ProductKey != "" {
			m = m.Where(c.ProductKey, in.ProductKey)
		}
		err = m.Order("utc_create desc").Scan(&list)
		if err != nil {
			err = gerror.New("获取数据失败")
		}
		for i := range list {
			list[i].ConfigNumber = fmt.Sprintf("%02d", len(list)-i)
		}
	})
	return
}

// GetRemoteconfById 获取指定ID数据
func (s *sMonitoropsRemoteconf) GetRemoteconfById(ctx context.Context, id int) (out *model.RemoteconfOutput, err error) {
	err = dao.Remoteconf.Ctx(ctx).Where(dao.Remoteconf.Columns().Id, id).Scan(&out)
	return
}

// AddRemoteconf 添加数据
func (s *sMonitoropsRemoteconf) AddRemoteconf(ctx context.Context, in model.RemoteconfAddInput) (err error) {
	var p []*entity.DevProduct
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, in.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}

	var param *do.Remoteconf
	err = gconv.Scan(in, &param)
	if err != nil {
		glog.Error(ctx, err)
		return
	}

	param.Id = guid.S()
	param.UtcCreate = gtime.Now().UTC()
	param.GmtCreate = gtime.Now().Format("Y/m/d H:i:s")
	param.ConfigName = fmt.Sprintf("%d", gtime.Now().UnixMilli())
	_, err = dao.Remoteconf.Ctx(ctx).Insert(param)
	if err != nil {
		glog.Error(ctx, err)
		return
	}

	// 只保留最新11条记录
	var r []*entity.Remoteconf
	err = dao.Remoteconf.Ctx(ctx).Where(dao.Remoteconf.Columns().ProductKey, in.ProductKey).OrderDesc(dao.Remoteconf.Columns().UtcCreate).Scan(&r)
	if err != nil {
		glog.Error(ctx, err)
		return
	}
	ids := []string{}
	if r != nil && len(r) <= 11 {
		return
	}
	for _, v := range r {
		ids = append(ids, v.Id)
	}
	_, err = dao.Remoteconf.Ctx(ctx).Delete(dao.Remoteconf.Columns().Id+" in (?)", ids[11:])
	if err != nil {
		glog.Error(ctx, err)
		return
	}
	return
}

// EditRemoteconf 修改数据
func (s *sMonitoropsRemoteconf) EditRemoteconf(ctx context.Context, in model.RemoteconfEditInput) (err error) {
	var c *entity.Remoteconf
	err = dao.Remoteconf.Ctx(ctx).Where(dao.DevProduct.Columns().Id, in.Id).Scan(&c)
	if err != nil {
		return
	}
	if c == nil {
		return gerror.New("远程配置文件不存在")
	}
	var p []*entity.DevProduct
	err = dao.DevProduct.Ctx(ctx).Where(dao.DevProduct.Columns().Key, c.ProductKey).Scan(&p)
	if err != nil {
		return
	}
	if p == nil {
		return gerror.New("产品不存在")
	}
	//todo 向所有设备推送配置文件
	if c.Scope == consts.DeviceScopeProduct {

	}


	return
}

// DeleteRemoteconf 删除数据
func (s *sMonitoropsRemoteconf) DeleteRemoteconf(ctx context.Context, Ids []int) (err error) {
	_, err = dao.Remoteconf.Ctx(ctx).Delete(dao.Remoteconf.Columns().Id+" in (?)", Ids)
	return
}
