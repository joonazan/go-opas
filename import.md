# `import`

`import`-komennolla otetaan käyttöön paketteja. Sanan `import` jälkeen tulee paketin polku. Paketin nimi on se osa polkua, joka tulee viimeisen kauttaviivan jälkeen.
- Go:n mukana tulevien pakettien käyttämiseen riittää pelkkä paketin nimi
- Go etsii muut paketit Go-kansiosta. Esimerkiksi minulla on kopio kirjastosta _vec2_ kansiossa `/home/joonazan/go/src/github.com/joonazan/vec2`.
- Kirjaston polku voi samalla olla myös URL, josta sen saa. Paketin ja kaikki sen vaatimat paketit saa ladattua komennolla `go get`. Esim. `go get "github.com/joonazan/vec2"`.

```Go
import "math"
import "math/rand"
import "github.com/joonazan/vec2"
```
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
Paketit voi importata eri nimellä laittamalla halutun nimen polun eteen. Alaviivan laittaminen nimen kohdalle importtaa kirjaston, mutta siihen ei pääse käsiksi millään nimellä.
```Go
import (
	glfw "github.com/go-gl/glfw3"
	_ "image/png"
)
```
