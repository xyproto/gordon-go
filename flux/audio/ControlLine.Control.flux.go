// Generated by Flux, not meant for human consumption.  Editing may make it unreadable by Flux.

package audio

func (x *ControlLine) Control(a Audio) () {
	var v float64
	var v2 *float64
	var v3 *ControlLine
	var v4 Audio
	var v5 float64
	var v6 float64
	var v7 float64
	var v8 *ControlLine
	v8 = x
	v3 = x
	v4 = a
	x2 := &v8.x
	v = *x2
	v2 = x2
	v6 = *x2
	x3 := &v3.step
	v5 = *x3
	for i := range v4 {
		var v9 = &v4[i]
		var v10 *float64
		v10 = v9
		*v10 = v
		x4 := v6 + v5
		v7 = x4
		v6 = x4
		v = x4
	}
	*v2 = v7
	return
}
