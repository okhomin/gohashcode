package gohashcode

import (
	"testing"
)

func TestHashcode(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{
			name: "Int",
			args: args{
				v: 123,
			},
			want: 3813,
		},
		{
			name: "Pointer Int",
			args: args{
				v: func() *int {
					i := 123
					return &i
				}(),
			},
			want: 3813,
		},
		{
			name: "String",
			args: args{
				v: "Hello",
			},
			want: 2358303207,
		},
		{
			name: "Byte",
			args: args{
				v: byte(123),
			},
			want: 3813,
		},
		{
			name: "Slice",
			args: args{
				v: []int{1, 2, 3},
			},
			want: 240343,
		},
		{
			name: "Map",
			args: args{
				v: map[string]int{
					"one": 1,
					"two": 2,
				},
			},
			want: 108091655972,
		},
		{
			name: "Struct",
			args: args{
				v: func() any {
					type S struct {
						A int
						B string
						C []uint32
						E map[string]any
					}
					type Sub struct {
						A int
						B string
						C []uint32
						E map[string]any
					}
					return S{
						A: 123,
						B: "Hello",
						C: []uint32{1, 2, 3},
						E: map[string]any{
							"one": 1,
							"two": "two",
							"three": func() *map[string]string {
								m := map[string]string{
									"one": "one",
								}
								return &m
							}(),
							"sub": Sub{
								A: 123,
								B: "Hello",
								C: []uint32{1, 2, 3},
								E: map[string]any{
									"one": 1,
									"two": "two",
								},
							},
							"sub-ptr": &Sub{
								A: 123,
								B: "Hello",
								C: []uint32{1, 2, 3},
								E: map[string]any{
									"one": 1,
								},
							},
							"sub-slice": []Sub{
								{
									A: 123,
									B: "Hello",
								},
							},
							"sub-nil": (*Sub)(nil),
						},
					}
				}(),
			},
			want: 15127202206556940116,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hashcode(tt.args.v); got != tt.want {
				t.Errorf("Hashcode() = %v, want %v", got, tt.want)
			}
		})
	}
}
