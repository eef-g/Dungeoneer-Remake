import React from "react";

export default function Monster(props) {
    
    const monster_style = {
      position: 'absolute',
      left: '26%',
      top: '6%',
    };

    return (
        <div style={monster_style} className="Monster Display">
            <img src={props.image} alt={props.name} />
        </div>
    );
}
