package main

import (
	"fmt"
)

// 策略模式 :只将变化的部分抽象成接口 简而言之，策略模式就是定义一系列的算法，将每个算法封装起来，并使它们可以互换。
// 和代理模式很像
// 1. 编写一个抽象策略 StorageStrategy 存储策略
type StorageStrategy interface {
	Save()
}

var strategys = map[string]StorageStrategy{
	"file":         &fileStorage{},
	"encrypt_file": &encryptFileStorage{},
}

// NewStorageStrategy
func NewStorageStrategy(t string) (StorageStrategy, error) {
	s, ok := strategys[t]
	if !ok {
		return nil, fmt.Errorf("not found StorageStrategy: %s", t)
	}

	return s, nil
}

// 2、编写两个具体策略，实现抽象策略中的抽象方法
// FileStorage 保存到文件
type fileStorage struct{}

// Save Save
func (s *fileStorage) Save() {
	fmt.Println("save to file")
}

// encryptFileStorage 加密保存到文件
type encryptFileStorage struct{}

// encrypt Save
func (s *encryptFileStorage) Save() {
	fmt.Println("encrypt save to file")
}
