import React from "react";
import { useNavigate } from 'react-router-dom';
import { WipeDungeon, DungeonToString, GetRoomStats, ProgressInDungeon } from "../../wailsjs/go/main/App";
import Monster from "../components/Monster";
import CombatChoices from "../components/CombatChoices";

export default function Game() {
    const navigate  = useNavigate(); // Get the history object using the useHistory hook

    let [monsterName, setMonsterName] = React.useState("");
    // let [monsterHealth, setMonsterHealth] = React.useState(0);
    // let [monsterAttack, setMonsterAttack] = React.useState(0);
    // let [monsterAttackBonus, setMonsterAttackBonus] = React.useState(0);
    let [monsterImagePath, setMonsterImagePath] = React.useState("");

    // This is run whenever the component is loaded
    React.useEffect(() => {
        DungeonToString().then((result) => {
            console.log(result);
        });
        GetRoomStats().then((result) => {
            UpdateCurrentRoomInfo(result);
        });
    }, []);
    
    function UpdateCurrentRoomInfo(room_info){
        let room_enemy = room_info.Enemy;
        // End the game if there are no more enemies
        if (room_enemy.Name == "Finished") {
            // End the game
            WipeDungeon().then((result) => {
                navigate("/menu");
            });
            return;
        }

        setMonsterName(room_enemy.Name);
        // setMonsterHealth(room_enemy.HP);
        // setMonsterAttack(room_enemy.Attack);
        // setMonsterAttackBonus(room_enemy.AttackBonus);
        setMonsterImagePath(room_enemy.Image);

    }

    function NextDungeonRoom() {
        // Chnge to different Go function
        ProgressInDungeon().then((result) => {
            GetRoomStats().then((final_result) => {
                UpdateCurrentRoomInfo(final_result);
            });
        });
    }

    function Run() {
        console.log("Run");
        WipeDungeon().then((result) => {
            navigate("/menu");
        });
    }


    function Attack() {
        console.log("Attack");
    }

    return (
        // Eventually, replace the main <div> w/ a <BackgroundImage> component
        <div>
            <Monster image={monsterImagePath} name={monsterName}/>
            <CombatChoices 
                AttackFunction={NextDungeonRoom}
                BlockFunction={NextDungeonRoom}
                HealFunction={NextDungeonRoom}
                RunFunction={Run}
            />
        </div>
    );
}