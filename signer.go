package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
)

var mu sync.Mutex // мютекс во избежание перегрева функции DataSignerMd5

func SingleHash(in, out chan interface{}) {
	wg := new(sync.WaitGroup)
	for dataInterf := range in {
		wg.Add(1)
		innerWrapper := func(data string) {
			outMD5, outCRC := make(chan interface{}), make(chan interface{})
			go func() {
				mu.Lock()
				resMD5 := DataSignerMd5(data)
				mu.Unlock()
				outMD5 <- DataSignerCrc32(resMD5)
			}()
			go func() {
				outCRC <- DataSignerCrc32(data)
			}()
			res := (<-outCRC).(string) + "~" + (<-outMD5).(string)
			out <- res
			wg.Done()
		}
		go innerWrapper(fmt.Sprintf("%v", dataInterf))
	}
	wg.Wait() // ждем завершения всех внутренних горутин, чтобы закрыть выходной канал
}

func MultiHash(in, out chan interface{}) {
	wg := new(sync.WaitGroup)
	for dataInterf := range in {
		wg.Add(1)
		innerWrapper := func(data string) {
			var outChans []chan interface{}
			var hashes []string
			crcWrapper := func(in, out chan interface{}) {
				out <- DataSignerCrc32((<-in).(string))
			}
			for th := 0; th <= 5; th++ {
				inInner, outInner := make(chan interface{}), make(chan interface{})
				go crcWrapper(inInner, outInner)
				ths := strconv.Itoa(th)
				inInner <- ths + data
				outChans = append(outChans, outInner)
			}
			for i := 0; i <= 5; i++ {
				hashes = append(hashes, (<-outChans[i]).(string))
			}
			out <- strings.Join(hashes, "")
			wg.Done()
		}
		go innerWrapper(fmt.Sprintf("%v", dataInterf))
	}
	wg.Wait() // ждем завершения всех внутренних горутин, чтобы закрыть выходной канал
}

func CombineResults(in, out chan interface{}) {
	var res []string
	for data := range in {
		res = append(res, data.(string))
	}
	sort.Strings(res)
	out <- strings.Join(res, "_")
}

func ExecutePipeline(jobs ...job) {
	channels := make([]chan interface{}, len(jobs)+1)
	jobWrapper := func(j job, in, out chan interface{}) {
		// закрываем выходной канал каждой горутины
		// закрытие внешнее, поскольку в готовом тесте закрытия канала первой конвейерной функции не происходит
		defer close(out)
		j(in, out)
	}
	for i := range channels {
		channels[i] = make(chan interface{})
	}
	for i := 0; i < len(jobs)-1; i++ {
		go jobWrapper(jobs[i], channels[i], channels[i+1])
	}
	ind := len(jobs)-1
	// запускаем последний job в основной горутине, не закрывая выходной канал
	jobs[ind](channels[ind], channels[ind+1])
}
