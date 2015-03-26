package types

type BitsByte byte

func (b *BitsByte) SetBit(pos uint, flag bool) {
	if flag {
		*b |= 1 << pos
    } else {	
		*b &= ^(1 << pos)
    }    
}