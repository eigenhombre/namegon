# namegon

<img src="namegon.jpg" width="400">

![build](https://github.com/eigenhombre/namegon/actions/workflows/build.yml/badge.svg)

Another random name generator, loosely based on
[`namejen`](https://github.com/eigenhombre/namejen).

# Usage

    go get github.com/eigenhombre/namegon@<version>

E.g.,

    go get github.com/eigenhombre/namegon@v0.0.1

The function `Namer` takes a slice of strings and a "chain length"
*n*, and returns a function which can be called repeatedly to generate
names based on *n*-grams from the supplied input.

# Example

    package main

    import (
        "fmt"

        "github.com/eigenhombre/namegon"
    )

    func main() {
        namer := namegon.Namer(namegon.LatinishNames, 4)
        for i := 0; i < 10; i++ {
            fmt.Println(namer())
        }
        fmt.Println("OK")
    }

Output:

    Ntamino
    Uredo
    Lorum
    Xusamplitus
    Apilludo
    Aepromptus
    Riuria
    Asportunitio
    Dminis
    Erebro
    OK

# Name Seed Lists

The package comes with three "seed lists" ported over from [`namejen`](https://github.com/eigenhombre/namejen):

- `ExampleMaleNames`
- `ExampleFemaleNames`
- `LatinWords`

... but of course, your own [corpora](https://github.com/dariusk/corpora) can be used.

# License: MIT

Copyright (c) 2023 John Jacobsen

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

