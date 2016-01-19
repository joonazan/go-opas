# Import

`import`-komennolla otetaan käyttöön paketteja. Sanan `import` jälkeen tulee paketin polku. Paketin nimi on se osa polkua, joka tulee viimeisen kauttaviivan jälkeen.

Go:n mukana tulevien pakettien käyttämiseen riittää pelkkä paketin nimi
```Go
import "math"
```
Go etsii muut paketit Go-kansiosta. Esimerkiksi minulla on kopio paketista _vec2_ kansiossa `/home/joonazan/go/src/github.com/joonazan/vec2`, koska minun GOPATH on `/home/joonazan/go`.
```Go
import "gogame/bresenham"
import "github.com/joonazan/vec2"
```
Paketin polku voi samalla olla myös URL, josta sen saa. Paketin ja kaikki sen vaatimat paketit saa ladattua komennolla `go get`. Esim. `go get "github.com/joonazan/vec2"`.

Monen asian importtaaminen ei vaadi importin moneen kertaan kirjoittamista.
```Go
import (
	"fmt"
	"github.com/go-gl-legacy/gl"
	"io/ioutil"
	"log"
	"strings"
)
```
Paketit voi importata eri nimellä laittamalla halutun nimen polun eteen.
```Go
import glfw "github.com/go-gl/glfw3"
```
Alaviivan laittaminen nimen kohdalle importtaa paketin, mutta siihen ei pääse käsiksi millään nimellä. Paketti siis tuodaan pelkästään sivuvaikutusten vuoksi. Esimerkiksi _image/png_ saa _image_-paketin lataamaan png-muotoisia kuvia.
```Go
import (
	"image"
	_ "image/png"
)
```
Pisteen laittaminen nimen kohdalle tuo paketin sisällön tämänhetkisen paketin nimiavaruuteen. Näin ei pitäisi tehdä, ellei siihen ole todella hyvää syytä.
```Go
package main

import . "fmt"

func main() {
	Println("Tämä on ongelmallista jos fmt:ssä ja tässä paketissa on samannimisiä asioita.")
}
```
