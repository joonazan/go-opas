package main

import (
	"encoding/json"
	"log"
	"path"
	"sync"
	"time"
)

const (
	EiTehty = iota
	Tehty
	Aloitettu
	TilojenMäärä
)

var (
	savefolder = "save"
	savefile   = path.Join(savefolder, "edistymiset.json")
	oldfile    = path.Join(savefolder, "vanhat_edistymiset.json")
)

type Edistymiset struct {
	sync.RWMutex
	data map[string]Edistyminen
}

func (e Edistymiset) Edistyminen(oppilas string) Edistyminen {
	if o, ok := e.get(oppilas); ok {
		return o
	}

	e.lisää(oppilas)
	o, _ := e.get(oppilas)
	return o
}

func (e Edistymiset) get(oppilas string) (ed Edistyminen, ok bool) {
	e.RLock()
	defer e.RUnlock()
	ed, ok = e.data[oppilas]
	return
}

func (e Edistymiset) lisää(oppilas string) {
	e.Lock()
	defer e.Unlock()
	e.data[oppilas] = teeEdistyminen()
}

func (e Edistymiset) TallennaVälein(kesto time.Duration) {
	ticker := time.NewTicker(kesto)
	go func() {
		for {
			<-ticker.C
			if err := e.Tallenna(); err == nil {
				log.Println("Automaattinen tallentaminen onnistui!")
			} else {
				log.Println("Virhe tallentaessa:", err)
			}
		}
	}()
}

func (e Edistymiset) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.data)
}

func (e Edistymiset) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &e.data)
}

type Edistyminen struct {
	sync.RWMutex
	data map[string]uint8
}

func (e Edistyminen) Set(aihe string, value uint8) {
	e.Lock()
	defer e.Unlock()
	e.data[aihe] = value
}

func (e Edistyminen) Get(aihe string) uint8 {
	e.RLock()
	defer e.RUnlock()
	return e.data[aihe]
}

func (e Edistyminen) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.data)
}

func (e *Edistyminen) UnmarshalJSON(data []byte) error {
	var with_strings map[string]json.Number
	err := json.Unmarshal(data, &with_strings)

	e.data = make(map[string]uint8, len(with_strings))
	for aihe, tila := range with_strings {
		num, err := tila.Int64()
		if err != nil {
			log.Println(err)
		}
		e.data[aihe] = uint8(num)
	}

	return err
}

func teeEdistyminen() Edistyminen {
	return Edistyminen{data: make(map[string]uint8)}
}
