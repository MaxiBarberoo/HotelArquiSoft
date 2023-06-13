import React from 'react';

function Reservas({ reservas, reservasTotales }) {
    console.log(reservas);
    console.log(reservasTotales);

    return (
        <div>
            <h2>Reservas</h2>
            {reservas && reservas.length > 0 ? (
                <div>
                    <ul>
                        {reservas.map((reserva) => (
                            <li key={reserva.id}>
                                <p>Fecha de ingreso: {reserva.fecha_ingreso}</p>
                                <p>Fecha de egreso: {reserva.fecha_egreso}</p>
                                <p>Hotel: {reserva.hotel_id}</p>
                            </li>
                        ))}
                    </ul>
                </div>
            ) : (
                <p>No tienes reservas realizadas.</p>
            )}

            {reservasTotales && reservasTotales.length > 0 && (
                <div>
                    <h3>Todas las reservas</h3>
                    <ul>
                        {reservasTotales.map((reserva) => {
                            console.log(reserva); // Verificar los datos de cada reserva en la consola
                            return (
                                <li key={reserva.id}>
                                    <p>Usuario: {reserva.user_id}</p>
                                    <p>Fecha de ingreso: {reserva.fecha_ingreso}</p>
                                    <p>Fecha de egreso: {reserva.fecha_egreso}</p>
                                    <p>Hotel: {reserva.hotel_id}</p>
                                </li>
                            );
                        })}
                    </ul>
                </div>
            )}
        </div>
    );
}
export default Reservas;

