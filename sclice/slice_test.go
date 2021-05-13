package sclice

import (
	"reflect"
	"strconv"
	"testing"
)

type CheckStruct struct {
	Name string
	Age  int
}

func CheckF(v interface{}) (interface{}, error) {
	s := v.(CheckStruct)
	return s.Name + strconv.Itoa(s.Age), nil
}

func TestIntersect(t *testing.T) {
	type args struct {
		l1 []interface{}
		l2 []interface{}
		f  func(v interface{}) (interface{}, error)
	}

	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "struct",
			args: args{
				l1: []interface{}{
					CheckStruct{
						Name: "a",
						Age:  1,
					},
					CheckStruct{
						Name: "b",
						Age:  2,
					},
					CheckStruct{
						Name: "c",
						Age:  3,
					},
					CheckStruct{
						Name: "d",
						Age:  4,
					},
				},
				l2: []interface{}{
					CheckStruct{
						Name: "d",
						Age:  4,
					},
					CheckStruct{
						Name: "e",
						Age:  5,
					},
					CheckStruct{
						Name: "f",
						Age:  6,
					},
				},
				f: CheckF,
			},
			want: []interface{}{
				CheckStruct{
					Name: "d",
					Age:  4,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Intersect(tt.args.l1, tt.args.l2, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Intersect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinus(t *testing.T) {
	type args struct {
		l1 []interface{}
		l2 []interface{}
		f  func(v interface{}) (interface{}, error)
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "struct",
			args: args{
				l1: []interface{}{
					CheckStruct{
						Name: "a",
						Age:  1,
					},
					CheckStruct{
						Name: "b",
						Age:  2,
					},
					CheckStruct{
						Name: "c",
						Age:  3,
					},
					CheckStruct{
						Name: "d",
						Age:  4,
					},
				},
				l2: []interface{}{
					CheckStruct{
						Name: "d",
						Age:  4,
					},
					CheckStruct{
						Name: "e",
						Age:  5,
					},
					CheckStruct{
						Name: "f",
						Age:  6,
					},
				},
				f: CheckF,
			},
			want: []interface{}{
				CheckStruct{
					Name: "a",
					Age:  1,
				},
				CheckStruct{
					Name: "b",
					Age:  2,
				},
				CheckStruct{
					Name: "c",
					Age:  3,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Minus(tt.args.l1, tt.args.l2, tt.args.f)
			if (err != nil) != tt.wantErr {
				t.Errorf("Minus() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Minus() got = %v, want %v", got, tt.want)
			}
		})
	}
}
