package main

import "fmt"

// AbstractMessage 抽象化角色
type AbstractMessage interface {
	SendMessage(text, to string)
}

// MessageImplementer 实现化角色
type MessageImplementer interface {
	Send(text, to string)
}

// MessageSMS 具体实现化角色1
type MessageSMS struct{}

func ViaSMS() MessageImplementer {
	return &MessageSMS{}
}

func (*MessageSMS) Send(text, to string) {
	fmt.Printf("send %s to %s via SMS", text, to)
}

// MessageEmail 具体实现化角色2
type MessageEmail struct{}

func ViaEmail() MessageImplementer {
	return &MessageEmail{}
}

func (*MessageEmail) Send(text, to string) {
	fmt.Printf("send %s to %s via Email", text, to)
}

// CommonMessage 抽象化角色的实现1
type CommonMessage struct {
	method MessageImplementer
}

func NewCommonMessage(method MessageImplementer) *CommonMessage {
	return &CommonMessage{
		method: method,
	}
}

func (m *CommonMessage) SendMessage(text, to string) {
	m.method.Send(text, to)
}

// UrgencyMessage 抽象化角色的实现2
type UrgencyMessage struct {
	method MessageImplementer
}

func NewUrgencyMessage(method MessageImplementer) *UrgencyMessage {
	return &UrgencyMessage{
		method: method,
	}
}

func (m *UrgencyMessage) SendMessage(text, to string) {
	m.method.Send(fmt.Sprintf("[Urgency] %s", text), to)
}

func main() {
	//实现化角色1
	sms := ViaSMS()
	//实现化角色2
	email := ViaEmail()
	//抽象化角色1
	commonMessage := NewCommonMessage(email)
	commonMessage.SendMessage("抽象是普通信息，邮件方式实现", "关羽")
	//抽象化角色2
	urgencyMessage := NewUrgencyMessage(sms)
	urgencyMessage.SendMessage("抽象是紧急消息，短信方式实现", "张飞")
}
