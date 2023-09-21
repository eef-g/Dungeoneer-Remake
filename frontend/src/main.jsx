import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'

const container = document.getElementById('root')

const root = createRoot(container)

root.render(
    // Have StrictMode disabled for now to make text displaying not a nightmare
    // Will re-enable it when the game is finished to be able to debug properly 

    //<React.StrictMode>
        <App/>
    //</React.StrictMode>
)
