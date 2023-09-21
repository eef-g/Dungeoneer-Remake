import React from "react";
import { useNavigate } from 'react-router-dom';
import { WipeDungeon, DungeonToString, GetRoomStats, ProgressInDungeon } from "../../wailsjs/go/main/App";
import Monster from "../components/Monster";
import CombatChoices from "../components/CombatChoices";
import StatusText from "../components/StatusText";

export default function Game() {
    const navigate  = useNavigate(); // Get the history object using the useHistory hook
    
    let [statusText, setStatusText] = React.useState("> You enter the dungeon!");
    let [monsterName, setMonsterName] = React.useState("");
    let [monsterImagePath, setMonsterImagePath] = React.useState("");
    let currentText = statusText;

    // This is run whenever the component is loaded
    React.useEffect(() => {
        DungeonToString().then((result) => {
            console.log(result);
        });
        GetRoomStats().then((result) => {
            UpdateCurrentRoomInfo(result);
        });
    }, []);

    // Delay function to call if I need to wait X seconds
    async function delay(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
    
    function UpdateCurrentRoomInfo(room_info){
        let room_enemy = room_info.Enemy;
        console.log(room_enemy);
        // End the game if there are no more enemies
        if (room_enemy.Name == "Finished") {
            // End the game
            WipeDungeon().then((result) => {
                navigate("/menu");
            });
            return;
        }
        setMonsterName(room_enemy.Name);
        setMonsterImagePath(room_enemy.Image);
    }

    function NextDungeonRoom() {
        // Chnge to different Go function
        let defeat_text = "\n> You defeat the " + monsterName + "!";
        UpdateText(defeat_text);
        ProgressInDungeon().then((result) => {
            GetRoomStats().then((final_result) => {
                UpdateCurrentRoomInfo(final_result);
            });
        });
    }

    function Run() {
        UpdateText("\n> You run away from the dungeon!");
        // Wait a few seconds then return to the main menu
        delay(2000).then(() => {
            WipeDungeon().then(() => {
                navigate("/menu");
            });
        });
    }

    function Attack() {
    }

    async function UpdateText(newText) {
        let words = newText.split(" ");
        let i = 1;
        console.log("Current text before intervalID: '" + currentText + "'")
        let estimatedText = currentText + words[0] + " ";
        setStatusText(estimatedText);
        for(i; i < words.length; i++) {
            await delay(25).then(() => {
                console.log(estimatedText);
                setStatusText(estimatedText + words[i] + " ");
                estimatedText += words[i] + " ";
            });
        }
    }

    React.useEffect(() => {
        currentText = statusText;
    }, [statusText]);
 

    return (
        // Eventually, replace the main <div> w/ a <BackgroundImage> component
        <div>
            <Monster image={monsterImagePath} name={monsterName}/>
            <StatusText text={statusText}/>
            <CombatChoices 
                AttackFunction={NextDungeonRoom}
                BlockFunction={NextDungeonRoom}
                HealFunction={NextDungeonRoom}
                RunFunction={Run}
            />
        </div>
    );
}