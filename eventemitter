package eventemitter

// EventEmitter struct: Hanya berupa deretan index yang berisi chan map
type EventEmitter struct {
	// Muatan untuk setiap chan buffer
	muatan int
	// Muatan Event On
	mapEvents map[chan interface{}][]string
	// Muatan Event Emit
	mapClients map[string][]chan interface{}
}

// New : membuat EventEmitter dengan data init yang diharapkan.
func New(initChanCapacity int) *EventEmitter {
	initmapEvents := make(map[chan interface{}][]string)
	initmapClients := make(map[string][]chan interface{})

	EventEmit := EventEmitter{
		mapEvents:  initmapEvents,
		mapClients: initmapClients,
	}
	EventEmit.muatan = initChanCapacity
	return &EventEmit
}

// On : Indormasi dari daftar Chan yang dikeluarkan
func (p *EventEmitter) On(namaEvents ...string) chan interface{} {
	//init new chan using muatan as channel buffer
	workChan := make(chan interface{}, p.muatan)

	var updateChanList []chan interface{}
	for _, namaEvent := range namaEvents {
		updateChanList, _ = p.mapClients[namaEvent]
		updateChanList = append(updateChanList, workChan)
		p.mapClients[namaEvent] = updateChanList
	}
	p.mapEvents[workChan] = namaEvents

	return workChan
}

// Emit : Emit berupa iformasi yang dikirim ke daftar channel
func (p *EventEmitter) Emit(content interface{}, namaEvents ...string) {
	for _, namaEvent := range namaEvents {
		if chanList, ok := p.mapClients[namaEvent]; ok {
			//Someone has subscribed this namaEvent
			for _, channel := range chanList {
				channel <- content
			}
		}
	}
}

// EventPlus :  menambahkan nama event kedalam daftar chan yang aktif
func (p *EventEmitter) EventPlus(clientChan chan interface{}, namaEvents ...string) {

	var updateChanList []chan interface{}
	for _, namaEvent := range namaEvents {
		updateChanList, _ = p.mapClients[namaEvent]
		updateChanList = append(updateChanList, clientChan)
		p.mapClients[namaEvent] = updateChanList
	}
	p.mapEvents[clientChan] = namaEvents

}

// EventMin : menghapus atau membuang nama event dari daftar chan yang aktif
func (p *EventEmitter) EventMin(clientChan chan interface{}, namaEvents ...string) {

	for _, namaEvent := range namaEvents {
		// memilah event yang harus dibuang
		if chanList, ok := p.mapClients[namaEvent]; ok {
			// membuang salah satu event dari daftar
			var updateChanList []chan interface{}
			for _, client := range chanList {
				if client != clientChan {
					updateChanList = append(updateChanList, client)
				}
			}
			p.mapClients[namaEvent] = updateChanList
		}
		if ListnamaEvent, ok := p.mapEvents[clientChan]; ok {
			var updateTopicList []string
			for _, updateTopic := range ListnamaEvent {
				if updateTopic != namaEvent {
					updateTopicList = append(updateTopicList, namaEvent)
				}
			}
			p.mapEvents[clientChan] = updateTopicList
		}
	}
}
