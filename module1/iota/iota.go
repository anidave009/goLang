package iota

const (
	KB = 1 << (10 * (iota + 1))
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func PrintIotaConstants() {
	println("KB:", KB)
	println("MB:", MB)
	println("GB:", GB)
	println("TB:", TB)
	println("PB:", PB)
	println("EB:", EB)
	println("ZB:", float64(ZB))
	println("YB:", float64(YB))
}
