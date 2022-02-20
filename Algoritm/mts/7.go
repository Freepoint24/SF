package main

import "fmt"

func main() {
	value := "АаааИИИИ"
	enc := 
	buf := Counter2(enc)

fmt.Printf("%s\n%s\n", value, enc)
}


func Counter2(buf []byte) []byte {
	l := len(buf)
	nl := l >> 1
	if nl < 16 {
		nl = 16
	}

	n := make([]byte, 0, nl)

	i := 0
	for {
		if i >= l {
			break
		}
		cnt := 1
		for j := i + 1; j < l; j++ {
			if buf[i] == buf[j] {
				cnt++
				if cnt == 257 {
					break
				}
			} else {
				break
			}
		}

		switch cnt {
		case 0:
			// should not be possible
			panic("case 0 triggered...")
		case 1:
			n = append(n, buf[i])
		case 2:
			n = append(n, buf[i])
			n = append(n, buf[i])
		default: // > 3
			n = append(n, buf[i])
			n = append(n, buf[i])
			n = append(n, buf[i])
			n = append(n, byte(cnt - 3))
		}

		i += cnt
	}

	return n
}
