package main

import "fmt"

// 责任链模式：一般由链表或者数组(逐次去调用)去实现：特别适合多级审批业务，或者中间件
// 优点：责任链模式的优点是可以动态地建立责任链，增加或删除处理者都很方便，同时也能避免请求发送者和接受者之间的紧耦合关系。
// 缺点：可能会造成请求在整个链上被处理多次或未被处理
type BaseApprover struct {
	Successor Approver //下一个处理者
	Name      string   //名字
}

func (a *BaseApprover) SetApprover(approver Approver) {
	a.Successor = approver
}

type Approver interface {
	ProcessRequest(*PurchaseRequest)
}

type PurchaseRequest struct {
	ptype int     //请求类型
	price float64 //请求金额
	id    int
}

type DepartmentApprover struct {
	BaseApprover
}

func (d *DepartmentApprover) ProcessRequest(purchaseRequest *PurchaseRequest) {
	if purchaseRequest.price <= 5000 {
		fmt.Println(" 请求编号 id= ", purchaseRequest.id, " 被 ", d.Name+" 处理")
	} else {
		d.Successor.ProcessRequest(purchaseRequest)
	}
}

type CollegeApprover struct {
	BaseApprover
	Approver
}

func (c *CollegeApprover) ProcessRequest(purchaseRequest *PurchaseRequest) {
	if purchaseRequest.price > 5000 && purchaseRequest.price <= 10000 {
		fmt.Println(" 请求编号 id= ", purchaseRequest.id, " 被 ", c.Name+" 处理")
	} else {
		c.Successor.ProcessRequest(purchaseRequest)
	}
}

type ViceSchoolMasterApprover struct {
	BaseApprover
	Approver
}

func (v *ViceSchoolMasterApprover) ProcessRequest(purchaseRequest *PurchaseRequest) {
	if purchaseRequest.price > 10000 && purchaseRequest.price <= 30000 {
		fmt.Println(" 请求编号 id= ", purchaseRequest.id, " 被 ", v.Name+" 处理")
	} else {
		v.Successor.ProcessRequest(purchaseRequest)
	}
}

type SchoolMasterApprover struct {
	BaseApprover
	Approver
}

func (s *SchoolMasterApprover) ProcessRequest(purchaseRequest *PurchaseRequest) {
	if purchaseRequest.price > 30000 {
		fmt.Println(" 请求编号 id= ", purchaseRequest.id, " 被 ", s.Name+" 处理")
	} else {
		s.Successor.ProcessRequest(purchaseRequest)
	}
}
