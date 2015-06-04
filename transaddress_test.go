// Copyright 2015 Factom Foundation
// Use of this source code is governed by the MIT
// license that can be found in the LICENSE file.

package simplecoin

import (
	"fmt"
	"github.com/agl/ed25519"
	"math/rand"
	"testing"
)

var _ = fmt.Printf
var _ = ed25519.Sign
var _ = rand.New

// An address
var taddress1 = [ADDRESS_LENGTH]byte{
	0x61, 0xe3, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72, 0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
	0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c, 0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
}

// An address
var taddress2 = [ADDRESS_LENGTH]byte{
	0x61, 0xe4 /* <= */, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72, 0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
	0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c, 0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00,
}

func Test_TAddressEquals(test *testing.T) {
	a1 := new(TransAddress)
	a2 := new(TransAddress)

	a1.amount = 5
	a2.amount = 5

	a1.address = new(Address)
	a2.address = new(Address)

	a1.address.SetBytes(address1[:])
	a2.address.SetBytes(address1[:])

	if !a1.IsEqual(a2) { // Out of the box, hashes should be equal
		PrtStk()
		test.Fail()
	}

	a1.address.SetBytes(address2[:])

	if a1.IsEqual(a2) { // Now they should not be equal
		PrtStk()
		test.Fail()
	}

	a2.address.SetBytes(address2[:])

	if !a1.IsEqual(a2) { // Back to equality!
		PrtStk()
		test.Fail()
	}

	a1.amount = 6

	if a1.IsEqual(a2) { // Amounts are not equal
		PrtStk()
		test.Fail()
	}

}