package game
 
import(
 "github.com/hajimehoshi/ebiten"
//"log"
)

type Layer struct{
	fields []Field
}

func (l *Layer)addField(f Field){
	l.fields = append(l.fields,f)
}


type Chunk struct{
	area *Rectangle2D
	layers map[int]*Layer
	qt *QuadTree
	th *TileHolder
	position *PointIso
}

func NewChunk(area *Rectangle2D, th *TileHolder,pos *PointIso) *Chunk{
	ch := &Chunk{}
	ch.area = area
	ch.th = th
	ch.layers = make(map[int]*Layer)
	ch.qt = NewQuadTree(area)
	ch.position = pos
	return ch
}

func (ch *Chunk)AddField(layernum int, field Field){
	if ch.layers[layernum] == nil{
		ch.layers[layernum] = &Layer{}
	}
	if ch.area.PointIsoWithin(field.Position){
		p := field.Position
		p2:= ch.position
		field.Area2d.MoveIso(p.IX,p.IY)
		field.Area2d.MoveIso(p2.IX,p2.IY)
		ch.layers[layernum].addField(field)
	}
}

func (ch *Chunk)AddObject(obj *Object){
	ch.qt.AppendObject(obj)
}

func (ch *Chunk)DrawLayer(layer int,screen *ebiten.Image, options *ebiten.DrawImageOptions){
	options.GeoM.Translate(ch.position.IX, ch.position.IY)
	for j := range ch.layers[layer].fields{
			ch.layers[layer].fields[j].DrawField(screen, options, ch.th)
		}
	options.GeoM.Translate(-ch.position.IX, -ch.position.IY)
}
func (ch *Chunk)DrawObjects(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	options.GeoM.Translate(ch.position.IX, ch.position.IY)
	ch.qt.Draw(screen,options, ch.th)
	options.GeoM.Translate(-ch.position.IX, -ch.position.IY)
}

func (ch Chunk)GetField(apoint *PointIso) int{
	testarea := *ch.area
	testarea.MoveIso(ch.position.IX, ch.position.IY)
	if testarea.IsoIsWithin(apoint.IX, apoint.IY){
		var smallarea Rectangle2D
		var k,l int
		if ch.layers[1] != nil{
		for i := range ch.layers[1].fields{
			smallarea = *ch.layers[1].fields[i].Area2d
			if smallarea.IsoIsWithin(apoint.IX,apoint.IY) {
					p1:= apoint.GetPoint2D()
					k= ch.layers[1].fields[i].GetTile(p1)				
			}
		}
		}
		if ch.layers[2] != nil{
			for i :=range ch.layers[2].fields{
				smallarea = *ch.layers[2].fields[i].Area2d
				if smallarea.IsoIsWithin(apoint.IX, apoint.IY){
					p1:= apoint.GetPoint2D()
					l= ch.layers[2].fields[i].GetTile(p1)
				}
			}
		}
		if k!= 0 {
			return k
		}
		if l != 0 {
			return l
		}
	}
	return 0
}
