package game

import (
	"container/list"
	//"log"
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

type DMGDIsplay struct{
	dmg int
	DPoint *PointIso
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

func (a AnAttack)Terminated() bool{
	return a.terminated
}

func (a AnAttack)GetStrike() *Strikeevent{
	a.count = 0
	return a.strike
}

func (a AnAttack)GetCharacter() *Character{
	return a.char
 }

func (a AnAttack)DidHit(r2D *Rectangle2D) bool{
	return a.area.Intersects(r2D)
}

func (a AnAttack)Trapped(pIso *PointIso) bool{
	return a.area.IsoIsWithin(pIso.IX,pIso.IY)
}

type AttackList struct{
	attacks *list.List
	bodys *list.List
}
func (al *AttackList)RegisterAttack(a *AnAttack){
	if a != nil{
		al.attacks.PushBack(a)
	}
}

func (al *AttackList)RegisterBody(body *Body){
	f:= true
	for d:= al.bodys.Front(); d!= nil; d = d.Next(){
		t:= d.Value.(*Body)
		if t.Char != body.Char{
			f= false
		}
	}
	if f{

			al.bodys.PushBack(body)
		}
}

func (al *AttackList)Initialize(){
	if al.attacks == nil{
		al.attacks = list.New()
	}
	if al.bodys == nil{
		al.bodys = list.New()
	}
}

func (al *AttackList)Update()   (int, *PointIso,bool){
	var dmg int
	var area *PointIso
	var hit bool
	for e := al.attacks.Front(); e != nil; e = e.Next(){
		s := e.Value.(*AnAttack)
		for d:= al.bodys.Front(); d!= nil; d = d.Next(){
			t:= d.Value.(*Body)
			if s.DidHit(t.Moving){
				if dmg, hit = t.Char.DealWithStrikeEvent(s.GetStrike(),s.GetCharacter()); hit{
					s.count = 0
					area = t.Moving.GetCenterIso()
				}
			}
		}
		s.Update()
		if s.Terminated(){
			defer al.attacks.Remove(e)
		}
		return dmg, area, hit
	}
	return 0, &PointIso{-10,-10}, false
}
