package main

import (
	"fmt"
)

type Player struct {
	name   string
	age    int
	Hp     int
	Jineng map[string]int
	DaoJu  []string
}

func NewPlayer(name string, age int, Hp int, Jineng map[string]int, DaoJu []string) *Player {
	return &Player{
		name:   name,
		age:    age,
		Hp:     Hp,
		Jineng: Jineng,
		DaoJu:  DaoJu,
	}
}

func (p *Player) Gongji(Jneng string) {
	var key string
	var value int
	for key, value = range p.Jineng {
		if key == Jneng {
			fmt.Printf("%s使用%s技能之后，打掉了%d的血，打掉的是谁呢？", p.name, key, value)
			break
		} else {
			fmt.Println("不存在该技能")
		}
	}
}

func (p *Player) Gongjied(Jinenged int) {
	fmt.Printf("%s被攻击啦\n", p.name)
	p.Hp = p.Hp - Jinenged //如果不是指针接收器的话，数据值拷贝，会另启一块地址，不会对源地址进行改变
}

func (p *Player) BuyDaoju(Daoju string) {
	fmt.Printf("%s买了%s到具", p.name, Daoju)
	p.DaoJu = append(p.DaoJu, Daoju)
}

func main() {
	var J string
	JN := map[string]int{
		"一刀斩": 10,
		"双刀斩": 20,
	}
	fmt.Println("请输入你要使用的技能")
	fmt.Scanf("%s", &J)
	player_1 := Player{"Npc1", 60, 100, JN, nil}
	player_1.Gongji(J)
	player_2 := Player{"Npc2", 90, 100, JN, nil}
	player_2.Gongjied(player_1.Jineng[J]) //传值之后在接收器中，p = &player_2
	fmt.Println(player_2.Hp)
	player_2.BuyDaoju("达摩圈套")
	fmt.Printf("道具栏中有%s", player_2.DaoJu)
}
