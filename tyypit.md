# Tyypit

Kaikki "esineet" ohjelmoinnissa ovat arvoja. `"Hello, world!"` on arvo. `fmt.Println` on arvo. `3` on arvo. `func` ei ole arvo. Eikä `package`. Edes `fmt` ei ole arvo. (Mutta `"fmt"` on.)

Kuitenkaan kaikki arvot eivät ole samanlaisia. Mitä on `"Hello, world!" + 3`? Entä `3 + fmt.Println`? Jokaisella arvolla on tyyppi.

Kaikki tyypit rakentuvat yksinkertaisista tyypeistä.

```Go
bool        true (tosi) tai false (epätosi)

uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32

string      merkkijono
```

```Go
package main

import "fmt"

func main() {
	fmt.Printf("%T", fmt.Println)
}
```
tulostaa: `func(...interface {}) (int, error)`
