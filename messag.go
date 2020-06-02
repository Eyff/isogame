package game

import(
)
type Message interface{
	Update()
	GetTerminated() bool
	GetArea() *Rectangle2D
	GetChar() *Character
	GetDMP() *string
	}

type DPD struct{
	loc *PointIso
	damage string
	count int
	terminated bool
}

func (d *DPD)Update(){
	if d.count < 0{
		d.terminated = true
	}
	d.loc.IY -= 0.4
	d.count--
}

func (d DPD)GetTerminated() bool{
	return d.terminated
}

func (d DPD)GetChar() *Character{
	return nil
}

func (d DPD)GetArea() *Rectangle2D{
	return nil
}

func (d DPD)GetDMP() *string{
	return &d.damage
}

func (d DPD)Terminate() bool{
	return d.terminated
}


type TechValue struct{
	tv int
	Char *Character
}

func (tv TechValue)Update(){}
func (tv TechValue)GetTerminated() bool{
	return false
}
func (tv TechValue)GetChar() *Character{
	return tv.Char
}

func (tv TechValue)GetArea() *Rectangle2D{
	return nil
}

func (tv TechValue)GetDMP() *string{
	return nil
}

