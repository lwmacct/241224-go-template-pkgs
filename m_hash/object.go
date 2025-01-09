package m_hash

type Ts struct{}

func (t *Ts) New() *Ts {
	return &Ts{}
}

// MD5 计算字符串的 MD5 哈希值并返回十六进制编码的字符串。
func (t *Ts) MD5(s string) string {
	return MD5(s)
}

// SHA1计算字符串的 SHA-1 哈希值并返回十六进制编码的字符串。
func (t *Ts) SHA1(s string) string {
	return SHA1(s)
}

// SHA3 计算字符串的 SHA3-256 哈希值并返回十六进制编码的字符串。
func (t *Ts) SHA3(s string) string {
	return SHA3(s)
}

// SHA256 计算字符串的 SHA-256 哈希值并返回十六进制编码的字符串。
func (t *Ts) SHA256(s string) string {
	return SHA256(s)
}

// SHA512 计算字符串的 SHA-512 哈希值并返回十六进制编码的字符串。
func (t *Ts) SHA512(s string) string {
	return SHA512(s)
}

// FNV 计算字符串的 FNV-1a 哈希值并返回 uint64 类型的哈希值。
func (t *Ts) FNV(s string) uint64 {
	return FNV(s)
}

// CRC32 计算字符串的 CRC32 哈希值并返回 uint32 类型的哈希值。
func (t *Ts) CRC32(s string) uint32 {
	return CRC32(s)
}

// Murmur 计算字符串的 MurmurHash3 哈希值并返回 uint32 类型的哈希值。
func (t *Ts) Murmur(s string) uint32 {
	return Murmur(s)
}

// Jenkins 计算字符串的 Jenkins Hash 哈希值并返回 uint32 类型的哈希值。
// 由于 Go 中没有官方的 Jenkins Hash 实现，这里提供一个简单的实现。
func (t *Ts) Jenkins(s string) uint32 {
	return Jenkins(s)
}

// XXHash64 计算字符串的 XXHash64 哈希值并返回 uint64 类型的哈希值。
func (t *Ts) XXHash64(s string) uint64 {
	return XXHash64(s)
}
