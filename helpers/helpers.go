package helpers

func IDGenerator() func() int32 {
	i := 0
	return func() int32 {
		i++
		return int32(i)
	}
}
