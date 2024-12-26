package m_closure

func NewCounter(start int) func() int {
	val := start - 1
	return func() int {
		val++
		return val
	}
}
