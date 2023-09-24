package entities

type Player struct{
  Entity
}

func NewPlayer(n string, h int, d int) *Player{
  p := &Player{
    Entity{
      Name: n,
      Health: h,
      Damage: d,
    },
  }

  return p;
}
