package cyberdemon

import "fmt"

type Cyberdemon interface {
	Ammu()
	Kuollut() bool
}

func New() Cyberdemon {
	cd := cyberdemon(1024)
	return &cd
}

type cyberdemon int

func (c *cyberdemon) Ammu() {
	*c--
	fmt.Println("AAAGH!")
	if c.Kuollut() {
		fmt.Println("Kyberdemoni räjähti.")
	}
}

func (c cyberdemon) Kuollut() bool {
	return c <= 0
}
