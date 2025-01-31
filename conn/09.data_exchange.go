package conn

import (
	"github.com/Hypdncy/gordp/core"
	"github.com/Hypdncy/gordp/glog"
	"github.com/Hypdncy/gordp/proto/t128"
)

func (c *Client) readPdu() t128.PDU {
	glog.Debugf("before peek")
	defer func() { glog.Debugf("exit readPDU") }()
	d := c.stream.Peek(1)
	switch d[0] {
	case 3:
		glog.Debugf("read tpkt pdu begin")
		return t128.ReadPDU(c.stream)
	case 0:
		glog.Debugf("read fastpath pdu begin")
		return t128.ReadFastPathPDU(c.stream)
	default:
		core.Throw("invalid package")
	}
	return nil
}

func (c *Client) sendMouseEvent(pointerFlags uint16, xPos, yPos uint16) {
	pdu := t128.NewFastPathMouseInputPDU(pointerFlags, xPos, yPos)
	data := pdu.Serialize()
	glog.Debugf("send mouse event data: %v - %x:", len(data), data)
	core.WriteFull(c.stream, pdu.Serialize())
}

func (c *Client) SendMouseMoveEvent(xPos, yPos uint16) {
	c.sendMouseEvent(t128.PTRFLAGS_MOVE, xPos, yPos)
}

func (c *Client) SendMouseLeftDownEvent(xPos, yPos uint16) {
	c.sendMouseEvent(t128.PTRFLAGS_DOWN|t128.PTRFLAGS_BUTTON1, xPos, yPos)
}

func (c *Client) SendMouseLeftUpEvent(xPos, yPos uint16) {
	c.sendMouseEvent(t128.PTRFLAGS_BUTTON1, xPos, yPos)
}
