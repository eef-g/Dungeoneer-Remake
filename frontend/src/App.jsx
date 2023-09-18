import * as React from 'react';
import {BrowserRouter, Routes, Route} from 'react-router-dom';
import MainMenu from './pages/MainMenu.jsx';
import Game from './pages/Game.jsx';
import Results from './pages/Results.jsx';
import './App.css';

export default function App() {
    return (
        <div>
            <BrowserRouter>
                <Routes>
                    <Route index element={<MainMenu/>}/>
                    <Route element={<MainMenu/>} path="/menu"/>
                    <Route element={<Game/>} path="/game"/>
                    <Route element={<Results/>} path="/results"/>

                    // 404 Not Found -- just in case
                    <Route path="*">"404 Not Found!"</Route>
                </Routes>
            </BrowserRouter>
        </div>
    );
}