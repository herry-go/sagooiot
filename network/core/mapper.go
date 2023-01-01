package core

import (
	"context"
	"github.com/sagoo-cloud/sagooiot/internal/consts"
	logicModel "github.com/sagoo-cloud/sagooiot/internal/model"
	"github.com/sagoo-cloud/sagooiot/network/model"
	"strconv"
)

// todo 需要处理下心跳和注册包，需要考虑到客户端可能不支持心跳或者注册包
func MapperServer(ctx context.Context, res logicModel.NetworkServerOut) model.Server {
	s := model.Server{
		Id:        res.Id,
		Name:      res.Name,
		Type:      res.Types,
		Addr:      res.Addr,
		Register:  model.RegisterPacket{},
		Heartbeat: model.HeartBeatPacket{},
		Protocol:  model.Protocol{},
		//TODO 这里暂时不写device，需要等待后续的device插入数据，考虑是不是启动的时候带入
		//Devices:  res.Devices,
		Disabled: true,
		Created:  res.CreatedAt.Time,
	}
	//TODO 等待model和前端修改补充下mqtt的options，主要是额外的一些配置
	if res.Status == consts.ServerStatusOnline {
		s.Disabled = false
	}
	StrToPointInterfaceWithoutError(ctx, res.Register, &s.Register)
	StrToPointInterfaceWithoutError(ctx, res.Heartbeat, &s.Heartbeat)
	StrToPointInterfaceWithoutError(ctx, res.Protocol, &s.Protocol)
	return s
}

func MapperDevice(res logicModel.DeviceOutput) model.Device {
	return model.Device{
		Id:        uint64(res.Id),
		TunnelId:  uint64(res.TunnelId),
		ProductId: strconv.Itoa(int(res.ProductId)),
		Name:      res.Name,
		//todo 这里的station等待后面处理
		Station:  res.Status,
		Disabled: false,
		Created:  res.CreatedAt.Time,
	}
}

func mapperProduct(res logicModel.DetailProductOutput) model.Product {
	return model.Product{
		Id:           strconv.Itoa(int(res.Id)),
		Name:         res.Name,
		Manufacturer: res.CategoryName,
		//Version:      res.,
		Protocol: model.Protocol{
			Name: res.MessageProtocol,
			//TODO 消息的一些其他的配置
			Options: nil,
		},
		Tags:     nil,
		Pollers:  nil,
		Commands: nil,
		Created:  res.CreatedAt.Time,
	}
}

func mapperTunnel(ctx context.Context, res logicModel.NetworkTunnelOut) model.Tunnel {
	t := model.Tunnel{
		Id:       uint64(res.Id),
		ServerId: res.ServerId,
		Name:     res.Name,
		//SN:        res.SN,
		Type:      res.Types,
		Addr:      res.Addr,
		Remote:    res.Remote,
		Retry:     model.Retry{},
		Heartbeat: model.HeartBeatPacket{},
		Serial:    model.SerialOptions{},
		Protocol:  model.Protocol{},
		Disabled:  res.Status == 0,
		Last:      res.Last.Time,
		Created:   res.CreatedAt.Time,
	}

	StrToPointInterfaceWithoutError(ctx, res.Retry, &t.Retry)
	StrToPointInterfaceWithoutError(ctx, res.Heartbeat, &t.Heartbeat)
	StrToPointInterfaceWithoutError(ctx, res.Serial, &t.Serial)
	StrToPointInterfaceWithoutError(ctx, res.Protoccol, &t.Protocol)
	return t
}
