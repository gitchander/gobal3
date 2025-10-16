package bal3

// import (
// 	"fmt"
// )

// // sum of 4 trits: [-4..4]

// func splitTrits1(v int) (hi, lo Trit) {
// 	const (
// 		v0 = 0
// 	)
// 	switch v {
// 	case -4:
// 		return -1, -1
// 	case -3:
// 		return -1, v0
// 	case -2:
// 		return -1, +1
// 	case -1:
// 		return v0, -1
// 	case 0:
// 		return v0, v0
// 	case +1:
// 		return v0, +1
// 	case +2:
// 		return +1, -1
// 	case +3:
// 		return +1, v0
// 	case +4:
// 		return +1, +1
// 	default:
// 		panic(fmt.Errorf("splitTrits1: invalid value %d", v))
// 	}
// }

// func splitTrits2(v int) (hi, lo Trit) {
// 	const (
// 		vN = -1
// 		vZ = 0
// 		vP = +1
// 	)
// 	switch v {
// 	case -4:
// 		return vN, vN
// 	case -3:
// 		return vN, vZ
// 	case -2:
// 		return vN, vP
// 	case -1:
// 		return vZ, vN
// 	case 0:
// 		return vZ, vZ
// 	case +1:
// 		return vZ, vP
// 	case +2:
// 		return vP, vN
// 	case +3:
// 		return vP, vZ
// 	case +4:
// 		return vP, vP
// 	default:
// 		panic(fmt.Errorf("splitTrits2: invalid value %d", v))
// 	}
// }
