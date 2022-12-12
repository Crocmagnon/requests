// Code generated by "stringer -type=ErrorKind"; DO NOT EDIT.

package requests

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[ErrUnknown-0]
	_ = x[ErrURL-1]
	_ = x[ErrRequest-2]
	_ = x[ErrConnect-3]
	_ = x[ErrValidator-4]
	_ = x[ErrHandler-5]
}

const _ErrorKind_name = "ErrUnknownErrURLErrRequestErrConnectErrValidatorErrHandler"

var _ErrorKind_index = [...]uint8{0, 10, 16, 26, 36, 48, 58}

func (i ErrorKind) String() string {
	if i < 0 || i >= ErrorKind(len(_ErrorKind_index)-1) {
		return "ErrorKind(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ErrorKind_name[_ErrorKind_index[i]:_ErrorKind_index[i+1]]
}
