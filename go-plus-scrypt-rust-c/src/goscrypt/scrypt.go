package goscrypt

/*
#cgo LDFLAGS: -lrustscryptbindings
#include "rustscryptbindings.h"
*/
import "C"
import (
	"errors"
	"golang.org/x/crypto/scrypt"
	"unsafe"
)

func RustScryptKey(password, salt []byte, r, p, n int) (result []byte, err error) {
	result = make([]byte, 32)

	ret := C.scrypt_key(
		C.CString(string(password)),
		C.CString(string(salt)),
		C.int(r),
		C.int(p),
		C.int(n),
		(*C.uchar)(unsafe.Pointer(&result[0])),
		C.size_t(32),
	)
	if ret != 0 {
		return nil, errors.New("some error")
	}

	return result, nil
}

func GoScryptKey(password, salt []byte, r, p, n int) ([]byte, error) {
	return scrypt.Key(password, salt, n, r, p, 32)
}
