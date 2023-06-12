import React, { useState, useEffect } from 'react';
import './App.css'
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import Hoteles from './Componentes/Hoteles.jsx';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [hotels, setHotels] = useState([]);
    const [userId, setUserId] = useState(null);

    const handleLogin = (tipoUsuario, userId) => {
        setIsLoggedIn(true);
        setIsAdmin(tipoUsuario === 1);
        setUserId(userId);
    };

    return (
    <div className = 'App'>
        <div>
           <h1>ENCONTRA LA MEJOR OPCION</h1>
           {!isLoggedIn && <BotonLogin handleLogin={handleLogin} />}
           {!isLoggedIn && <BotonRegister />}
           {!isAdmin && hotels.map(hotel => (
          <Hoteles
            key={hotel.id}
            nombreHotel={hotel.name}
            piezas={hotel.cantHabitaciones}
            isLoggedIn={isLoggedIn}
            hotelId={hotel.id}
            userId={userId}
          />
        ))}
        </div>
    </div>
    )
}
export default App;
