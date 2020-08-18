package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"reflect"
	"testing"
)

func Test_isIntsEnough(t *testing.T) {
	type args struct {
		ints     []int
		minCount int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ints is nil, must return false", args: args{ints: nil, minCount: 5}, want: false,
		},
		{
			name: "ints has len() == 5, min = 6 , must return false", args: args{ints: []int{1, 2, 3, 4, 5}, minCount: 6}, want: false,
		},
		{
			name: "ints has len() == 5, min = 5, must return true", args: args{ints: []int{1, 2, 3, 4, 5}, minCount: 5}, want: true,
		},
		{
			name: "ints has len() == 6, min = 5, must return true", args: args{ints: []int{1, 2, 3, 4, 5, 6}, minCount: 5}, want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIntsEnough(tt.args.ints, tt.args.minCount); got != tt.want {
				t.Errorf("isIntsEnough() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeJSON(t *testing.T) {
	type args struct {
		reader io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name: "valid, expected equal values, not error", args: args{reader: getTest_decodeJSONReader1()}, want: []int{1, 2, 3, 4}, wantErr: false,
		},
		{
			name: "invalid, expected error", args: args{reader: getTest_decodeJSONReader2()}, want: nil, wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeJSON(tt.args.reader)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func getTest_decodeJSONReader1() io.Reader {
	var intsPayload initValues
	intsPayload.Ints = []int{1, 2, 3, 4}

	bb, _ := json.Marshal(intsPayload)

	return bytes.NewBuffer(bb)
}

func getTest_decodeJSONReader2() io.Reader {

	return bytes.NewBuffer([]byte{1, 23, 4})
}

func TestWithMinCount(t *testing.T) {
	type args struct {
		minCount int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "01", args: args{minCount: 5}, want: 5,
		},
		{
			name: "02", args: args{minCount: 6}, want: 6,
		},
		{
			name: "03", args: args{minCount: 0}, want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := WithMinCount(tt.args.minCount); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("WithMinCount() = %v, want %v", got, tt.want)
			// }

			testOptsStruct := opts{}

			option := WithMinCount(tt.args.minCount)
			option(&testOptsStruct)

			if testOptsStruct.minCount != tt.want {
				t.Errorf("WithMinCount(opt).minValue = %v, want %d", testOptsStruct.minCount, tt.want)
			}
		})
	}
}

func TestFromReader(t *testing.T) {
	t.Run("decodeJSON returns json, must return error", func(t *testing.T) {
		buf := bytes.NewBuffer([]byte{1, 3, 4})

		_, err := FromReader(buf)
		if err == nil {
			t.Error("expected not nil error, got nil")
		}
	})

	t.Run("ints is not enough, must return error", func(t *testing.T) {
		var intsPayload struct {
			Ints []int `json"ints"`
		}
		intsPayload.Ints = make([]int, 29)

		bb, _ := json.Marshal(intsPayload)

		buf := bytes.NewBuffer(bb)

		_, err := FromReader(buf)
		if err == nil {
			t.Error("expected not nil error, got nil")
		}
	})

	t.Run("must return not nil tree, not error", func(t *testing.T) {
		var intsPayload struct {
			Ints []int `json"ints"`
		}
		intsPayload.Ints = make([]int, 30)

		bb, _ := json.Marshal(intsPayload)

		buf := bytes.NewBuffer(bb)

		tree, err := FromReader(buf, WithMinCount(30))
		if err != nil {
			t.Errorf("expected nil error, got %v", err)
		}

		if tree == nil {
			t.Error("expected not nil tree, got nil")
		}
	})
}
