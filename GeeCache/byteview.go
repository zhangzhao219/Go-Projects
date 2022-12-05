package geecache

// 只读数据结构用来表示缓存值
type ByteView struct {
	b []byte
}

// 返回缓存值的长度
func (v ByteView) Len() int {
	return len(v.b)
}

// 返回拷贝从而防止这个值被外部操作修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// 将数据作为一个字符串进行返回
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
