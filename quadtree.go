package game

import(
"errors"
"github.com/hajimehoshi/ebiten"
"log"
)

//Used to hold an object
type Node struct{
	Area *Rectangle2D
	childNE, childNW, childSE, childSW *Node
	Anobject *Object
	level int
}

func NewNode(area *Rectangle2D, lvl int) *Node {
	n := &Node{}
	n.Area = area
	n.level = lvl
	return n
}

//TestBranchesNil looks if Branches of a node are empty
func (n Node)TestBranchesNil() bool{
	return n.childNE == nil && n.childNW == nil && n.childSE == nil && n.childSW == nil
}

//TestNodeNil looks if Branches and object of a Node are nil
func (n *Node)TestNodeNil() bool{
	return n.TestBranchesNil() && n.Anobject == nil
}

//Finds out which of the Branches an Object fits in
func (n Node)TestQuad(pos *PointIso) (*Node, error){
	pos2D := pos.GetPoint2D()
	if n.Area.L + n.Area.GetLength()/2 > pos2D.X{
		if n.Area.U + n.Area.GetLength()/2 > pos2D.Y{
			return n.childNE,nil
		} else{
			return n.childSE,nil
		}
	} else {
		if n.Area.U+n.Area.GetLength()/2 >pos2D.Y{
			return n.childNW,nil
		} else{
			return n.childSW,nil
		}
	}
	err := errors.New("Not within TreeArea")
	return nil, err
}

//Appends new Branches to a Node
func (n *Node)AddALevel(){
	halflength := n.Area.GetLength()/2
	n.childNE = NewNode(&Rectangle2D{n.Area.L, n.Area.U, n.Area.L+halflength, n.Area.U+halflength}, n.level+1)
	n.childNW = NewNode(&Rectangle2D{n.Area.R-halflength, n.Area.U, n.Area.R, n.Area.U+halflength}, n.level+1)
	n.childSE = NewNode(&Rectangle2D{n.Area.L,n.Area.D-halflength, n.Area.L+halflength, n.Area.D},n.level+1)
	n.childSW = NewNode(&Rectangle2D{n.Area.R-halflength,n.Area.D-halflength, n.Area.R, n.Area.D},n.level+1)
}

//Appends an object to the tree
func (n *Node)AppendObject(obj *Object) *Node{
	if n.Anobject == nil{
		if n.TestBranchesNil(){
			log.Println("Appended")
			n.Anobject = obj
			return n
		}
		if !n.TestBranchesNil(){
			node,_ := n.TestQuad(obj.Position)
			node.AppendObject(obj)
			return node
		}
	}
	if n.Anobject != nil {
		if !n.TestBranchesNil(){
			node,_:=n.TestQuad(obj.Position)
			node.AppendObject(obj)
			ob2 := n.Anobject
			n.Anobject = nil
			node2,_:= n.TestQuad(ob2.Position)
			node2.AppendObject(ob2)
			return node
		} else if n.TestBranchesNil() {
			n.AddALevel()
			node,_:=n.TestQuad(obj.Position)
			node.AppendObject(obj)
			ob2 := n.Anobject
			n.Anobject = nil
			node2,_:= n.TestQuad(ob2.Position)
			node2.AppendObject(ob2)
			return node2
		}
	}
	return n
}

//GraphicHolder intermediate Structure that holds ScreenPosition and tilenumber of an object
type GraphicHolder struct{
	tilenum int
	position *PointIso
}

//Collects Data for intermediate Structure of all objects in the tree
func (n *Node)ReturnGraphic(temp []GraphicHolder) []GraphicHolder{
	if n.TestNodeNil(){
		return temp
	}
	if n.Anobject != nil{
		gH := GraphicHolder{n.Anobject.Graphic, n.Anobject.Position}
		temp = append(temp, gH)
		return temp
	}
	temp = n.childNE.ReturnGraphic(temp)
	temp =n.childNW.ReturnGraphic(temp)
	temp =n.childSE.ReturnGraphic(temp)
	temp = n.childSW.ReturnGraphic(temp)
	return temp
}

//The main quadtree class
type QuadTree struct{
	Initial *Node
}

func NewQuadTree(area *Rectangle2D) *QuadTree{
	init :=NewNode(area,0)
	return &QuadTree{init}
}

//Appends an object to the tree
func (qt *QuadTree)AppendObject(obj *Object) {
	if qt.Initial.Area.PointIsoWithin(obj.Position){
		qt.Initial.AppendObject(obj)
	}
}
func (qt *QuadTree)ReturnGraphic() []GraphicHolder{
	var holder []GraphicHolder
	holder = qt.Initial.ReturnGraphic(holder)
	return holder
}

func (qt *QuadTree)Draw(screen *ebiten.Image, options *ebiten.DrawImageOptions, th *TileHolder){
	arr := qt.ReturnGraphic()
	for i:= range arr{
		options.GeoM.Translate(arr[i].position.IX,arr[i].position.IY - float64(th.Tiles[arr[i].tilenum].Off))
		screen.DrawImage(th.DisplayTile(arr[i].tilenum), options)
		options.GeoM.Translate(-arr[i].position.IX,-arr[i].position.IY + float64(th.Tiles[arr[i].tilenum].Off))
	}
}

