package main

import(
"log"
"github.com/hajimehoshi/ebiten"
"github.com/hajimehoshi/ebiten/ebitenutil"
"github.com/Eyff/isogame"
"image/color"
)

// Game implements ebiten.Game interface.
type Game struct{
	th *game.TileHolder
	oc *game.ObjectCreator
	op *ebiten.DrawImageOptions
	ch *game.Chunk
	ca *game.Avatar
	al *game.AttackList
	ddl *game.DMGDisplayList
	mo *game.Monster
	loaded bool
	}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update(screen *ebiten.Image) error {
	if !g.loaded{
		g.op= &ebiten.DrawImageOptions{}
		g.th = game.NewTileHolder()
		g.ddl = &game.DMGDisplayList{}
		g.ddl.Init()
		g.th.AddImage(0,"isotiles2.png")
		g.th.AddTile(0,0,game.NewTile(23,1,42,11,0,20,false))
		g.th.AddTile(1,0,game.NewTile(2,1,21,11,0,20,false))
		g.th.AddTile(2,0,game.NewTile(64,23,83,33,14,20,true))
		g.th.AddTile(3,0,game.NewTile(85,23,104,32,19,20,true))
		g.th.AddTile(4,0,game.NewTile(106,23,125,32,18,20,true))
		g.th.AddTile(5,0,game.NewTile(64,56,83,65,18,20,true))
		g.th.AddTile(6,0,game.NewTile(85,56,104,65,22,20,true))
		g.th.AddTile(7,0,game.NewTile(106,56,125,65,22,20,true))
		g.th.AddTile(100,0,game.NewTile(877,67,928,92,72,39,true))
		g.th.AddTile(101,0,game.NewTile(937,12,1011,44,12,85, true))
		g.th.AddTile(1001,0,game.NewTile(436,56,486,73,48,41,true))
		g.th.AddTile(1002,0,game.NewTile(484,56,534,73,48,50,true))
		g.th.AddTile(1003,0,game.NewTile(546,50,597,67,48,50,true))
		g.th.AddTile(1004,0,game.NewTile(601,52,651,69,48,50,true))
		g.th.AddTile(1005,0,game.NewTile(650,51,700,68,48,50,true))
		g.th.AddTile(1006,0,game.NewTile(713,51,763,68,48,50,true))
		g.th.AddTile(1007,0,game.NewTile(435,127,485,144,48,50,true))
		g.th.AddTile(1008,0,game.NewTile(486,129,536,146,48,50,true))
		g.th.AddTile(1009,0,game.NewTile(544,134,594,151,48,50,true))
		g.th.AddTile(1010,0,game.NewTile(601,126,651,143,48,50,true))
		g.th.AddTile(1011,0,game.NewTile(652,123,702,140,48,50,true))
		g.th.AddTile(1012,0,game.NewTile(719,127,769,144,48,50,true))
		g.th.AddTile(1013,0,game.NewTile(726,211,749,219,18,24,true))
		g.oc = game.NewObjectCreator(g.th)
		cp := game.NewPoint2D(300,-100)
		g.ch = game.NewChunk(&game.Rectangle2D{0,0,440,440},g.th,cp.GetPointIso())
		p:= game.NewPoint2D(195,70)
		
		g.ch.AddObject(g.oc.NewObject(100,p.GetPointIso(),true))
		
		p1 := game.NewPoint2D(150,150)
		g.ch.AddObject(g.oc.NewObject(101,p1.GetPointIso(),true))
		p1 = game.NewPoint2D(200,100)
		g.ch.AddObject(g.oc.NewObject(100,p1.GetPointIso(),true))
		a:=[]int{1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,}
		b := game.NewPoint2D(1,1)
		g.ch.AddField(0,*game.NewField(a,20,b.GetPointIso()))
		c:= game.NewPoint2D(101,1)
		g.ch.AddField(0,*game.NewField(a,20,c.GetPointIso()))
		d:= game.NewPoint2D(201,1)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d= game.NewPoint2D(301,1)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(1,101)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(101,101)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(201,101)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(301,101)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(1,201)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(101,201)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(201,201)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(301,201)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(1,301)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(101,301)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(201,301)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		d = game.NewPoint2D(301,301)
		g.ch.AddField(0,*game.NewField(a,20,d.GetPointIso()))
		e:=[]int{6,7,7,7,7,7,7,7,7,7,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,}
		f:= game.NewPoint2D(1,1)
		log.Println("layer 1 field 1: " , f)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))
		e= []int{7,7,7,7,7,7,7,7,7,7,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,}
		f=game.NewPoint2D(101,1)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))
		e= []int{7,7,7,7,7,7,7,7,7,7,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,}
		f=game.NewPoint2D(201,1)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{7,7,7,7,7,7,7,7,7,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,}
		f=game.NewPoint2D(301,1)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,}
		f= game.NewPoint2D(1,101)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,}
		f= game.NewPoint2D(301,101)
		g.ch.AddField(2,*game.NewField(e,20,f.GetPointIso()))		 		 
		e=[]int{6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,}
		f= game.NewPoint2D(1,201)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))	
		e=[]int{0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,}
		f= game.NewPoint2D(301,201)
		g.ch.AddField(2,*game.NewField(e,20,f.GetPointIso()))	
		e=[]int{6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,6,0,0,0,0,0,0,0,0,0,}
		f= game.NewPoint2D(1,301)
		g.ch.AddField(1,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4,4,4,4,4,4,4,4,4,}
		f= game.NewPoint2D(1,301)
		g.ch.AddField(2,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4,4,4,4,4,4,4,4,4,4,}
		f= game.NewPoint2D(101,301)
		g.ch.AddField(2,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,4,4,4,4,4,4,4,4,4,4,}
		f= game.NewPoint2D(201,301)
		g.ch.AddField(2,*game.NewField(e,20,f.GetPointIso()))
		e=[]int{0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,0,0,0,0,0,0,0,0,0,3,4,4,4,4,4,4,4,4,4,3,}
		f= game.NewPoint2D(301,301)
		g.ch.AddField(2,*game.NewField(e,20,f.GetPointIso()))	
		g.al = &game.AttackList{}
		g.al.Initialize()
		log.Println("Character creation")	
		trt := game.NewTraits(101,31, 41,51,31,21,21, 16,21,5)
		trtb := game.NewClassLevelTraits(5,1.5,2.334,2,1.75,0.75,0.65,0.9,1.1,0.2)
		fighter :=game.NewClass(trt,trtb)
		p4:= game.NewPoint2D(400,100)
		attr := game.NewAttributes(3,3,3,3,3)
		skill := game.NewSkills(0,0)
		exp := &game.Expirience{0,0}
		player := game.NewCharacter(attr,skill, exp,fighter)
		log.Println("Avatar creation")
		g.ca = game.NewAvatar(&game.Rectangle2D{50,50,82,67},p4.GetPointIso(),player)
		g.ca.SetTile(2,1001)
		g.ca.SetTile(3,1002)
		g.ca.SetTile(4,1001)
		g.ca.SetTile(5,1003)
		g.ca.SetTile(6,1004)
		g.ca.SetTile(7,1005)
		g.ca.SetTile(8,1004)
		g.ca.SetTile(9,1006)
		g.ca.SetTile(10,1007) 
		g.ca.SetTile(11,1008)
		g.ca.SetTile(12,1007)
		g.ca.SetTile(13,1009)
		g.ca.SetTile(14,1010)
		g.ca.SetTile(15,1011)
		g.ca.SetTile(16,1010)
		g.ca.SetTile(17,1012)

		moattr := game.NewAttributes(3,3,3,3,3)
		moskill := game.NewSkills(0,0)
		moexp := &game.Expirience{0,0}
		monster := game.NewCharacter(moattr,moskill, moexp,fighter)
		point := game.NewPoint2D(350,200)
		momob := game.NewMob(point.GetPointIso(),&game.Rectangle2D{0,0,50, 20})
		g.mo = game.NewMonster(1013,momob,monster)
		g.loaded = true
	}
    // Write your game's logical update.
    
    g.al.RegisterBody(g.mo.GetBody())
    g.al.RegisterBody(g.ca.GetBody())
    h:= g.ca.Controll()
    g.al.RegisterAttack(h)
    g.ca.CollisionPoint()
	k:= g.ch.GetField(g.ca.ColPo)
    g.ca.Pathblock = g.th.Tiles[k].GetBlock()
    g.ddl.Update()
    g.ddl.GetDMGP(g.al.Update())
    return nil
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	
	g.ch.DrawLayer(0,screen,g.op)
	g.ch.DrawLayer(1,screen,g.op)
	g.mo.Draw(screen,g.op,g.th)
	g.ca.Draw(screen,g.op,g.th)
	
	g.ch.DrawObjects(screen,g.op)
	g.ch.DrawLayer(2,screen,g.op)
	g.ddl.Draw(screen,g.op)
	ebitenutil.DrawRect(screen,g.ca.ColPo.IX-5, g.ca.ColPo.IY-5,10,10,color.White)
	
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
    return 800, 600
}

func main() {
    game := &Game{}
    // Sepcify the window size as you like. Here, a doulbed size is specified.
    ebiten.SetWindowSize(800, 600)
    ebiten.SetWindowTitle("Your game's title")
    // Call ebiten.RunGame to start your game loop.
    if err := ebiten.RunGame(game); err != nil {
        log.Fatal(err)
    }
}

