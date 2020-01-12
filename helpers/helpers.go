package helpers

func IDGenerator(start int32) func() int32 {
	i := start - 1
	return func() int32 {
		i++
		return i
	}
}
