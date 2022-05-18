package main

import (
	"crypto/md5"
	"encoding/binary"
	"flag"
	"fmt"
)

func mix_magic_strings(magic_strings_littlendian []uint32, in [16]byte, out *[64]byte) {

	var magic_strings []byte
	for _, n := range magic_strings_littlendian {
		bs := make([]byte, 4)
		binary.LittleEndian.PutUint32(bs, n)
		magic_strings = append(magic_strings, bs...)
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

func AP230(str_serial string, str_version string) string {
	platformSecret := [16]byte{'T', 'I', 'U', 't', '8', 'K', 'k', '5', 'A', '7', 'd', '4', 'W', 'i', 'r', 'H'}

	//"02301601202422"
	serial := []byte(str_serial)

	magic_strings_littlendian := []uint32{
		0x1cbe3df0,
		0xd034ed14,
		0x86487844,
		0x3dc2c339,
		0x427fdcb7,
		0xf06e8381,
		0xe82c9e44,
		0xeafcc130,
	}

	var thing [64]byte
	mix_magic_strings(magic_strings_littlendian, platformSecret, &thing)

	strlen := 0
	for ; strlen < len(thing) && thing[strlen] != 0; strlen++ {
	}

	md5digest2str(&platformSecret)

	platformSecret = md5.Sum(thing[:strlen])

	md5digest2str(&platformSecret)

	serial = append(serial, platformSecret[:]...)

	platformSecret = md5.Sum(serial)

	md5digest2str(&platformSecret)

	version := []byte(str_version)
	version = append(version, platformSecret[:]...)

	platformSecret = md5.Sum(version)

	md5digest2str(&platformSecret)

	return string(platformSecret[:])
}

func AP130(str_serial string) string {
	platformSecret := [16]byte{'J', 'P', 'E', 'i', 'X', '5', 'c', 'j', 's', 'b', 'c', 'R', 'T', 'P', '3', 'X'}

	//"02301601202422"
	serial := []byte(str_serial)

	magic_strings_littlendian := []uint32{
		0x58ad91d4,
		0x5d8d8176,
		0xe7ca7c76,
		0xb2c33e4a,
		0xc6cd6203,
		0x11b3895,
		0x1e581aed,
		0x67b10ed4,
	}

	var thing [64]byte
	mix_magic_strings(magic_strings_littlendian, platformSecret, &thing)

	strlen := 0
	for ; strlen < len(thing) && thing[strlen] != 0; strlen++ {
	}

	platformSecret = md5.Sum(thing[:strlen])
	md5digest2str(&platformSecret)

	//Add serial number
	serial = append(serial, platformSecret[:]...)
	platformSecret = md5.Sum(serial)

	md5digest2str(&platformSecret)

	return string(platformSecret[:])
}

func main() {

	strSerial := flag.String("serial", "", "Device serial number")
	strVersion := flag.String("version", "", "Firmware version (eg 10.3)")

	flag.Parse()

	if len(*strSerial) == 0 {
		fmt.Println("You need the devices serial number")
		return
	}

	password := ""
	switch (*strSerial)[1:4] {
	case "230":
		if len(*strVersion) == 0 {
			fmt.Println("Assuming version 10.0")
			*strVersion = "10.0"
		}

		password = AP230(*strSerial, *strVersion)
	case "130":
		password = AP130(*strSerial)
	default:
		fmt.Println("Unknown device type: ", (*strSerial)[1:4])
		return
	}

	fmt.Println(password)

}
