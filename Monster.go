package game

import(
	//"math/rand"
	"github.com/hajimehoshi/ebiten"
)


type Monster struct{
	tilenum int
	moving *Mob
	char *Character
}

func NewMonster(tilenum int, mob *Mob,char *Character) *Monster{
	return &Monster{tilenum,mob,char}
}

func (m Monster)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions, th *TileHolder){
	point,_,_,_ := m.moving.Area.GetCornersIso()
	options.GeoM.Translate(point.IX,point.IY)
	options.GeoM.Translate(0,-float64(th.Tiles[m.tilenum].GetOff()))
	screen.DrawImage(th.DisplayTile(m.tilenum),options)
	options.GeoM.Translate(0,float64(th.Tiles[m.tilenum].GetOff()))
	options.GeoM.Translate(-point.IX,-point.IY)
}

func (m Monster)GetBody() *Body {
	return&Body{m.moving.Area,m.char}
}
