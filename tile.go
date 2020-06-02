package game

import (
"image"
"github.com/hajimehoshi/ebiten"
"github.com/hajimehoshi/ebiten/ebitenutil"
"log"
)

type Tile struct{
	L,U,R,D,Off int
	Tilelength int
	Block bool
	Area2D *Rectangle2D
}

func NewTile(l,u,r,d,off, tilelength int, block bool) *Tile{
	r2d := Calc2DRectangleFromIsoInt(l,u,r,d)
	return &Tile{l,u,r,d,off,tilelength,block,r2d}
}

func (t Tile)GetRect() image.Rectangle{
	return image.Rect(t.L, t.U-t.Off, t.R, t.D)
}

func (t Tile)GetOff() int{
	return t.Off
}

func (t Tile)GetBlock() bool{
	return t.Block
}

func (t Tile)GetArea2D() *Rectangle2D{
	return t.Area2D
}

type TileHolder struct{
	Sheets map[int]*ebiten.Image
	Tiles map[int]*Tile
	TileSheet map[int]int
}

func NewTileHolder() *TileHolder{
	th := &TileHolder{}
	th.Sheets= make(map[int]*ebiten.Image)
	th.Tiles = make(map[int]*Tile)
	th.TileSheet = make(map[int]int)
	return th
}

func (th *TileHolder)AddImage(sheetnum int, filename string) {
	eimg,_, err:= ebitenutil.NewImageFromFile(filename, ebiten.FilterDefault)
	if err != nil{
		log.Fatal(err)
	}
	th.Sheets[sheetnum]= eimg
}

func (th *TileHolder)AddTile(tilenum int, sheetnum int, tile *Tile){
	th.TileSheet[tilenum] = sheetnum
	th.Tiles[tilenum] = tile
}

func (th TileHolder)DisplayTile(tilenum int) *ebiten.Image {
	return th.Sheets[th.TileSheet[tilenum]].SubImage(th.Tiles[tilenum].GetRect()).(*ebiten.Image)
}

func (th TileHolder)GetTileLength(tilenum int) int {
	return th.Tiles[tilenum].Tilelength
}
