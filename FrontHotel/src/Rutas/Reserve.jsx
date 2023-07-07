import React, { useState, useEffect } from "react"
import '../Stylesheet/Reserve.css'
import Header from '../Componentes/Header'
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

function Reserve() {
    const [fechaDesde, setFechaDesde] = useState(null);
    const [fechaHasta, setFechaHasta] = useState(null);
    const [hotelesDisponibles, setHotelesDisponibles] = useState([]);
    const [reservas, setReservas] = useState([]); // Agregar el estado de las reservas

    const clavetoken = "secreto";

    const handleFechaDesdeChange = (date) => {
        setFechaDesde(date);
    };

    const handleFechaHastaChange = (date) => {
        setFechaHasta(date);
    };

    const buscarHotelesDisponibles = () => {
        if (!fechaDesde || !fechaHasta) {
            alert("Debes completar los campos de fecha desde y fecha hasta.");
        } else if (fechaDesde >= fechaHasta) {
            alert("La fecha desde debe ser anterior a la fecha hasta.");
        } else {
            // Realizar la solicitud al backend para obtener los hoteles disponibles y las reservas del usuario
            fetch("/reservas/hotelsbyfecha", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${clavetoken}`
                },
                body: JSON.stringify({
                    fechaDesde,
                    fechaHasta,
                }),
            })
                .then((response) => response.json())
                .then((data) => {
                    // Actualizar el estado con los hoteles disponibles obtenidos del backend
                    setHotelesDisponibles(data.hoteles);
                    // Actualizar el estado con las reservas del usuario obtenidas del backend
                    setReservas(data.reservas);
                })
                .catch((error) => console.error(error));
        }
    };

    useEffect(() => {
        // Realizar la solicitud al backend para obtener las reservas del usuario
        fetch("/reservas", {
            headers: {
                Authorization: `Bearer ${clavetoken}`, // Asegúrate de incluir el token de autenticación
            },
        })
            .then((response) => response.json())
            .then((data) => {
                // Actualizar el estado con las reservas obtenidas del backend
                setReservas(data);
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    }, []);


    return (
        <div>
            <Header />
            <div>
                <br></br>
                <h2>INGRESE LAS FECHAS para su estadia</h2>
                <div>
                    <p>Desde: </p>
                    <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
                </div>
                <div>
                    <p>Hasta: </p>
                    <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
                </div>
                <button onClick={buscarHotelesDisponibles}>BUSCAR</button>
            </div>
            {hotelesDisponibles.length > 0 && (
                <div>
                    <h2>Hoteles Disponibles:</h2>
                    <ul>
                        {hotelesDisponibles.map((hotel) => (
                            <li key={hotel.id}>
                                <p>Nombre: {hotel.nombre}</p>
                                <p>Cantidad de Habitaciones: {hotel.cantidadHabitaciones}</p>
                            </li>
                        ))}
                    </ul>
                </div>
            )}

            {reservas.length > 0 && (
                <div>
                    <h2>Tus Reservas:</h2>
                    <ul>
                        {reservas.map((reserva) => (
                            <li key={reserva.id}>
                                <p>Nombre: {reserva.nombre}</p>
                                <p>Habitaciones: {reserva.cantidadHabitaciones}</p>
                                {/* Agrega más detalles de la reserva aquí */}
                            </li>
                        ))}
                    </ul>
                </div>
            )}
        </div>
    );
}

export default Reserve;
