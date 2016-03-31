package mylib

import ( "fmt")

/*
	just a test.. no real use
*/

type Ort struct {
    Bez string
    Ref1 *Ort
    Ref2 *Ort
    id int
    zeit int
}

func PrintRef1 (x Ort){
    fmt.Println(x.Bez + " Ref1=" +  x.Ref1.Bez)        
}

