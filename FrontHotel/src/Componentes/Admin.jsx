import React, { useState, useEffect } from 'react'

function Admin() {
    const [reservasTotales, setReservasTotales] = useState([]);
    return (
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