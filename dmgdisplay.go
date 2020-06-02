package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/text"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font"
	"io/ioutil"
	"image/color"
	"container/list"
	"log"
	"strconv"
)

type DPD struct{
	loc *PointIso
	damage string
	count int
	terminated bool
}

func (d *DPD)Update(){
	if d.count < 0{
		d.terminated = true
	}
	d.loc.IY -= 0.4
	d.count--
}

func (d DPD)Terminate() bool{
	return d.terminated
}


type DMGDisplayList struct{
	DMGP  *list.List
	font *truetype.Font
	mplusNormalFont font.Face
	color color.Color
}
func (DDL *DMGDisplayList)Init() {
	DDL.DMGP  =list.New()
	fontBytes, err := ioutil.ReadFile("Blockstepped.ttf")
	if err != nil {
		log.Println(err)
		return
	}
	var err2 error
	DDL.font, err2 = truetype.Parse(fontBytes)
	if err2 != nil {
		log.Fatal(err)
	}
	DDL.mplusNormalFont = truetype.NewFace(DDL.font, &truetype.Options{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingFull,
	})
}

func (DDL *DMGDisplayList)GetDMGP(dmgPoints int, loc *PointIso, ahit bool){
	if ahit{
		a:=strconv.Itoa(dmgPoints)
		v:= &DPD{loc,a,150,false}
		DDL.DMGP.PushBack(v)
	}
}

func (DDL DMGDisplayList)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions){
	for e := DDL.DMGP.Front(); e != nil; e = e.Next(){
		s:= e.Value.(*DPD)		
		text.Draw(screen, s.damage, DDL.mplusNormalFont,int(s.loc.IX), int(s.loc.IY), color.White)
	}
}

func (DDL *DMGDisplayList)Update() error{
	for e := DDL.DMGP.Front(); e != nil; e = e.Next(){
		s:= e.Value.(*DPD)
		s.Update()
		if s.Terminate(){
			 DDL.DMGP.Remove(e)
		}
	}
	return nil
}
