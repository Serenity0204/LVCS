package resources

import (
	_ "embed"
)

//go:embed ASCIIArts/ascii1.txt
var ASCII1 string

//go:embed ASCIIArts/ascii2.txt
var ASCII2 string

//go:embed ASCIIArts/ascii3.txt
var ASCII3 string

//go:embed ASCIIArts/ascii4.txt
var ASCII4 string

//go:embed ASCIIArts/ascii5.txt
var ASCII5 string

//go:embed docs/list.txt
var LIST string

//go:embed docs/detail.txt
var DETAIL string
