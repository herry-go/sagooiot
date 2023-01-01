package core

import (
	"context"
	"errors"
	"github.com/sagoo-cloud/sagooiot/network/events"
	"github.com/sagoo-cloud/sagooiot/network/model"
)

// Device 设备
type Device struct {
	model.Device
	events.EventEmitter

	product *model.Product

	Context map[string]interface{}

	pollers []*Poller

	//命令索引
	commandIndex map[string]*model.Command

	running bool
	tunnel  *Tunnel
}

func NewDevice(ctx context.Context, m *model.Device) (*Device, error) {
	dev := &Device{
		Device:       *m,
		Context:      make(map[string]interface{}),
		commandIndex: make(map[string]*model.Command, 0),
		pollers:      make([]*Poller, 0),
	}

	//加载产品
	var err error
	dev.product, err = LoadProduct(ctx, dev.ProductId)
	if err != nil {
		return nil, err
	}

	//索引命令
	for _, cmd := range dev.product.Commands {
		dev.commandIndex[cmd.Name] = cmd
	}

	//初始化
	for _, v := range dev.product.Pollers {
		dev.pollers = append(dev.pollers, &Poller{Poller: *v, Device: dev})
	}

	return dev, nil
}

func (dev *Device) BindTunnel(tunnel *Tunnel) error {
	if tunnel == nil {
		return errors.New("通道未加载")
	}
	dev.tunnel = tunnel
	return nil
}

func (dev *Device) onData(data map[string]interface{}) {

	//向上广播
	dev.Emit("data", data)
}

func (dev *Device) Start(ctx context.Context) error {
	tunnel := GetTunnel(int(dev.TunnelId))
	if tunnel == nil {
		return errors.New("找不到链接")
	}
	err := dev.BindTunnel(tunnel)
	if err != nil {
		return err
	}
	for _, poller := range dev.pollers {
		err := poller.Start(ctx)
		if err != nil {
			return err
		}
	}

	dev.running = true

	return nil
}

func (dev *Device) Stop() error {
	dev.running = false

	for _, poller := range dev.pollers {
		poller.Stop()
	}
	return nil
}

func (dev *Device) Running() bool {
	return dev.running
}

func (dev *Device) read(size int) error {
	//todo  等待实现
	return nil
}

func (dev *Device) Refresh(ctx context.Context) error {
	if !dev.running {
		return errors.New("设备未运行")
	}
	for _, poller := range dev.pollers {
		poller.Execute(ctx)
	}
	return nil
}
