package entities

type Entity struct {
  Name string
  Damage int
  Health int
}

func Attack(attacker *Entity, defender *Entity) int{
  // Make the attacker deal damage to the defender
  result := defender.Health - attacker.Damage;
  return result;
}
