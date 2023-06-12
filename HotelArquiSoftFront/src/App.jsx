import React, { useState, useEffect } from 'react';
import './App.css'
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import Hoteles from './Componentes/Hoteles.jsx'

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [hotels, setHotels] = useState([]);

    const handleLogin = (tipoUsuario) => {
        setIsLoggedIn(true);
        setIsAdmin(tipoUsuario === 1);
    };

    useEffect(() => {
        // Realizar la solicitud GET para obtener la lista de hoteles
        fetch('http://localhost:8090/hotels')
          .then(response => response.json())
          .then(data => setHotels(data))
          .catch(error => console.error(error));
      }, []);

    return (
    <div className = 'App'>
        <div>
           <h1>ENCONTRA LA MEJOR OPCION</h1>
           {!isLoggedIn && <BotonLogin handleLogin={handleLogin} />}
           {!isLoggedIn && <BotonRegister />}
           {hotels.map(hotel => (
          <Hoteles
            key={hotel.id}
            nombreHotel={hotel.name}
            piezas={hotel.cantHabitaciones}
            isLoggedIn={isLoggedIn}
            hotelId={hotel.id}
            userId={hotel.userId}
          />
        ))}
        </div>
    </div>
    )
}
export default App;
