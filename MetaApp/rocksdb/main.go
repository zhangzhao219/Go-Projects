package main

import (
	"fmt"
	"sync/atomic"

	"github.com/linxGnu/grocksdb"
)

func main() {
	bbto := grocksdb.NewDefaultBlockBasedTableOptions() // SST文件的默认格式
	bbto.SetBlockCache(grocksdb.NewLRUCache(3 << 30))   // 新建LRU缓存

	// 添加一个布隆过滤器
	filter := grocksdb.NewBloomFilter(10)
	bbto.SetFilterPolicy(filter)

	// 根据配置创建数据库
	opts := grocksdb.NewDefaultOptions()
	opts.SetBlockBasedTableFactory(bbto)
	opts.SetCreateIfMissing(true) // 如果文件夹不存在则创建

	db, err := grocksdb.OpenDb(opts, "rocksdb-demo") // 打开数据库（实际上就是一个文件夹）
	if err != nil {
		fmt.Printf("RocksDB open failed, error=%v\n", err)
		return
	}

	// 5. 批量写入
	wo := grocksdb.NewDefaultWriteOptions()
	defer wo.Destroy()

	wb := grocksdb.NewWriteBatch()
	defer wb.Destroy()

	wb.Put([]byte("foo"), []byte("bar"))
	wb.Put([]byte("bar"), []byte("foo"))

	err = db.Write(wo, wb) // 真正写入
	if err != nil {
		fmt.Printf("RocksDB batch write failed, error=%v\n", err)
		return
	}

	// 4. 批量读取
	ro := grocksdb.NewDefaultReadOptions()
	defer ro.Destroy()

	ro.SetFillCache(false) // 批量读取的时候要关闭

	// 创建迭代器
	itdb := db.NewIterator(ro)
	defer itdb.Close()

	itdb.Seek([]byte("bar")) // 为有序存储，移动迭代器到指定的位置，然后再开始迭代
	for it := itdb; it.Valid(); it.Next() {
		key := it.Key()
		value := it.Value()
		fmt.Printf("Key: %v Value: %v\n", string(key.Data()), string(value.Data()))
		key.Free()
		value.Free()
	}
	if err := itdb.Err(); err != nil {
		fmt.Printf("RocksDB iterator failed, error=%v\n", err)
	}
	var a atomic.Value
	a.Store(false)
	if b := a.Load(); b != nil && b.(bool) {
		fmt.Println(b)
	}

}
