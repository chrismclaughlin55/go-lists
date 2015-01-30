package list

import (
"net/url"
"errors"
"sync"
"fmt"
)


type ArrayList struct {
	lock *sync.Mutex
	cap int
	len int
	list []interface{}

}
func (l *ArrayList) Init(){
l.lock = new(sync.Mutex)
fmt.Println("init")
}

func (l *ArrayList) Push(item interface {}){
	l.lock.Lock()
	defer func() { l.lock.Unlock()}()
	defer func() { l.len++}()
	l.list = append(l.list, item)

}

func (l *ArrayList) Pop() (interface{},error) {
	l.lock.Lock()
	defer func(){ l.lock.Unlock()}()
	if l.len == 0{
	return nil,errors.New("nothing in queue")
	}
	ret := l.list[0]
	l.list = l.list[1:]
	l.len--
	return ret, nil
}	

func (l *ArrayList) Print(){
for i := range l.list{
fmt.Printf("%d -> ", l.list[i])
}
fmt.Println()
}
func (l ArrayList) Length() int{
	return len(l.list)

}
