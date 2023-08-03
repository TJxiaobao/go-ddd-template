package errno

import "testing"

func TestBizError_Error(t *testing.T) {
	eo := NewSimpleError(ErrParameterInvalid, nil, "id")
	t.Log(eo.Error())

	eo = NewSimpleError(ErrMissingParameter, nil, "Code")
	t.Log(eo.Error())
}
