package main

// DO NOT USE ANY IMPORT

type QutexV1 struct {
	status string
}

func NewQutexV1() *QutexV1 {
	return &QutexV1{
		status: "unlocked",
	}
}

func (q *QutexV1) Lock() {
	for q.status == "locked" {
		continue
	}

	q.status = "locked"
}

func (q *QutexV1) Unlock() {
	if q.status == "unlocked" {
		panic("unlocking something unlocked")
	}

	q.status = "unlocked"
}

type QutexV2 struct {
	locked       bool
	unlcok_notif chan bool
}

func NewQutexV2() *QutexV2 {
	return &QutexV2{
		locked:       false,
		unlcok_notif: make(chan bool),
	}
}

func (q *QutexV2) Lock() {
	if q.locked {
		<-q.unlcok_notif
	}
	q.locked = true
}

func (q *QutexV2) Unlock() {
	if !q.locked {
		panic("")
	}
	q.locked = false
}
