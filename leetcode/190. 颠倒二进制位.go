package main

func reverseBits(num uint32) uint32 {
	var r uint32
	for i := 0; i < 32; i++ {
		r = (r << 1) | (num & 1)
		num >>= 1
	}
	return r
}
