package game

import(
)

type Mob struct{
	Position *PointIso
	Area *Rectangle2D
}

func NewMob(pos *PointIso, area *Rectangle2D) *Mob{
	area.PositionIso(pos.IX,pos.IY)
	return &Mob{pos, area}
}

func (m *Mob)PositionPIso(pI *PointIso){
	m.Position = pI
	m.Area.PositionIso(pI.IX, pI.IY)
}

func (m *Mob)PositionIso(ix,iy float64){
	m.Position = NewPointIso(ix,iy)
	m.Area.PositionIso(ix,iy)
}

func (m *Mob)PositionP2D(p2d *Point2D){
	pIso := p2d.GetPointIso()
	m.Position = pIso
	m.Area.PositionIso(pIso.IX,pIso.IY)
}

func (m *Mob)Position2d(x,y float64){
	xI,yI := TwoDtoIso(x,y)
	m.Position = NewPointIso(xI,yI)
	m.Area.PositionIso(xI,yI)
}


func (m *Mob)MoveIso(xI,yI float64){
	m.Position.IX = m.Position.IX+xI
	m.Position.IY = m.Position.IY+yI
	m.Area.MoveIso(xI,yI)
}

func (m *Mob)Move2D(x,y float64){
	xI,yI:= TwoDtoIso(x,y)
	m.Position.IX = m.Position.IX+xI
	m.Position.IY = m.Position.IY+yI
	m.Area.MoveIso(xI,yI)
}
