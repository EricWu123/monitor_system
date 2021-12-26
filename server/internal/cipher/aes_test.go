package cipher

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAesEncrypt(t *testing.T) {
	type args struct {
		origData string
		key      []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"test-2", args{"11111111111111asdfaadsfasdffffsODFLkj", []byte("PHvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")}, "11111111111111asdfaadsfasdffffsODFLkj", false},
		{"test-1", args{"admin", []byte("PHvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")}, "admin", false},
		{"test-1", args{"guest", []byte("PHvL%$o0oNNoOZnk#o2qbqCeQB1iXeIR")}, "guest", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encData, err := AesEncrypt(tt.args.origData, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(string(encData))
			got, err := AesDecrypt(encData, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesEncrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesEncrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
