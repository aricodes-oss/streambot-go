package channels

import (
	"streambot/events"
)

var EventDispatch = make(chan *events.Event, 8*1024)
