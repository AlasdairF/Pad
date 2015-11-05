package pad

func Bytes(input []byte, to int, padder byte) []byte {

	l := len(input)
	if l >= to {
		return input
	}
	
	b := make([]byte, to)
	dif := to - l
	var i int
	for i=0; i<dif; i++ {
		b[i] = padder
	}
	copy(b[i:], input)

	return b
}

func String(input string, to int, padder byte) string {

	l := len(input)
	if l >= to {
		return input
	}
	
	b := make([]byte, to)
	dif := to - l
	var i int
	for i=0; i<dif; i++ {
		b[i] = padder
	}
	copy(b[i:], input)

	return string(b)
}

func StringBytes(input string, to int, padder byte) []byte {

	l := len(input)
	if l >= to {
		return input
	}
	
	b := make([]byte, to)
	dif := to - l
	var i int
	for i=0; i<dif; i++ {
		b[i] = padder
	}
	copy(b[i:], input)

	return b
}
