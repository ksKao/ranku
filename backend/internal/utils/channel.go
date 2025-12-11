package utils

import "log"

// the data in the channel doesn't matter, we just want to detect whether an event is sent so that the SSE endpoint can retrieve the data
var channels []*chan int = []*chan int{}

func AddChannel(ch *chan int) {
	channels = append(channels, ch)
}

func RemoveChannel(ch *chan int) {
	pos := -1
	storeLen := len(channels)

	for i, channel := range channels {
		if ch == channel {
			pos = i
		}
	}

	if pos == -1 {
		return
	}

	channels[pos] = channels[storeLen-1]
	channels = channels[:storeLen-1]
	log.Printf("Connection remains: %d", len(channels))
}

func BroadcastUpdate() {
	for _, ch := range channels {
		*ch <- 0
	}
}
