package kuncore

type channelCountStruct struct {
	key   string
	reply chan int64
}

type ChannelCC struct {
	ic chan string
	dc chan string
	lc chan channelCountStruct
	m  map[string]int64
}

func NewChannelCC() *ChannelCC {
	cc := &ChannelCC{
		ic: make(chan string, 100000),
		dc: make(chan string, 100000),
		lc: make(chan channelCountStruct),
		m:  make(map[string]int64),
	}
	go cc.loop()
	return cc
}

func (cc *ChannelCC) loop() {
	for {
		select {
		case ci := <-cc.ic:
			cc.m[ci]++
			break
		case cd := <-cc.dc:
			cc.m[cd]++
			break
		case cl := <-cc.lc:
			cl.reply <- cc.m[cl.key]
			break
		}
	}
}

func (cc *ChannelCC) Increase(key string) {
	cc.ic <- key
}

func (cc *ChannelCC) Decrease(key string) {
	cc.dc <- key
}

func (cc *ChannelCC) Count(key string) int64 {
	cs := &channelCountStruct{
		key:   key,
		reply: make(chan int64),
	}
	return <-cs.reply
}
