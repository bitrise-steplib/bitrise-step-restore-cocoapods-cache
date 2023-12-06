// Code generated by "stringer -type=ExitCode"; DO NOT EDIT.

package exitcode

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Success-0]
	_ = x[Failure-1]
}

const _ExitCode_name = "SuccessFailure"

var _ExitCode_index = [...]uint8{0, 7, 14}

func (i ExitCode) String() string {
	if i < 0 || i >= ExitCode(len(_ExitCode_index)-1) {
		return "ExitCode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ExitCode_name[_ExitCode_index[i]:_ExitCode_index[i+1]]
}