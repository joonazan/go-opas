// +build !heroku

package main

import "os"

func (e Edistymiset) Tallenna() error {
	e.RLock()
	data, err := json.Marshal(e)
	if err != nil {
		return err
	}
	e.RUnlock()

	err = os.MkdirAll(savefolder, os.ModePerm)
	if err != nil {
		return err
	}
	err = os.Rename(savefile, oldfile)
	if err != nil && !os.IsNotExist(err) {
		return err
	}
	err = ioutil.WriteFile(savefile, data, os.ModePerm)
	if err != nil {
		os.Rename(oldfile, savefile)
		return err
	}

	return nil
}

func lataaEdistymiset() Edistymiset {

	e := Edistymiset{}
	bytes, err := ioutil.ReadFile(savefile)
	if err == nil {
		err := json.Unmarshal(bytes, &e.data)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Tila ladattu levyltä!")
	} else if os.IsNotExist(err) {
		log.Println("Tallennustiedostoa ei löytynyt. Aloitan tyhjästä.")
		e.data = make(map[string]Edistyminen)
	} else {
		log.Fatal(err)
	}

	return e
}
