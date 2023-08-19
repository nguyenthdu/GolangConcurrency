package pubsub

import (
	"sync"
	"time"
)

type (
	//subcriber thuoc kieu channel
	subscriber chan interface{}
	// topic la mot filter
	topicFunc func(v interface{}) bool
)
type Publisher struct {
	//Read/Write Mutex
	m sync.Mutex
	// kich thuc hang doi
	buffer int
	//  timeout cho viec publishing
	timeout time.Duration
	// timeout da subcriber vao topic nao
	subcribers map[subscriber]topicFunc
}

// contructor voi timeout va doi dai hang doi
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:     buffer,
		timeout:    publishTimeout,
		subcribers: make(map[subscriber]topicFunc),
	}
}

// them subcriber moi, dang ky het tat ca topic
func (p *Publisher) Subscriber() chan interface{} {
	return p.SubscribeTopic(nil)
}

// them subscirber moi, subsdcribe cac topic da duoc filter loc
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subcribers[ch] = topic
	p.m.Unlock()
	return ch
}

// huy subscribe
func (p *Publisher) Evict(sub chan interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	delete(p.subcribers, sub)
	close(sub)
}

// pushlish ra 1 topic
func (p *Publisher) Publish(v interface{}) {
	p.m.Lock()
	defer p.m.Unlock()
	var wg sync.WaitGroup
	for sub, topic := range p.subcribers {
		wg.Add(1)
		go p.sendTopic(sub, topic, v, &wg)
	}
	wg.Wait()
}

// dong 1 doi tuongw publisher va dong tat ca cac subscriber
func (p *Publisher) Close() {
	p.m.Lock()
	defer p.m.Unlock()
	for sub := range p.subcribers {
		delete(p.subcribers, sub)
		close(sub)
	}
}

// gui 1 topic co the duy tri trong thoi gian cho wg
func (p *Publisher) sendTopic(sub subscriber, topic topicFunc, v interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	if topic != nil && !topic(v) {
		return
	}
	select {
	case sub <- v:
	case <-time.After(p.timeout):

	}
}
