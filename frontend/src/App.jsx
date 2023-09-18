import {useState} from 'react';
import './BackgroundImage.jsx';
import './ImageButton.jsx';
import './App.css';
import './style.css';
import {Greet} from "../wailsjs/go/main/App";
import BackgroundImage from './BackgroundImage.jsx';
import React from 'react';
import mainMenuImage from "./assets/images/main-menu.png" // Tell webpack this JS file uses this image
import buttonImage from "./assets/images/button-bg.png"
import ImageButton from './ImageButton.jsx';


function App() {
    const [resultText, setResultText] = useState("Press the button to generate room length.");
    const updateResultText = (result) => setResultText(result);

    const appStyle = {
        display: 'flex', // Set display to flex
        flexDirection: 'column', // Set main axis of flex to column
        justifyContent: 'center', // Center horizontally
        alignItems: 'center', // Center vertically
        height: '100vh', // Set height to 100vh
    };

    const resultStyle = {
        fontFamily: 'Runescape',
        fontSize: '20px',
        color: 'yellow',
    };

    const titleStyle = {
        fontFamily: 'Runescape',
        fontSize: '70px',
        color: 'yellow',
        position: 'absolute',
        top: '15%',
    }

    function greet() {
        Greet("{USER}").then(updateResultText);
    }

    return (
        <BackgroundImage backgroundImage={mainMenuImage}>
             <div id="App" style={appStyle}>
                <div id="title" className="title" style={titleStyle}>Dungeoneering</div>
                <div id="result" className="result" style={resultStyle}>{resultText}</div>
                <div id="input" className="input-box">
                    <ImageButton backgroundImage={buttonImage} onClick={greet} text="Begin"/>
                </div>
            </div>
        </BackgroundImage>
    )
}

export default App
