package main

import (
	."code.google.com/p/gordon-go/gui"
	"code.google.com/p/go.exp/go/types"
)

type valueNode struct {
	*ViewBase
	AggregateMouseHandler
	obj types.Object
	addr, indirect, set bool
	blk *block
	text Text
	val, in, out *port
	focused bool
}

func newValueNode(obj types.Object, addr, indirect, set bool) *valueNode {
	n := &valueNode{obj: obj, addr: addr, indirect: indirect, set: set}
	n.ViewBase = NewView(n)
	n.AggregateMouseHandler = AggregateMouseHandler{NewClickKeyboardFocuser(n), NewViewDragger(n)}
	n.text = NewText("")
	n.text.SetBackgroundColor(Color{0, 0, 0, 0})
	n.text.MoveCenter(Pt(0, n.text.Height() / 2))
	n.AddChild(n.text)
	if f, ok := obj.(field); ok {
		n.val = newInput(n, &types.Var{Name: "x", Type: f.recv})
		n.AddChild(n.val)
	} else if obj == nil {
		n.val = newInput(n, &types.Var{Name: "value"})
		n.AddChild(n.val)
		n.val.connsChanged = n.reform
	}
	n.reform()
	return n
}

func (n *valueNode) reform() {
	text := ""
	if n.obj != nil {
		if n.addr {
			text = "&"
		} else if n.indirect {
			text = "*"
		}
		if _, ok := n.obj.(field); ok {
			text += "."
		}
		text += n.obj.GetName()
	} else if n.addr {
		text = "addr"
	} else if n.indirect {
		text = "indirect"
	}
	n.text.SetText(text)
	
	if n.set {
		if n.in == nil {
			n.in = newInput(n, &types.Var{})
			n.AddChild(n.in)
			if n.val == nil {
				n.in.MoveCenter(Pt(-8 - 2*portSize, 0))
			} else {
				n.in.connsChanged = n.reform
				n.in.MoveCenter(Pt(-8 - 2*portSize, -portSize / 2))
				n.val.MoveCenter(Pt(-8 - 2*portSize, portSize / 2))
			}
		}
		if n.out != nil {
			for _, c := range n.out.conns {
				c.blk.removeConnection(c)
			}
			n.RemoveChild(n.out)
			n.out = nil
		}
	} else {
		if n.out == nil {
			n.out = newOutput(n, &types.Var{})
			n.AddChild(n.out)
			n.out.MoveCenter(Pt(8 + 2*portSize, 0))
			if n.val != nil {
				n.val.MoveCenter(Pt(-8 - 2*portSize, 0))
			}
		}
		if n.in != nil {
			for _, c := range n.in.conns {
				c.blk.removeConnection(c)
			}
			n.RemoveChild(n.in)
			n.in = nil
		}
	}
	var t types.Type
	if n.obj != nil {
		t = n.obj.GetType()
		if n.addr {
			t = &types.Pointer{Base: t}
		} else if n.indirect {
			t, _ = indirect(t)
		}
	} else {
		var valType types.Type
		if len(n.val.conns) > 0 {
			valType = n.val.conns[0].src.obj.Type
			t = valType
			if n.addr {
				t = &types.Pointer{Base: t}
			} else if n.indirect {
				t, _ = indirect(t)
			}
		} else if n.set && len(n.in.conns) > 0 {
			t = n.in.conns[0].src.obj.Type
			valType = t
			if n.addr {
				valType, _ = indirect(valType)
			} else if n.indirect {
				valType = &types.Pointer{Base: valType}
			}
		}
		n.val.setType(valType)
	}
	if n.set {
		n.in.setType(t)
	} else {
		n.out.setType(t)
	}
	ResizeToFit(n, 0)
}

func (n valueNode) block() *block { return n.blk }
func (n *valueNode) setBlock(b *block) { n.blk = b }
func (n valueNode) inputs() (p []*port) {
	if n.val != nil {
		p = []*port{n.val}
	}
	if n.set {
		p = append(p, n.in)
	}
	return
}
func (n valueNode) outputs() []*port {
	if n.set {
		return nil
	}
	return []*port{n.out}
}
func (n valueNode) inConns() (c []*connection) {
	for _, p := range n.inputs() {
		c = append(c, p.conns...)
	}
	return
}
func (n valueNode) outConns() []*connection {
	if n.set {
		return nil
	}
	return n.out.conns
}

func (n *valueNode) Move(p Point) {
	n.ViewBase.Move(p)
	for _, c := range append(n.inConns(), n.outConns()...) {
		c.reform()
	}
}

func (n *valueNode) TookKeyboardFocus() { n.focused = true; n.Repaint() }
func (n *valueNode) LostKeyboardFocus() { n.focused = false; n.Repaint() }

func (n *valueNode) KeyPressed(event KeyEvent) {
	switch event.Key {
	case KeyLeft, KeyRight, KeyUp, KeyDown:
		n.blk.outermost().focusNearestView(n, event.Key)
	case KeyEscape:
		n.blk.TakeKeyboardFocus()
	default:
		if _, ok := n.obj.(*types.Const); ok {
			n.ViewBase.KeyPressed(event)
		} else {
			switch event.Text {
			case "&":
				if !n.set {
					n.addr = !n.addr
					n.indirect = false
					n.reform()
				}
			case "*":
				var t types.Type
				if n.obj != nil {
					t = n.obj.GetType()
				} else {
					t = n.val.obj.Type
				}
				if _, ok := indirect(t); ok || t == nil {
					n.addr = false
					n.indirect = !n.indirect
					n.reform()
				}
			case "=":
				if !n.addr {
					n.set = !n.set
					n.reform()
					n.TakeKeyboardFocus()
				}
			default:
				n.ViewBase.KeyPressed(event)
			}
		}
	}
}

func (n valueNode) Paint() {
	SetColor(map[bool]Color{false:{.5, .5, .5, 1}, true:{.3, .3, .7, 1}}[n.focused])
	if n.val != nil {
		DrawLine(n.val.MapToParent(n.val.Center()), ZP)
	}
	if n.set {
		DrawLine(n.in.MapToParent(n.in.Center()), ZP)
	} else {
		DrawLine(ZP, n.out.MapToParent(n.out.Center()))
	}
}
