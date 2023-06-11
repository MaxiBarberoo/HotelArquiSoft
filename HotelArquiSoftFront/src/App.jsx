import React, { useState } from 'react';
import './App.css'
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import Hoteles from './Componentes/Hoteles.jsx'

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);

    const handleLogin = () => {
        setIsLoggedIn(true);
    };

    return (
    <div className = 'App'>
        <div>
           <h1>ENCONTRA LA MEJOR OPCION</h1>
           <BotonLogin handleLogin={handleLogin}/>
           <BotonRegister />
           <Hoteles
                nombreHotel='HOTEL LUXURY'
                image='1'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='HOTEL PALACE'
                image='2'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='HOTEL CENTRAL PARK'
                image='3'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='HOTEL PARAÍSO'
                image='4'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='CASA DE LAS MONTAÑAS'
                image='5'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='HOTEL EXOTIC'
                image='6'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='HOTEL MANSION DEL SUR'
                image='7'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='HOTEL JARDIN'
                image='8'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='GRAN HOTEL'
                image='9'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
            <Hoteles
                nombreHotel='CASA NATURAL'
                image='10'
                piezas= '5'
                isLoggedIn = {isLoggedIn}
            />
        </div>
    </div>
    )
}
export default App;
