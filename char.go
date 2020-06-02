package game

import(
	"github.com/hajimehoshi/ebiten"
	)


type Avatar struct{
	moving *Mob
	char *Character
	DisplayPoint *PointIso
	ColPo *PointIso
	Pathblock bool
	images map[int]int
	stance,returnstance int
	counter int
	attack bool
	charbody *Body
}

func (ava *Avatar)Blocked(blocked bool){
	ava.Pathblock = blocked
}


func NewAvatar(area *Rectangle2D, displayPos *PointIso, char *Character) *Avatar{
	ava := &Avatar{}
	ava.charbody = &Body{}
	ava.moving = NewMob(NewPointIso(0,0),area)
	ava.charbody.Moving = area
	ava.charbody.Char = char
	ava.char = char
	ava.DisplayPoint = displayPos
	ava.images = make(map[int]int)
	ava.stance = 2
	ava.returnstance = 2
	return ava
}


func (char *Avatar)Move2D(x,y float64){
	if y >0 {
		char.stance = 2
	}else if ( y>0 && x>0 ){
		char.stance = 4
	}else if y < 0 {
		char.stance = 6
	}else if(y<0 &&x<0){
		char.stance = 8
	}else if x <0 {
		char.stance = 10
	}else if (x<0 && y >0) {
		char.stance = 12
	}else if x >0  {
		char.stance = 14
	}else if (x >0 && y<0){
		char.stance =16
	} 
	xI,yI := TwoDtoIso(x,y)
	if xI !=0 || yI != 0{
		char.DisplayPoint.IX += xI/5
		char.DisplayPoint.IY += yI/5
	}
}

func (char *Avatar)SetTile(stance int, tilenum int){
	char.images[stance] = tilenum
}

func (char Avatar)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions, th *TileHolder){
	options.GeoM.Translate(char.DisplayPoint.IX, char.DisplayPoint.IY - float64(th.Tiles[char.images[char.stance]].GetOff()))
	screen.DrawImage(th.DisplayTile(char.images[char.stance]),options)
	options.GeoM.Translate(-char.DisplayPoint.IX, -char.DisplayPoint.IY+ float64(th.Tiles[char.images[char.stance]].GetOff()))
}

func (ava Avatar)GetBody() *Body{
 return ava.charbody
}
func (char *Avatar)Steer(){
	var f uint8
	if ebiten.IsKeyPressed(ebiten.KeyA){
		f |= 0x1
	}
	if ebiten.IsKeyPressed(ebiten.KeyW){
		f  |= 0x4
	}
	if ebiten.IsKeyPressed(ebiten.KeyD){
		f |= 0x2
	}
	if ebiten.IsKeyPressed(ebiten.KeyS){
		f|= 0x8
	}
	var b float64
	b = 1
	if char.Pathblock{
		b = 0.000000001
	}
	if f == 0x8 {
		char.Move2D(0*b,2*b)
	} else if f == 0xA {
		char.Move2D(1*b,1*b)
	}else if f == 0x4{
		char.Move2D(0*b,-2*b)
	}else if f == 0x5{
		char.Move2D(-1*b,-1*b)
	}else if f == 0x1{
		char.Move2D(-2*b,0*b)
	}else if f == 0x9{
		char.Move2D(-1*b, 1*b)
	}else if f == 0x2{
		char.Move2D(2*b,0*b)
	}else if f == 0x6{
		char.Move2D(1*b,-1*b)
	} else {
		char.stance = char.returnstance
	}
}

func (char *Avatar)Attack() *AnAttack{
	if char.counter >0{
		char.counter--
	}
	if char.counter == 0{
			char.attack = false
	}
	char.returnstance = char.stance
	if ebiten.IsKeyPressed(ebiten.KeySpace)|| ebiten.IsKeyPressed(ebiten.KeyJ){
		char.stance +=1
		if char.attack == false {
			char.counter = 120
		}
		
		if char.attack == false{
			strike, achar := char.char.CreateStrikeEvent()
			char.attack = true
			return NewAttack(strike,achar,char.GenerateStrikeField(),20)
		}
	}
	return nil
}

func (char *Avatar)Controll() *AnAttack{
	char.Steer()
	h:=char.Attack()
	if char.counter == 120{
		return h
	}
	return nil
}

func (char *Avatar)CollisionPoint() {
	var p2D *PointIso
	a:= *char.moving.Area
	a.PositionIso(char.DisplayPoint.IX,char.DisplayPoint.IY)
	a.Move2D(5,-10)
	var l,u,r,d *PointIso
	if char.returnstance  == 2 || char.returnstance   == 6 || char.returnstance   == 10 || char.returnstance   == 14{
		l,u,r,d   = a.GetEdgesCenterIso()
	} 
	if char.returnstance   == 4 || char.returnstance   == 10|| char.returnstance   == 12 ||char.returnstance  == 16{
		u,l,d,r  = a.GetCornersIso()
	}
	if char.returnstance   == 2 || char.returnstance   == 4{
		p2D = d
	}
	if char.returnstance   == 6 || char.returnstance   == 8{
		p2D = u
	}
	if char.returnstance   == 10 || char.returnstance   == 12{
		p2D = l
	} 
	if char.returnstance   == 14 || char.returnstance   == 16{
		p2D = r
	}
	//log.Println("Display Point: ", p2D)
	char.ColPo = p2D
}

func (ava Avatar)GenerateStrikeField() *Rectangle2D{
	a:= *ava.moving.Area
	a.PositionIso(ava.DisplayPoint.IX,ava.DisplayPoint.IY)
	a.Move2D(5,-10)
	l := a.GetLength()
	h := a.GetHeight()
	if ava.returnstance == 2 {
		a.Move2D(0,h)
	}
	if ava.returnstance ==6{
		a.Move2D(0,-h)
	}
	if ava.returnstance ==10 {
		a.Move2D(-l,0)
	}
	if ava.returnstance ==14 {
		a.Move2D(l,0)
	}
	return &a
}
