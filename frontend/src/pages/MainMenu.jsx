import React from 'react';
import { useNavigate } from 'react-router-dom';
import { GenerateDungeon } from '../../wailsjs/go/main/App';
import '../style.css';
import '../components/ImageButton.jsx';
import '../components/BackgroundImage.jsx';
import '../assets/images/main-menu.png'
import '../assets/images/button-bg.png'


import BackgroundImage from '../components/BackgroundImage.jsx';
import mainMenuImage from "../assets/images/main-menu.png" // Tell webpack this JS file uses this image
import buttonImage from "../assets/images/button-bg.png"
import ImageButton from '../components/ImageButton.jsx';


export default function MainMenu() {
    const navigate  = useNavigate(); // Get the history object using the useHistory hook
    let [dungeon, setDungeon] = React.useState("Dungeon Info Goes Here"); // Create a state variable for the dungeon
    
    const appStyle = {
        display: 'flex', // Set display to flex
        flexDirection: 'column', // Set main axis of flex to column
        justifyContent: 'center', // Center horizontally
        alignItems: 'center', // Center vertically
        height: '100vh', // Set height to 100vh
    };

    const titleStyle = {
        fontFamily: 'Runescape',
        fontSize: '70px',
        color: 'yellow',
        position: 'absolute',
        top: '15%',
    }

    function startGame() {
        GenerateDungeon().then((result) => {
            console.log(result);
            setDungeon(result);
            //navigate("/game");
          });
    }

    return (
        <BackgroundImage backgroundImage={mainMenuImage}>
            <div id="main-menu" style={appStyle}>
                <div id="title" className="title" style={titleStyle}>Dungeoneer</div>
                <div id="input" className="input-box">
                    <ImageButton backgroundImage={buttonImage} onClick={startGame} text="Begin"/>
                    <p id="dungeon-info">{dungeon}</p>
                </div>
            </div>
        </BackgroundImage>
    )
}