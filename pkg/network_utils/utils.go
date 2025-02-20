package networkutils

func ToInteger(s string) int {
	if len(s) == 0 {
		return 0
	}
	zero := int('0')
	value := int(s[0])

	return value - zero
}
