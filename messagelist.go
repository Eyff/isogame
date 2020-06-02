package game

import (
"container/list"
"log"
)


type MessageList struct{
	messages *list.List
	bodies *list.List
	display *DMGDisplayList
}

func (ml *MessageList)Init(disp *DMGDisplayList){
	ml.messages = list.New()
	ml.bodies = list.New()
	ml.display = disp
}

func (ml *MessageList)RegisterBody(bo *Body){
	f:= true
	for b:= ml.bodies.Front();b!= nil; b= b.Next(){
		p := b.Value.(*Body)
		if bo.Char == p.Char{
			f = false
		}
	}
	if f{
		log.Println("registerd Bodies")
		ml.bodies.PushBack(bo)
	}
}

func (ml *MessageList)Update(){
	for b := ml.bodies.Front(); b!= nil; b = b.Next(){
		for m:= ml.messages.Front(); m!= nil; m= m.Next(){
				n := m.Value.(Message)
				if a := n.GetArea();a!= nil{
					log.Println("Did strike")
					p := b.Value.(*Body)
					if a.Intersects(p.Moving){
						o:= n.(*AnAttack)
						d, tv:=p.Char.DealWithAttack(o)
						ml.GetMessage(d)
						ml.GetMessage(tv)
						ml.messages.Remove(m)
					}
				}
				if a:= n.GetChar(); a!= nil{
					p:= b.Value.(*Body)
					if a== p.Char {
						o := n.(*TechValue)
						p.Char.GetTechValue(o)
						ml.messages.Remove(m)
					}
				}
				if a:= n.GetDMP(); a!= nil{
					o := n.(*DPD)
					ml.display.GetDPD(o)
					ml.messages.Remove(m)
				}
			}
		}
	for m:= ml.messages.Front(); m!= nil; m= m.Next(){
		n := m.Value.(Message)
		n.Update()
		if n.GetTerminated(){
			defer ml.messages.Remove(m)
		 }
	}
}


func (ml *MessageList)GetMessage(e Message){
	ml.messages.PushBack(e)
}

