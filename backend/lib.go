package backend

import (
  "fmt"
	"eef.dungeoneer/backend/entities"
  "eef.dungeoneer/backend/dungeons"
)

// Have structs that will be sent to the frontend as data packets in an event
type CombatData struct{
  log string
  NewRoom dungeons.Room 
} 

func CombatTurn(player *entities.Player, dungeon *dungeons.Dungeon, choice string) CombatData{
  // Make output_packet an empty CombatData packet
  output_packet := CombatData{
    log: "", 
    NewRoom: dungeons.Room{
      Enemy: entities.CraftMonster("None", -1, -1, -1, "None"),
      NextRoom: nil,
    },
  };

  monster := dungeon.CurrentRoom.Enemy;

  switch choice{
    case "attack":
      start_health := monster.Base.Health;
      monster.Base.Health = entities.Attack(&player.Entity, &monster.Base);
      output_packet.log += fmt.Sprintf("\n> %v does %d damage to %v", player.Name, start_health - monster.Base.Health, monster.Base.Name);
      // Check and make sure we aren't moving to the next room
      if (monster.Base.Health <= 0) {
        output_packet.log += fmt.Sprintf("\n> You defeated the %v!", monster.Base.Name);
        dungeon.ProgressInDungeon();
        output_packet.NewRoom = *dungeon.CurrentRoom;
      }

    case "block":
      output_packet.log += "You chose block. Loser :P";
    
    case "heal":
      output_packet.log += "You decided to heal. What a gamer.";

    case "run":
      output_packet.log += "You ran away like a little baby. Waaaah";

    default:
      output_packet.log += fmt.Sprintf("ERROR: '%v' is an invalid choice.", choice);
  }


  return output_packet;
}
