// Code generated by "stringer -type=Optimizer"; DO NOT EDIT.

package manduco

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Ect-0]
	_ = x[Oxipng-1]
}

const _Optimizer_name = "EctOxipng"

var _Optimizer_index = [...]uint8{0, 3, 9}

func (i Optimizer) String() string {
	if i < 0 || i >= Optimizer(len(_Optimizer_index)-1) {
		return "Optimizer(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Optimizer_name[_Optimizer_index[i]:_Optimizer_index[i+1]]
}
