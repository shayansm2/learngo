package main

import (
	"errors"
	"gorm.io/gorm"
)

type TLDRDBCached struct {
	nonCachedProvider *TLDRProvider
	db                *gorm.DB
}

func (t *TLDRDBCached) Retrieve(key string) string {
	var entity *TLDREntity
	err := t.db.Where("Key = ?", key).First(&entity).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		entity = &TLDREntity{
			Key: key,
			Val: (*t.nonCachedProvider).Retrieve(key),
		}
		t.db.Create(entity)
	}
	return entity.Val
}

func (t *TLDRDBCached) List() []string {
	var cachedList []string
	t.db.Model(&TLDREntity{}).Select("Key").Scan(&cachedList)

	cachedKeys := make(map[string]bool)
	for _, key := range cachedList {
		cachedKeys[key] = true
	}

	nonCachedList := make([]string, 0)
	for _, key := range (*t.nonCachedProvider).List() {
		if _, found := cachedKeys[key]; !found {
			nonCachedList = append(nonCachedList, key)
		}
	}

	return append(cachedList, nonCachedList...)
}

func NewTLDRDBCached(nonCachedProvider TLDRProvider) TLDRProvider {
	return &TLDRDBCached{
		nonCachedProvider: &nonCachedProvider,
		db:                GetConnection(),
	}
}
