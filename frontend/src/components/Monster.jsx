import React from "react";

export default function Monster(props) {
    return (

        <div className="Monster Display">
            <img src={props.image} alt={props.name} />
        </div>
    );
}