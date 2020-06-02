package game

import (
)

type Strikeevent struct{
	Attack, Ability, Skill, Luck int
}

type Body struct{
	Moving *Rectangle2D
	Char *Character
}

func NewStrikeEvent(attack, ability, skill, luck int) *Strikeevent{
	return &Strikeevent{attack,ability, skill, luck}
}
	
type AnAttack struct{
	strike *Strikeevent
	char *Character
	area *Rectangle2D
	count int
	terminated bool
}

func NewAttack(strike *Strikeevent, char *Character, area *Rectangle2D, count int) *AnAttack{
	return &AnAttack{strike,char,area,count,false}
}
	

func (a *AnAttack)Update(){
	if a.count == 0{
		a.terminated = true
	}
	a.count--
}

func (a AnAttack)GetTerminated() bool{
	return a.terminated
}

func (a AnAttack)GetCharacter() *Character{
	return a.char
 }
 
func (a AnAttack)GetChar() *Character{
	return nil
}

func (a AnAttack)GetDMP() *string {
	return nil	
}
func (a AnAttack)GetArea() *Rectangle2D{
	return a.area
}

func (a AnAttack)Terminated() bool{
	return a.terminated
}


func (a AnAttack)GetStrike() *Strikeevent{
	a.count = 0
	return a.strike
}


func (a AnAttack)DidHit(r2D *Rectangle2D) bool{
	return a.area.Intersects(r2D)
}

func (a AnAttack)Trapped(pIso *PointIso) bool{
	return a.area.IsoIsWithin(pIso.IX,pIso.IY)
}
