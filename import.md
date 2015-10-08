# `import`

`import`-komennolla otetaan käyttöön kirjastoja. Sanan `import` jälkeen tulee kirjaston polku. Kirjaston nimi on se osa polkua, joka tulee viimeisen kauttaviivan jälkeen.
- Go:n mukana tulevan kirjaston käyttämiseen riittää pelkkä kirjaston nimi
- Muut kirjastot etsitään Go-kansiosta. Esimerkiksi minulla on kopio kirjastosta _vec2_ kansiossa `/home/joonazan/go/src/github.com/joonazan/vec2`.
- Kirjaston polku voi samalla olla myös URL, josta sen saa. Paketin ja kaikki sen vaatimat

```Go
import "math"
import "math/rand"
import "github.com/joonazan/vec2"
```
