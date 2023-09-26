import React from 'react';
import "../style.css";

export default function StatusText(props) {
    const textBoxRef = React.useRef(null);
    
    const textBox_style = {
        width: "90vw",
        height: "30vh",
        backgroundColor: "#696969",
        border: "solid",
        borderColor: "#cfd700",
        borderRadius: "2px",
        overflow: "auto",
        position: "absolute",
        bottom: "25%",
        left: "4%",
        textAlign: "left",
        color: "yellow",
        fontFamily: "Runescape",
        fontSize: "18px",
    }

    const p_style = {
        margin: "0",
        padding: "0",
        lineHeight: "1.2" // Adjust this value to make the lines closer together
    }

    React.useEffect(() => {
        // Scroll to the bottom of the div when the text is updated
        textBoxRef.current.scrollTop = textBoxRef.current.scrollHeight;
    }, [props.text]);

    const lines = props.text.split("\n");
    return (
        <div ref={textBoxRef} style={textBox_style} className="status-text">
            {lines.map((line, index) => (
                <p key={index} style={p_style}>{line}</p>
            ))}
        </div>
    );
}
