package main

import "fmt"

// 适配器模式:将某个接口转换成期望的另一个接口表示，主要的目的就是兼容性，其别名叫包装器wrapper
// 主要分三类，对象适配器、类适配器、接口适配器
// 	接口适配器:重写接口
// 	类适配器:使用继承
// 	对象适配器:使用参数传递实现

// 类适配器
type Adapter interface {
	ConvTo() string
}

type Voltage struct {
}

func (c *Voltage) OutPut() string {
	return "220V"
}

type VoltageAdapter struct {
	Voltage
}

func (v *VoltageAdapter) ConvTo() string {
	voltage := v.OutPut()
	if voltage == "220V" {
		voltage = "5V"
	}
	return voltage
}

// 对象适配器 可以替代组合、推荐使用
//func (v *VoltageAdapter) ConvTo2(c Voltage) string {
//	voltage := c.OutPut()
//	if voltage == "220V" {
//		voltage = "5V"
//	}
//	return voltage
//}

type VoltageAdapter2 struct {
	Voltage
}

func (v *VoltageAdapter2) ConvTo() string {
	voltage := v.OutPut()
	if voltage == "220V" {
		voltage = "10V"
	}
	return voltage
}

type Phone struct {
}

func (p Phone) Charging(adapter Adapter) {
	voltage := adapter.ConvTo()
	if voltage == "5V" {
		fmt.Println(voltage + "充电中")
		return
	}
	fmt.Println(voltage + "电压过大,无法充电")
}
