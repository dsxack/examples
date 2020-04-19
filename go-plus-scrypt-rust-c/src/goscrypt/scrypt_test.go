package goscrypt

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRustScryptKey(t *testing.T) {
	type args struct {
		password []byte
		salt     []byte
		r        int
		p        int
		n        int
	}
	tests := []struct {
		name     string
		args     args
		wantHash []byte
		wantErr  bool
	}{
		{
			name: "success",
			args: args{
				password: []byte("test1"),
				salt:     []byte("salt1"),
				r:        8,
				p:        1,
				n:        16,
			},
			wantHash: []byte{155, 14, 28, 179, 58, 64, 198, 176, 74, 48, 161, 168, 12, 134, 150, 179, 43, 33, 197, 179, 150, 167, 44, 227, 168, 25, 55, 244, 63, 24, 242, 107},
			wantErr:  false,
		},
		{
			name: "error",
			args: args{
				password: []byte("test1"),
				salt:     []byte("salt1"),
				r:        8,
				p:        1,
				n:        111,
			},
			wantHash: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotHash, err := RustScryptKey(tt.args.password, tt.args.salt, tt.args.r, tt.args.p, tt.args.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("RustScryptKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			require.Equal(t, tt.wantHash, gotHash)
		})
	}
}

func BenchmarkRustScryptKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = RustScryptKey([]byte("test1"), []byte("salt1"), 8, 1, 16)
	}
}

func BenchmarkGoScryptKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = GoScryptKey([]byte("test1"), []byte("salt1"), 8, 1, 16)
	}
}
