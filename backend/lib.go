package backend

import(
  "eef.dungeoneer/backend/entities"
)
// This is only here so we can access the backend package  & it's subpackages from the frontend


// Have structs that will be sent to the frontend as data packets in an event
type CombatData struct{
  log string
  NewRoom RoomData
} 

type RoomData struct{
  Enemy entities.Monster
  RoomNum int
}


func CombatTurn(choice string) CombatData{
  return CombatData{
    log: "Test log!",
    NewRoom: RoomData{
      Enemy: *entities.CraftMonster("None", 0, 0, 0, "None"),
      RoomNum: -1,
    },
  };
}
