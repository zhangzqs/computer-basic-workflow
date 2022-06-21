package main

func Ins2OpAd(ins uint16) (op uint16, ad uint16) {
	op = ins >> 10
	ad = ins & 0b1111111111
	return
}

func OpAd2Ins(op uint16, ad uint16) (ins uint16) {
	ins = op << 10
	ins = ins | ad
	return
}
