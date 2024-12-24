package m_to

type Ts struct{}

func New() *Ts {
	return &Ts{}
}
func (t *Ts) Bool(v interface{}) bool {
	return Bool(v)
}

func (t *Ts) String(v interface{}) string {
	return String(v)
}

func (t *Ts) Int(v interface{}) int {
	return Int(v)
}

func (t *Ts) Int64(v interface{}) int64 {
	return Int64(v)
}

func (t *Ts) Float64(v interface{}) float64 {
	return Float64(v)
}

func (t *Ts) Uint(v interface{}) uint {
	return uint(t.Uint64(v))
}

func (t *Ts) Uint64(v interface{}) uint64 {
	return Uint64(v)
}

func (t *Ts) Uint32(v interface{}) uint32 {
	return Uint32(v)
}

func (t *Ts) Float32(v interface{}) float32 {
	return Float32(v)
}

func (t *Ts) Json(v interface{}) string {
	return Json(v)
}
