// Generated by Flux, not meant for human consumption.  Editing may make it unreadable by Flux.

package audio

func (x *ControlConst) Control(a Audio) () {
	var v Audio
	var v2 *ControlConst
	var v3 float64
	v2 = x
	v = a
	x2 := &v2.Value
	v3 = *x2
	for i := range v {
		var v4 = &v[i]
		var v5 *float64
		v5 = v4
		*v5 = v3
	}
	return
}
