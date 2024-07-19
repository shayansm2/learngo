package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TLDRSample struct {
	data map[string]string
}

func (t *TLDRSample) Retrieve(key string) string {
	return t.data[key]
}

func (t *TLDRSample) List() []string {
	res := make([]string, 0, len(t.data))
	for k := range t.data {
		res = append(res, k)
	}
	return res
}

func NewTLDRSample() TLDRProvider {
	return &TLDRSample{
		data: map[string]string{
			"ls":   "ls is good",
			"bash": "bash os also good",
			"lsd":  "even better than ls",
			"zsh":  "a posix complaint shell with cool features",
			"go":   "the famous go compiler from Google inc",
		},
	}
}

func BeforeEachSample() {
	GetConnection().Migrator().DropTable(&TLDREntity{})
	GetConnection().AutoMigrate(&TLDREntity{})
}

func TestRetrieveSample(t *testing.T) {
	BeforeEachSample()
	imd := NewTLDRSample()
	cached := NewTLDRDBCached(imd)
	assert.Equal(t, imd.Retrieve("bash"), cached.Retrieve("bash")) // slow
	assert.Equal(t, imd.Retrieve("bash"), cached.Retrieve("bash")) // fast(cached)
}

func TestListSample(t *testing.T) {
	BeforeEachSample()
	imd := NewTLDRSample()
	cached := NewTLDRDBCached(imd)
	assert.ElementsMatch(t, imd.List(), cached.List())
}

func TestFunctionality(t *testing.T) {
	BeforeEachSample()
	imd := NewTLDRSample()
	cached := NewTLDRDBCached(imd)
	assert.Equal(t, imd.Retrieve("bash"), cached.Retrieve("bash")) // slow
	assert.Equal(t, imd.Retrieve("bash"), cached.Retrieve("bash")) // fast(cached)
	assert.Equal(t, imd.Retrieve("bash"), cached.Retrieve("bash")) // fast(cached)
	var rowCount int64
	GetConnection().Model(TLDREntity{}).Count(&rowCount)
	assert.Equal(t, rowCount, int64(1))

	assert.Equal(t, imd.Retrieve("go"), cached.Retrieve("go")) // slow
	assert.Equal(t, imd.Retrieve("go"), cached.Retrieve("go")) // fast
	assert.Equal(t, imd.Retrieve("go"), cached.Retrieve("go")) // fast
	GetConnection().Model(TLDREntity{}).Count(&rowCount)
	assert.Equal(t, rowCount, int64(2))

	ls := cached.List()
	fmt.Println(ls)
	assert.True(t, ls[0] == "go" || ls[1] == "go")
	assert.True(t, ls[0] == "bash" || ls[1] == "bash")
	assert.Equal(t, len(imd.List()), len(ls))
}
