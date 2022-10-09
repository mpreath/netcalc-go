package utils

import (
	"fmt"
)

func GetNetworkAddress(address uint32, mask uint32) uint32 {
	return address & mask
}

func GetBroadcastAddress(address uint32, mask uint32) uint32 {
	return address | (^mask)
}

func IsValidMask(mask uint32) bool {
	for i := 1; i <= 32; i++ {
		calcMask, _ := GetMaskFromBits(i)
		if mask == calcMask {
			return true
		}
	}

	return false
}

func GetBitsInMask(mask uint32) int {
	bc := 0
	for mask != 0 {
		mask = mask << 1
		bc++
	}
	return bc
}

func GetMaskFromBits(bits int) (uint32, error) {
	if bits <= 32 {
		var mask uint32 = 0
		mask = ^mask
		bc := 32 - bits
		mask = mask << bc
		return mask, nil
	} else {
		return 0, fmt.Errorf("utils:GetMaskFromBits: bits must be 32 or less")
	}

}

func GetCommonBitMask(n1 uint32, n2 uint32) uint32 {
	commonBits := n1 ^ n2

	idx := 0

	for commonBits != 0 {
		commonBits = commonBits >> 1
		idx++
	}

	newMask, _ := GetMaskFromBits(32 - idx)

	return newMask
}
