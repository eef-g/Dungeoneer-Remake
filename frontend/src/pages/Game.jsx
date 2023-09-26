import React from "react";
import { useNavigate } from 'react-router-dom';
import { CombatTurn, AddTextToQueue, CheckTextQueue, WipeDungeon, DungeonToString, GetRoomStats} from "../../wailsjs/go/main/App";
import Monster from "../components/Monster";
import CombatChoices from "../components/CombatChoices";
import StatusText from "../components/StatusText";
import BackgroundImage from "../components/BackgroundImage";
import combatBg from "../assets/images/combat-bgv2.png" // Tell webpack this JS file uses this image

export default function Game() {
    const navigate  = useNavigate(); // Get the history object using the useHistory hook
    let [isDisabled, setIsDisabled] = React.useState(false); // Disable the buttons when the user clicks one
    let [statusText, setStatusText] = React.useState("");
    let [monsterName, setMonsterName] = React.useState("");
    let [monsterImagePath, setMonsterImagePath] = React.useState("");
    
    // This is run whenever the component is loaded
    React.useEffect(() => {
        AddTextToQueue("\n> You enter the dungeon!").then(() => {
            DungeonToString().then((result) => {
                console.log(result);
                GetRoomStats().then(async (result) => {
                    await UpdateCurrentRoomInfo(result);
                });
            });
        });

        // Start the loop that will listen for new messages
        // This is the main loop that will listen for new messages from the Go backend
        let estimatedText = statusText; // Have to have this variable to keep track of the current text
        const interval = setInterval(() => {
            CheckTextQueue().then(async (result) => {
                if (result != "") {
                    await UpdateText(result, estimatedText);
                    estimatedText += result;
                }
            });
        }, 250);

        return () => clearInterval(interval);
    }, []);

    // Delay function to call if I need to wait X milliseconds
    async function delay(ms) {
        return new Promise(resolve => setTimeout(resolve, ms));
    }
 
    // Move this logic to the backend eventually
    async function UpdateCurrentRoomInfo(room_info){
        let room_enemy = room_info.Enemy;
        
        // End the game if there are no more enemies
        if (room_enemy.Base.Name == "Finished") {
            // End the game
            await AddTextToQueue("\n> You have defeated all the enemies in the dungeon!").then(async () => {
                setIsDisabled(true);
                delay(1500).then(() => {
                    WipeDungeon().then(() => {
                        navigate("/menu");
                        setIsDisabled(false);
                    });   
                });
            });
            return;
        }
        setMonsterName(room_enemy.Base.Name);
        setMonsterImagePath(room_enemy.Image);
        await AddTextToQueue("\n> You are stopped by a " + room_enemy.Base.Name + "!");
    }

    async function Run() {
        setIsDisabled(true);
        AddTextToQueue("\n> You run away from the dungeon!").then(() => {;
            // Wait a few seconds then return to the main menu
            delay(2000).then(() => {
                WipeDungeon().then(() => {
                    navigate("/menu");
                    setIsDisabled(false);
                });
            });
        });
    }

    async function UpdateText(newText, currentGuess) {
        let words = newText.split(" ");
        let i = 1;
        let estimatedText = currentGuess + words[0] + " ";
        setStatusText(estimatedText);
        for(i; i < words.length; i++) {
            await delay(25).then(() => {
                setStatusText(estimatedText + words[i] + " ");
                estimatedText += words[i] + " ";
            });
        }
    }

    React.useEffect(() => {
        console.log("isDisabled: " + isDisabled);
    }, [isDisabled]);

    async function Attack() {
      await CombatChoice("attack").then((result) =>{
        if(result.NextRoom) {
          AddTextToQueue("\n> You defeat the " + monsterName + "!");
          GetRoomStats().then((room_result) => {
            UpdateCurrentRoomInfo(room_result);
          });
        }
      });
      await delay(1000);
      setIsDisabled(false);
    }

    async function CombatChoice(choice) {
      setIsDisabled(true);
      let outcome_packet = await CombatTurn(choice);
      return outcome_packet;
    }

    return (
        // Eventually, replace the main <div> w/ a <BackgroundImage> component
        <BackgroundImage backgroundImage={combatBg}>
            <Monster image={monsterImagePath} name={monsterName}/>
            <StatusText text={statusText}/>
            <CombatChoices
                disabled={isDisabled} 
                AttackFunction={Attack}
                BlockFunction={Run}
                HealFunction={Run}
                RunFunction={Run}
            />
        </BackgroundImage>
    );
}
