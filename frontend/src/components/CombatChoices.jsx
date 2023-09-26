import React from "react";
import ImageButton from "./ImageButton";
import buttonImage from "../assets/images/button-bg.png";
import cooldownButtonImage from "../assets/images/disabled-button-bg.png";
// ^^^ Need to make a greyed out button image that lets us show when the button is on cooldown

export default function CombatChoices(props) {

  const combat_style = {
    position: "absolute",
    bottom: "2%",
    left: "3%",
    display: "flex",
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "space-evenly",
    alignItems: "center",
    width: "94vw",
  };

  return (
    <div className="CombatChoices Display" style={combat_style}>
      <ImageButton
        backgroundImage={props.disabled ? cooldownButtonImage : buttonImage}
        //backgroundImage={buttonImage}
        onClick={props.disabled ? null: props.AttackFunction}
        text="Attack"
      />
      <ImageButton
        backgroundImage={props.disabled ? cooldownButtonImage : buttonImage}
        //backgroundImage={buttonImage}
        onClick={props.disabled ? null : props.BlockFunction}
        text="Block"
      />
      <ImageButton
        backgroundImage={props.disabled ? cooldownButtonImage : buttonImage}
        //backgroundImage={buttonImage}
        onClick={props.disabled ? null : props.HealFunction}
        text="Heal"
      />
      <ImageButton
        backgroundImage={props.disabled ? cooldownButtonImage : buttonImage}
        //backgroundImage={buttonImage}
        onClick={props.disabled ? null : props.RunFunction}
        text="Run"
      />
    </div>
  );
}
