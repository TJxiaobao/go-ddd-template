package errno

import "testing"

func TestBizError_Error(t *testing.T) {
	eo := NewSimpleBizError(ErrParameterInvalid, nil, "ins_id")
	t.Log(eo.Error())

	eo = NewSimpleBizError(ErrMissingParameter, nil, "corpCode")
	t.Log(eo.Error())
}
