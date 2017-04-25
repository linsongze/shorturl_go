package service

import "sync/atomic"

type RAMStore struct {
	 store map[string]string
	 id *int64
}
func NewRamStore()*RAMStore{
	rs := new(RAMStore)
	rs.id = new(int64)
	rs.store=make(map[string]string)
	return rs
}

func (rs *RAMStore)Save(shortCode,url string)  {
	rs.store[shortCode]=url
}
func (rs *RAMStore)IncAndGet()int64{
	newId := atomic.AddInt64(rs.id,1)
	return newId
}
func(rs *RAMStore)Get(shortCode string)string{
	return rs.store[shortCode]
}

