import React from 'react'
import '../Stylesheet/Admin.css'
import 'Header.jsx'

function Admin() {
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