package main

import (
	"reflect"
)

type Maailma struct {
	komponentit []*Komponentti
}

func (m *Maailma) LuoKomponentti(kuolleetVanhemmat *[]int, viipaleet ...interface{}) *Komponentti {
	k := &Komponentti{}
	k.Kierrättäjä = *UusiKierrättäjä(append(viipaleet, &k.Vanhemmat)...)
	k.kuolleetVanhemmat = kuolleetVanhemmat

	m.komponentit = append(m.komponentit, k)

	return k
}

func (m *Maailma) VapautaKuolleet() {
	for _, k := range m.komponentit {
		k.TapaOrvot()
	}

	for _, k := range m.komponentit {
		for _, kuollut := range k.Kuolleet {
			k.Vapauta(kuollut)
		}
		k.Kuolleet = k.Kuolleet[:0]
	}
}

type Komponentti struct {
	Kierrättäjä
	kuolleetVanhemmat   *[]int
	Vanhemmat, Kuolleet []int
}

func (k *Komponentti) Varaa(vanhempi int) (id int) {
	id = k.Kierrättäjä.Varaa()
	k.Vanhemmat[id] = vanhempi
	return
}

func (k *Komponentti) TapaOrvot() {
	for i, v := range k.Vanhemmat {
		for _, kuollut := range *k.kuolleetVanhemmat {
			if v == kuollut {
				k.Kuolleet = append(k.Kuolleet, i)
			}
		}
	}
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
