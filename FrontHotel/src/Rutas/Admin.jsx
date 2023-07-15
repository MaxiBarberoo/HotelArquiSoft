import React, { useState, useEffect } from 'react';

function Admin(){
    const [reservasTotales, setReservasTotales] = useState([]);
    const [reservasUsuario, setReservasUsuario] = useState([]);
    const [nuevoHotel, setNuevoHotel] = useState({
        nombre: '',
        cantidadHabitaciones: 0
    });

    const [hoteles, setHoteles] = useState([]);
    const [isLoggedIn, setIsLoggedIn] = useState(false);
    const [isAdmin, setIsAdmin] = useState(false);
    const [userId, setUserId] = useState('');

    useEffect(() => {
        fetchReservas();
    }, []);

    const fetchReservas = () => {
        getReservas()
            .then((reservas) => {
                setReservasTotales(reservas);
            })
            .catch((error) => {
                console.error(error);
            });
    };

    const handleNuevoHotelChange = (event) => {
        const { name, value } = event.target;
        setNuevoHotel((prevNuevoHotel) => ({
            ...prevNuevoHotel,
            [name]: value
        }));
    };

    const crearNuevoHotel = () => {
        const cantHabitaciones = parseInt(nuevoHotel.cantidadHabitaciones);
        fetch('http://localhost:8090/hotels', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                name: nuevoHotel.nombre,
                cantHabitaciones: cantHabitaciones
            })
        })
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.error(data.error);
                } else {
                    setHoteles((prevHoteles) => [...prevHoteles, data]);
                    alert('Se ha creado un nuevo hotel con éxito');
                    setNuevoHotel({
                        nombre: '',
                        cantidadHabitaciones: 0
                    });
                }
            })
            .catch((error) => console.error(error));
    };

    const handleLogin = (tipoUsuario, userId) => {
        setIsLoggedIn(true);
        setIsAdmin(tipoUsuario === 1);
        setUserId(userId);

    };



    return(
        <div>

            <h2>Reservas Totales</h2>
            <ul>
                {reservasTotales.map((reserva) => (
                    <li key={reserva.id}>{reserva.nombre}</li>
                ))}
            </ul>


        </div>
    );
}

export default Admin;