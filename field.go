package game

import(
 "math"
 "github.com/hajimehoshi/ebiten"
//  "log"
)

type Field struct{
	Tiles []int
	Area2d *Rectangle2D
	Position *PointIso
}

func NewField(tilenums []int, tilelength int, pos *PointIso) *Field {
	sidesteps := int(math.Sqrt(float64(len(tilenums))))
	sidelength := tilelength*sidesteps
	area2d := &Rectangle2D{0,0,float64(sidelength/2), float64(sidelength/2)}
	return &Field{tilenums,area2d,pos}
}

func (f Field)DrawField(screen *ebiten.Image, options *ebiten.DrawImageOptions, th *TileHolder){
	options.GeoM.Translate(f.Position.IX, f.Position.IY)
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++{
			k := j + i*10
			l,m := TwoDtoIso(float64(j),float64(i))
			options.GeoM.Translate(l*float64(th.Tiles[f.Tiles[k]].Tilelength/2), m*float64(th.Tiles[f.Tiles[k]].Tilelength/2)- float64(th.Tiles[f.Tiles[k]].Off))
			screen.DrawImage(th.DisplayTile(f.Tiles[k]),options)
			options.GeoM.Translate(-l*float64(th.Tiles[f.Tiles[k]].Tilelength/2), -m*float64(th.Tiles[f.Tiles[k]].Tilelength/2)+ float64(th.Tiles[f.Tiles[k]].Off))
		}
	}
	options.GeoM.Translate(-f.Position.IX, -f.Position.IY)
}

func (f Field)GetTile(apoint *Point2D) int{
	testx := apoint.X - f.Area2d.L
	testy :=apoint.Y - f.Area2d.U
	k :=int(testx)/10 +(int(testy)/10)*10
	return f.Tiles[k]
}


type Object struct{
	Graphic int
	Position *PointIso
	Block bool
	Area2D *Rectangle2D
}

type ObjectCreator struct {
	tileHolder *TileHolder
}

func NewObjectCreator(th *TileHolder) *ObjectCreator {
	return &ObjectCreator{th}
}

func (oc *ObjectCreator)NewObject(tilenum int, position *PointIso, block bool) *Object{
	area2d := oc.tileHolder.Tiles[tilenum].GetArea2D()
	return &Object{tilenum, position, block, area2d}
}


