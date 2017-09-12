package secret

import (
	"reflect"
	"testing"
)

const targetTestVersion = 2

func TestTestVersion(t *testing.T) {
	if testVersion != targetTestVersion {
		t.Fatalf("Found testVersion = %v, want %v", testVersion, targetTestVersion)
	}
}

func TestHandshake(t *testing.T) {
	for _, test := range tests {
		h := Handshake(test.code)
		// use len() to allow either nil or empty list, because
		// they are not equal by DeepEqual
		if len(h) == 0 && len(test.h) == 0 {
			continue
		}
		if !reflect.DeepEqual(h, test.h) {
			t.Fatalf("Handshake(%d) = %q, want %q.", test.code, h, test.h)
		}
	}
}

func BenchmarkHandshake(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Handshake(31)
	}
}

func Test_parseBinary(t *testing.T) {
	type args struct {
		code uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test 3", args: args{code: 3}, want: "11"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBinary(tt.args.code); got != tt.want {
				t.Errorf("parseBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
