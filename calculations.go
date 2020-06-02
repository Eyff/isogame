package game

import(
//"log"
	)
	
	
func TwoDtoIso(x,y float64) (float64,float64){
	ix := x-y
	iy := (x+y)/2
	return ix,iy
}

func IsoTo2D(ix,iy float64) (float64,float64){
	x :=(ix+2*iy)/2
	y :=(2*iy -ix)/2
	return x,y
}

type PointIso struct{
	IX,IY float64
}

func (pI PointIso)GetPoint2D() *Point2D{
	
	x,y := IsoTo2D(pI.IX, pI.IY)

	return NewPoint2D(x,y)
}

func NewPointIso(x,y float64) *PointIso {
	return &PointIso{x,y}
}

type Point2D struct{
	X,Y float64
}

func (p2 Point2D)GetPointIso() *PointIso{

	ix,iy:= TwoDtoIso(p2.X,p2.Y)
	
	return NewPointIso(ix,iy)
}

func NewPoint2D(x,y float64) *Point2D {
	return &Point2D{x,y}
}

type Rectangle2D struct{
	L,U,R,D float64
}

func Calc2DRectangleFromIsoInt(l,u,r,d int) *Rectangle2D{
	l2,u2 := IsoTo2D(float64(l), float64(u))
	r2,d2 := IsoTo2D(float64(r),float64(d))
	return &Rectangle2D{l2,u2,r2,d2}
}

func Calc2DRectangleFromIsoFloat(l,u,r,d float64) *Rectangle2D{
	l2,u2 := IsoTo2D(l,u)
	r2,d2 := IsoTo2D(r,d)
	return &Rectangle2D{l2,u2,r2,d2}
}

func (r2d Rectangle2D)IsoIsWithin(x,y float64) bool{
	x2d,y2d := IsoTo2D(x,y)
	if x2d > r2d.L && x2d < r2d.R && y2d > r2d.U && y2d < r2d.D {
		return true
	}
	return false
}

func (r2d Rectangle2D)TwoDIsWithin(x,y float64) bool{
	if x >r2d.L && x < r2d.R && y > r2d.U && y < r2d.D {
		return true
	}
	return false
}

func (r2d Rectangle2D)Intersects(rtd *Rectangle2D) bool {
	if r2d.TwoDIsWithin(rtd.L, rtd.U)|| r2d.TwoDIsWithin(rtd.R, rtd.U)|| r2d.TwoDIsWithin(rtd.L,rtd.D)|| r2d.TwoDIsWithin(rtd.R,rtd.D){
		return true
	}
	if rtd.TwoDIsWithin(r2d.L,r2d.U)|| rtd.TwoDIsWithin(r2d.R, r2d.U)|| rtd.TwoDIsWithin(r2d.L,r2d.D)|| rtd.TwoDIsWithin(r2d.R, r2d.D){
		return true
	}
	return false
}

func (r2D *Rectangle2D)Move2D(x,y float64){
	r2D.L += x
	r2D.R += x
	r2D.U += y
	r2D.D +=y
}

func (r2D *Rectangle2D)MoveIso(ix,iy float64){
	x,y := IsoTo2D(ix,iy)
	r2D.Move2D(x,y)
}

func (r2D *Rectangle2D)Position2D(x,y float64){
	rectl := r2D.R - r2D.L
	recth := r2D.D - r2D.U
	r2D.L = x
	r2D.R = x+rectl
	r2D.U = y
	r2D.D = y + recth
}

func (r2D *Rectangle2D)PositionIso(ix,iy float64) {
	x,y :=  IsoTo2D(ix,iy)
	r2D.Position2D(x,y)
}

func (r2D Rectangle2D)PointIsoWithin(p *PointIso) bool{
	return r2D.IsoIsWithin(p.IX,p.IY)
}

func (r2D Rectangle2D)GetLength() float64{
	return r2D.R - r2D.L
}

func (r2D Rectangle2D)GetCenter2D() *Point2D{
	x:=(r2D.R-r2D.L)/2+r2D.L
	y:=(r2D.D-r2D.U)/2+r2D.U
	return NewPoint2D(x,y)
}

func (r2D Rectangle2D)GetCenterIso() *PointIso{
	p2D := r2D.GetCenter2D()
	return p2D.GetPointIso()
}

func (r2D Rectangle2D)GetHeight() float64{
	return r2D.D - r2D.U
}

func (r2D Rectangle2D)GetCornersIso() (*PointIso,*PointIso,*PointIso,*PointIso){
	lup := NewPoint2D(r2D.L, r2D.U)
	ldo := NewPoint2D(r2D.L, r2D.D)
	rup := NewPoint2D(r2D.R, r2D.U)
	rdo := NewPoint2D(r2D.R, r2D.D)
	return lup.GetPointIso(), ldo.GetPointIso(), rup.GetPointIso(), rdo.GetPointIso()
}

func (r2D Rectangle2D)GetEdgesCenterIso() (*PointIso,*PointIso,*PointIso,*PointIso){
	x:=r2D.GetLength()/2
	y:=r2D.GetHeight()/2
	l := NewPoint2D(r2D.L,r2D.U+y)
	u := NewPoint2D(r2D.L+x, r2D.U)
	r := NewPoint2D(r2D.R,r2D.U+y)
	d := NewPoint2D(r2D.L+x, r2D.D)
	return l.GetPointIso(), u.GetPointIso(), r.GetPointIso(), d.GetPointIso()
}
