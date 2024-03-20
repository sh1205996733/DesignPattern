package main

import "fmt"

// 适配器模式:将某个接口转换成期望的另一个接口表示，主要的目的就是兼容性，其别名叫包装器wrapper
// 适配器类和适配者都需要实现目标接口，适配者需要聚合到适配器类，在适配器类执行目标接口中的方法时，在执行支配者中的具体方法时做一些其他操作。
// 主要分三类，对象适配器、类适配器、接口适配器
// 	接口适配器:重写接口
// 	类适配器:使用继承
// 	对象适配器:使用参数传递实现

// 类适配器
type Adapter interface {
	OutPut() string
	OutPut2(*Voltage) string
}

type Voltage struct {
}

func (c *Voltage) OutPut() string {
	return "220V"
}

type VoltageAdapter struct {
	v0 *Voltage
}

func (v *VoltageAdapter) OutPut() string {
	if v.v0 == nil {
		v.v0 = new(Voltage)
	}
	voltage := v.v0.OutPut()
	if voltage == "220V" {
		voltage = "5V"
	}
	return voltage
}

// 对象适配器 可以替代组合、推荐使用
func (v *VoltageAdapter) OutPut2(c *Voltage) string {
	voltage := c.OutPut()
	if voltage == "220V" {
		voltage = "5V"
	}
	return voltage
}

type Phone struct {
}

func (p Phone) Charging(adapter Adapter) {
	voltage := adapter.OutPut()
	//voltage := adapter.ConvTo2(new(Voltage))
	if voltage == "5V" {
		fmt.Println(voltage + "充电中")
		return
	}
	fmt.Println(voltage + "电压过大,无法充电")
}
