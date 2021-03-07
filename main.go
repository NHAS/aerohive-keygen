package main

import (
	"crypto/md5"
	"flag"
	"fmt"
)

func mix_magic_strings(in [16]byte, out *[64]byte) {

	magic_strings := [32]byte{
		0xf0, 0x3d, 0xbe, 0x1c,
		0x14, 0xed, 0x34, 0xd0,
		0x44, 0x78, 0x48, 0x86,
		0x39, 0xc3, 0xc2, 0x3d,
		0xb7, 0xdc, 0x7f, 0x42,
		0x81, 0x83, 0x6e, 0xf0,
		0x44, 0x9e, 0x2c, 0xe8,
		0x30, 0xc1, 0xfc, 0xea,
	}

	swapLocations := [8]int{
		3, 1, 4, 5, 0, 7, 2, 6,
	}

	for i := 0; i < len(in); i++ {
		out[i] = in[i]
	}

	for i := 0; i < len(magic_strings); i++ {
		out[len(in)+i] = magic_strings[i]
	}

	inputLen := len(in)

	for i := range swapLocations {
		cVar1 := out[inputLen+swapLocations[i]]
		out[inputLen+swapLocations[i]] = out[i]
		out[i] = cVar1
	}
}

func md5digest2str(input *[16]byte) {

	if len(input) > 0 {

		for i := range input {
			uVar5 := input[i] % 59
			bVar1 := byte(uVar5)

			input[i] = bVar1

			bVar4 := bVar1 + 48

			if uVar5 < 9 {
				input[i] = bVar4
			} else {
				bVar4 = bVar1 + 56
				if uVar5 < 0x22 {
					input[i] = bVar4
					continue
				}
				input[i] = bVar1 + 63
			}

		}
	}
}

func main() {

	strSerial := flag.String("serial", "", "Device serial number")

	flag.Parse()

	if len(*strSerial) == 0 {
		fmt.Println("You need the devices serial number")
		return
	}

	platformSecret := [16]byte{'T', 'I', 'U', 't', '8', 'K', 'k', '5', 'A', '7', 'd', '4', 'W', 'i', 'r', 'H'}

	//"02301601202422"
	serial := []byte(*strSerial)

	var thing [64]byte
	mix_magic_strings(platformSecret, &thing)

	strlen := 0
	for ; strlen < len(thing) && thing[strlen] != 0; strlen++ {
	}

	md5digest2str(&platformSecret)

	platformSecret = md5.Sum(thing[:strlen])

	md5digest2str(&platformSecret)

	serial = append(serial, platformSecret[:]...)

	platformSecret = md5.Sum(serial)

	md5digest2str(&platformSecret)

	version := []byte("10.0")
	version = append(version, platformSecret[:]...)

	platformSecret = md5.Sum(version)

	md5digest2str(&platformSecret)

	fmt.Println("Shell secret key: ", string(platformSecret[:]))
}
