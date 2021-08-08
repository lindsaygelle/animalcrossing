package villager

type Villagers map[Key]Villager

func (v Villagers) Add(villager Villager) {
	v[villager.Key] = villager
}

func NewVillagers(v ...Villager) Villagers {
	villagers := Villagers{}
	for _, v := range v {
		villagers.Add(v)
	}
	return villagers
}
