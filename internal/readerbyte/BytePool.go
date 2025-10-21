package readerbyte

import (
	"bytes"
	"io"
	"sync"
)

var BufferPool = sync.Pool{
	New: func() interface{} {
		//32 KB 缓冲（按需调整）
		return make([]byte, 32*1024)
	},
}

func ReadBodyWithFixedBuf(r io.Reader) ([]byte, error) {
	buf := BufferPool.Get().([]byte)
	defer BufferPool.Put(buf)

	// 使用 bytes.Buffer + io.CopyBuffer 以确保读完整个 body（适用于不确定长度）
	var b bytes.Buffer
	// io.CopyBuffer 会重复使用 buf 作为临时缓冲
	if _, err := io.CopyBuffer(&b, r, buf); err != nil {
		return nil, err
	}

	// 返回一个独立的拷贝，如果你要长期保存结果，最好复制一份
	// 如果只是临时序列化解析，可以直接用 b.Bytes()（注意不要放回 pooled buf 之前逃逸）
	return b.Bytes(), nil
}
