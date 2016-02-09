package main

import "reflect"

func PoistaKuolleet(huoltajat, kuolleet []int, kierrättäjä Kierrättäjä) {
	for i, h := range huoltajat {
		for _, kuollut := range kuolleet {
			if h == kuollut {
				kierrättäjä.Vapauta(i)
			}
		}
	}
}

type Komponentti struct {
	Kierrättäjä
	Huoltajat []int
}

func UusiKomponentti(viipaleet ...interface{}) *Komponentti {
	k := new(Komponentti)
	k.Kierrättäjä = *UusiKierrättäjä(append(viipaleet, &k.Huoltajat)...)
	return k
}

func (k *Komponentti) Varaa(huoltaja int) (id int) {
	id = k.Kierrättäjä.Varaa()
	k.Huoltajat[id] = huoltaja
	return
}

func (k *Komponentti) PoistaKuolleet(kuolleetHuoltajat []int) {
	PoistaKuolleet(k.Huoltajat, kuolleetHuoltajat, k.Kierrättäjä)
}

type Kierrättäjä struct {
	viipaleet  []reflect.Value
	nollaarvot []reflect.Value
	vapaat     []int
}

func UusiKierrättäjä(viipaleet ...interface{}) *Kierrättäjä {
	k := new(Kierrättäjä)
	k.viipaleet = make([]reflect.Value, len(viipaleet))
	k.nollaarvot = make([]reflect.Value, len(viipaleet))

	for i, v := range viipaleet {
		viipale := reflect.ValueOf(v).Elem()
		if viipale.Kind() != reflect.Slice {
			panic("Kierrättäjälle yritettiin antaa jokin joka ei ole viipaleen osoite.")
		}
		k.viipaleet[i] = viipale
		k.nollaarvot[i] = reflect.Zero(viipale.Type().Elem())
	}
	return k
}

func (k *Kierrättäjä) Varaa() (id int) {
	if len(k.vapaat) == 0 {
		id = k.varaa()
	} else {
		viimeinen := len(k.vapaat) - 1
		id = k.vapaat[viimeinen]
		k.vapaat = k.vapaat[:viimeinen]
	}
	return
}

func (k *Kierrättäjä) Vapauta(id int) {
	k.vapaat = append(k.vapaat, id)
	for i, v := range k.viipaleet {
		v.Index(id).Set(k.nollaarvot[i])
	}
}

func (k *Kierrättäjä) varaa() (vanhaPituus int) {
	vanhaPituus = k.viipaleet[0].Len()

	for i, v := range k.viipaleet {
		v.Set(reflect.Append(v, k.nollaarvot[i]))
	}
	return
}
