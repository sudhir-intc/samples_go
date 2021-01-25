package onhash256

// #cgo LDFLAGS: -lcrypto
// #include <stdio.h>
// #include <errno.h>
// #include <openssl/sha.h>
import "C"

import (
	"errors"
	"hash"
	"unsafe"
)

// The size of an SHA256 checksum in bytes.
const Size = 32

// The blocksize of SHA256 in bytes.
const BlockSize = 64

type digest struct {
	context *C.SHA256_CTX
}

// New returns a new hash.Hash computing the SHA1 checksum.
func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

func (d *digest) Reset() {
	d.context = &C.SHA256_CTX{}
	C.SHA256_Init(d.context)
}

func (d *digest) Write(p []byte) (nn int, err error) {
	if len(p) == 0 || C.SHA256_Update(d.context, unsafe.Pointer(&p[0]),
		C.size_t(len(p))) == 1 {
		return len(p), nil
	}

	return 0, errors.New("SHA256_Update failed")
}

func (d *digest) Sum(in []byte) []byte {
	context := *d.context
	defer func() { *d.context = context }()

	md := make([]byte, Size)
	if C.SHA256_Final((*C.uchar)(&md[0]), d.context) == 1 {
		return append(in, md...)
	}
	return nil
}

func (d *digest) Size() int { return Size }

func (d *digest) BlockSize() int { return BlockSize }
