import React, { useState, useEffect } from 'react';
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import Hoteles from './Componentes/Hoteles.jsx';

function App() {
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [hoteles, setHoteles] = useState([]);
    const [userId, setUserId] = useState(null);

    const handleLogin = (tipoUsuario, userId) => {
        setIsLoggedIn(true);
        setIsAdmin(tipoUsuario === 1);
        setUserId(userId);
    };

    useEffect(() => {
        fetch('http://localhost:8090/hotels')
            .then(response => response.json())
            .then(data => setHoteles(data))
            .catch(error => console.error(error));
    }, []);

    return (
        <div className="App">
            <div>
                <h1>ENCONTRA LA MEJOR OPCION</h1>
                {!isLoggedIn && <BotonLogin handleLogin={handleLogin} />}
                {!isLoggedIn && <BotonRegister />}
                {!isAdmin && hoteles.map(hotel => (
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
    );
}

export default App;