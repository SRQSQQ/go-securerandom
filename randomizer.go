package randomizer

import (
	crand "crypto/rand"
	"encoding/base64"
)

// GenerateRandomBytes is a function that takes an integer and returns
// a slice of that length containing random bytes.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)

	if _, err := crand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

// maximumBytes is used to calculate the how many bytes we can generate a
// base64 string from, if we don't want the string to be longer than size.
func maximumBytes(size int) int {
	return int((float64(size) / 4) * 3)
}

// GenerateRandomBase64 is a function that returns a randomize standard Base64
// string. This does not use the URL encoding. It takes a single parameter that
// is the maximum possible length of the string, it will get as close as possible.
func GenerateRandomBase64(max int) (string, error) {
	b, err := GenerateRandomBytes(maximumBytes(max))
	return base64.StdEncoding.EncodeToString(b), err
}

// GenerateRandomBase64URLEncoding is a function that returns a randomize standard Base64
// string. This does not use the URL encoding. It takes a single parameter that
// is the maximum possible length of the string, it will get as close as possible.
func GenerateRandomBase64URLEncoding(max int) (string, error) {
	b, err := GenerateRandomBytes(maximumBytes(max))
	return base64.URLEncoding.EncodeToString(b), err
}

// GenerateRandomUint16 is a function that returns a uint16 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func GenerateRandomUint16() (uint16, error) {
	b, err := GenerateRandomBytes(2)

	if err != nil {
		return 0, err
	}

	var u16 uint16

	for i := range b {
		offset := uint16(i) + 1
		shift := 16 - (8 * offset)
		u16 = u16 | uint16(b[i])<<shift
	}

	return u16, nil
}

// GenerateRandomUint32 is a function that returns a uint32 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func GenerateRandomUint32() (uint32, error) {
	b, err := GenerateRandomBytes(4)

	if err != nil {
		return 0, err
	}

	var u32 uint32

	for i := range b {
		offset := uint32(i) + 1
		shift := 32 - (8 * offset)
		u32 = u32 | uint32(b[i])<<shift
	}

	return u32, nil
}

// GenerateRandomUint64 is a function that returns a uint64 generated by
// 'bitwise-or'ing 4 bytes from crypto/rand.Read().
func GenerateRandomUint64() (uint64, error) {
	b, err := GenerateRandomBytes(8)

	if err != nil {
		return 0, err
	}

	var u64 uint64

	for i := range b {
		offset := uint64(i) + 1
		shift := 64 - (8 * offset)
		u64 = u64 | uint64(b[i])<<shift
	}

	return u64, nil
}