package game

import(
"math"
"strconv"
)

var AttributeNames []string
var TraitNames []string
func init(){
 AttributeNames = []string{"Strength","Dexterity","Endurance","Lifeforce","Soulforce"}

TraitNames = []string{"Lifeenergy", "Magicenergy","Techenergy", "Attack", "Defense","Ability", "Tactic", "Luck", "Mobility", "Toughness"}
}
type Attributes struct{
	AAttribute map[string]int
}

func NewAttributes(strength, dexterity, endurance, lifeforce, soulforce int) *Attributes{
	atr := &Attributes{}
	atr.AAttribute = make(map[string]int)
	atr.AAttribute[AttributeNames[0]] = strength
	atr.AAttribute[AttributeNames[1]] = dexterity
	atr.AAttribute[AttributeNames[2]] = endurance
	atr.AAttribute[AttributeNames[3]] = lifeforce
	atr.AAttribute[AttributeNames[4]] = soulforce
	return atr
}

type Traits struct{
	ATrait map[string]int
	CAttributes map[string][]string
}


func NewTraits(LE,ME,TE,ATT,DEF,ABL,TAC,LUCK,MOBL,TOUGH int) *Traits{
	trt := &Traits{}
	trt.ATrait = make(map[string]int)
	trt.CAttributes = make(map[string][]string)
	trt.ATrait[TraitNames[0]] = LE
	trt.ATrait[TraitNames[1]] = ME
	trt.ATrait[TraitNames[2]] = TE
	trt.ATrait[TraitNames[3]] = ATT
	trt.ATrait[TraitNames[4]] = DEF
	trt.ATrait[TraitNames[5]] = ABL
	trt.ATrait[TraitNames[6]] = TAC
	trt.ATrait[TraitNames[7]] = LUCK
	trt.ATrait[TraitNames[8]]= MOBL
	trt.ATrait[TraitNames[9]]= TOUGH
	trt.CAttributes[TraitNames[0]] = []string{AttributeNames[3],AttributeNames[2], AttributeNames[0]}
	trt.CAttributes[TraitNames[1]] = []string{AttributeNames[4],AttributeNames[1], AttributeNames[0]}
	trt.CAttributes[TraitNames[2]] = []string{AttributeNames[2],AttributeNames[0], AttributeNames[4]}
	trt.CAttributes[TraitNames[3]] = []string{AttributeNames[0],AttributeNames[1], AttributeNames[3]}
	trt.CAttributes[TraitNames[4]] = []string{AttributeNames[1],AttributeNames[2], AttributeNames[3]}
	trt.CAttributes[TraitNames[5]] = []string{AttributeNames[1],AttributeNames[2], AttributeNames[4]}
	trt.CAttributes[TraitNames[6]] = []string{AttributeNames[4],AttributeNames[2], AttributeNames[3]}
	trt.CAttributes[TraitNames[7]] = []string{AttributeNames[2],AttributeNames[0], AttributeNames[4]}
	trt.CAttributes[TraitNames[8]]= []string{AttributeNames[3],AttributeNames[4], AttributeNames[1]}
	trt.CAttributes[TraitNames[9]]= []string{AttributeNames[0],AttributeNames[3], AttributeNames[1]}
	return trt
}


func (t *Traits)Calc(attr  *Attributes, class *Class){
	for i := range TraitNames{
		attrvalue :=float64(((attr.AAttribute[class.traitbonus.CAttributes[TraitNames[i]][0]])*3+(attr.AAttribute[class.traitbonus.CAttributes[TraitNames[i]][1]])*2+(attr.AAttribute[class.traitbonus.CAttributes[TraitNames[i]][2]])*1)/2)
	t.ATrait[TraitNames[i]] = int(float64(class.traitbonus.ATrait[TraitNames[i]]) + float64(class.Level)*class.levelbonus.TraitBonus[TraitNames[i]] + attrvalue)
	}
}

type ClassLevelTraits struct{
	TraitBonus map[string]float64
}
func NewClassLevelTraits(LE,ME,TE,ATT,DEF,ABL,TAC,LUCK,MOBL,TOUGH float64) *ClassLevelTraits{
	clt := &ClassLevelTraits{}
	clt.TraitBonus = make(map[string]float64)
	clt.TraitBonus["Lifeenergy"] = LE
	clt.TraitBonus["Magicenergy"] = ME
	clt.TraitBonus["Techenergy"] = TE
	clt.TraitBonus["Attack"] = ATT
	clt.TraitBonus["Defense"] = DEF
	clt.TraitBonus["Ability"] = ABL
	clt.TraitBonus["Tactic"] = TAC
	clt.TraitBonus["Luck"] = LUCK
	clt.TraitBonus["Mobility"]= MOBL
	clt.TraitBonus["Toughness"]= TOUGH
	return clt
}

type Expirience struct{
	ExpiriencePoints int
	Classpoints int
}

type Skills struct{
	Weapon int
	Armor int
}

func NewSkills(weap, armor int) *Skills{
	return &Skills{weap, armor}
}

type Class struct{
	traitbonus *Traits
	Level int
	levelbonus *ClassLevelTraits
}

func NewClass(traitboni *Traits, levelboni *ClassLevelTraits) *Class{
	return &Class{traitboni, 1, levelboni}
}

type Character struct{
	attributes *Attributes
	traits *Traits
	Traitbonus *Traits
	skills *Skills
	exp *Expirience
	class *Class
	damage int
	tech int
	collectedtech int
	magic int
}

func NewCharacter(attrib *Attributes, skills *Skills, exp *Expirience, class *Class) *Character{ 
	chartraits := NewTraits(0,0,0,0,0,0,0,0,0,0)
	traitboni :=NewTraits(0,0,0,0,0,0,0,0,0,0)
	chartraits.Calc(attrib, class)
	return &Character{attrib, chartraits, traitboni, skills, exp, class,0,0,0,0}
}

func (char *Character)RecalculateTraits(){
	char.traits.Calc(char.attributes,char.class)
}

func (char Character)GetLifepoints() int{
	return char.traits.ATrait["Lifeenergy"]- char.damage
}

func (char *Character)GetTechEnergy(energy int) int{
	char.tech += energy
	if char.tech > char.traits.ATrait["Techenergy"] {
		char.tech =  char.traits.ATrait["Techenergy"] 
	}
	if char.collectedtech%2 == 0{
	char.exp.ExpiriencePoints += energy/2
	char.exp.ExpiriencePoints += energy/10
    }
	return energy
}

func (char *Character)DealWithStrikeEvent(strike *Strikeevent, achar *Character) (int,bool){
	if char != achar{
		dmg:= int((float64(strike.Attack)*math.Exp((float64(strike.Ability+strike.Skill))/float64(char.traits.ATrait["Toughness"])))/((float64(char.traits.ATrait["Defense"])*math.Exp(float64(char.traits.ATrait["Mobility"]+char.skills.Armor)/float64(strike.Luck)))))
		char.damage+=dmg
		te := dmg/10
		if te == 0{
			te = 1
		}
		achar.GetTechEnergy(te)
		return dmg,true
	}
	return 0, false
}

func (char *Character)DealWithAttack(a *AnAttack) (*DPD,*TechValue){
	strike:= a.GetStrike()
	dmg:= int((float64(strike.Attack)*math.Exp((float64(strike.Ability+strike.Skill))/float64(char.traits.ATrait["Toughness"])))/((float64(char.traits.ATrait["Defense"])*math.Exp(float64(char.traits.ATrait["Mobility"]+char.skills.Armor)/float64(strike.Luck)))))
	char.damage+=dmg
	te := dmg/10
	if te == 0{
			te = 1
	}
	dmp := strconv.Itoa(dmg)
	ar := a.area.GetCenterIso()
	dpd := &DPD{ar,dmp,120,false}
	tv := &TechValue{te, a.char}
	return dpd,tv
	
}

func (char *Character)GetTechValue(tv *TechValue){
	char.tech += tv.tv
	char.collectedtech += tv.tv
	if char.tech > char.traits.ATrait["Techenergy"] {
		char.tech =  char.traits.ATrait["Techenergy"] 
	}
	if char.collectedtech%2 == 0{
	char.exp.ExpiriencePoints += char.collectedtech/2
	char.exp.ExpiriencePoints += char.collectedtech/10
    }
}

func (char Character)CreateStrikeEvent() (*Strikeevent, *Character){
	return NewStrikeEvent(char.traits.ATrait[TraitNames[3]], char.traits.ATrait[TraitNames[5]], char.skills.Weapon, char.traits.ATrait[TraitNames[7]]), &char
}

