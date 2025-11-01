package types_test

import (
	"testing"

	"github.com/ManasRaut/lexe/types"
)

type stackActionType string

const (
	push   stackActionType = "push"
	pop    stackActionType = "pop"
	length stackActionType = "len"
	clear  stackActionType = "clear"
	top    stackActionType = "top"
)

type stackAction struct {
	t          stackActionType
	value      int
	wantCurr   int
	wantLen    int
	wantErr    error
	wantReturn int
}

func TestStack(t *testing.T) {
	tests := []struct {
		name    string
		actions []stackAction
	}{
		{name: "All operations", actions: []stackAction{
			{t: length, value: 0, wantCurr: 0, wantLen: 0, wantErr: nil, wantReturn: 0},
			{t: push, value: -5, wantCurr: -5, wantLen: 1, wantErr: nil, wantReturn: 0},
			{t: length, value: 0, wantCurr: -5, wantLen: 1, wantErr: nil, wantReturn: 1},
			{t: push, value: 455456465, wantCurr: 455456465, wantLen: 2, wantErr: nil, wantReturn: 0},
			{t: pop, value: 0, wantCurr: -5, wantLen: 1, wantErr: nil, wantReturn: 455456465},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := types.NewStack[int](0)
			for _, action := range tt.actions {
				var tempVal *int = nil
				var gotReturn int = 0
				var gotErr error = nil
				switch action.t {
				case push:
					s.Push(action.value)
				case pop:
					tempVal, gotErr = s.Pop()
					gotReturn = *tempVal
				case length:
					gotReturn = s.Len()
				case clear:
					s.Clear()
				case top:
					tempVal, gotErr = s.Top()
					gotReturn = *tempVal
				}

				if action.wantErr != gotErr {
					t.Fatalf("Failed stack operation %s : got error %v, want %v", action.t, gotErr, action.wantErr)
				}
				if action.wantReturn != gotReturn {
					t.Fatalf("Failed stack operation %s : got return value %v, want %v", action.t, gotReturn, action.wantReturn)
				}
				gotLen := s.Len()
				if action.wantLen != gotLen {
					t.Fatalf("Failed stack operation %s : got length %v, want %v", action.t, gotLen, action.wantLen)
				}

				gotCurr, gotErr := s.Top()
				if gotLen > 0 {
					if action.wantCurr != *gotCurr {
						t.Fatalf("Failed stack operation %s : got top element %v, want %v", action.t, *gotCurr, action.wantCurr)
					}
				} else {
					if gotErr == nil {
						t.Fatalf("Failed stack operation %s : got error nill, want error %v", action.t, types.ErrInvalidStackOperation)
					}
				}
			}
		})
	}
}
