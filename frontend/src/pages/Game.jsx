import React from "react";
import { DungeonToString, GetRoomStats, ProgressInDungeon } from "../../wailsjs/go/main/App";
import ImageButton from "../components/ImageButton";
import buttonImage from "../assets/images/button-bg.png"

export default function Game() {
    let [monsterName, setMonsterName] = React.useState("");
    let [monsterHealth, setMonsterHealth] = React.useState(0);
    let [monsterAttack, setMonsterAttack] = React.useState(0);
    let [monsterAttackBonus, setMonsterAttackBonus] = React.useState(0);
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
        setMonsterName(room_enemy.Name);
        setMonsterHealth(room_enemy.HP);
        setMonsterAttack(room_enemy.Attack);
        setMonsterAttackBonus(room_enemy.AttackBonus);
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

    return (
        // Eventually, replace the main <div> w/ a <BackgroundImage> component
        <div>
            <h1>Game</h1>
            <ImageButton backgroundImage={buttonImage} onClick={NextDungeonRoom} text="Next Room"/>
            <p>Current room stats:</p>
            <ul>
                <li>Current Monster: {monsterName}</li>
                <li>Current Monster Health: {monsterHealth}</li>
                <li>Current Monster Attack: {monsterAttack}</li>
                <li>Current Monster Attack Bonus: {monsterAttackBonus}</li>
                <li>Current Monster Image Path: <p>{monsterImagePath}</p></li>
            </ul>
        </div>
    );
}