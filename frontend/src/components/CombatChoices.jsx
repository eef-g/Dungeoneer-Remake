import React from "react";
import ImageButton from "./ImageButton";
import buttonImage from "../assets/images/button-bg.png";

export default function CombatChoices(props) {

  const combat_style = {
    position: "absolute",
    bottom: "2%",
    display: "flex",
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "space-evenly",
    alignItems: "center",
    width: "100vw",
  };

  return (
    <div className="CombatChoices Display" style={combat_style}>
      <ImageButton
        backgroundImage={buttonImage}
        onClick={props.AttackFunction}
        text="Attack"
      />
      <ImageButton
        backgroundImage={buttonImage}
        onClick={props.BlockFunction}
        text="Block"
      />
      <ImageButton
        backgroundImage={buttonImage}
        onClick={props.HealFunction}
        text="Heal"
      />
      <ImageButton
        backgroundImage={buttonImage}
        onClick={props.RunFunction}
        text="Run"
      />
    </div>
  );
}