package control_structures

func collatzChainLength(n int) int {
	var i int
	for i=0;n > 1;i++ {
		if n%2 == 0 {
			n = n/2
		} else {
			n = n*3 + 1
		}
	}
	return i;
}
