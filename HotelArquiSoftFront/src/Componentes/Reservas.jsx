import React from 'react';

function Reservas(props) {
  const { reservas, reservasTotales } = props;

  return (
    <div>
      <h2>Reservas</h2>
      {reservas.length > 0 ? (
        <ul>
          {reservas.map((reserva) => (
            <li key={reserva.id}>
              <p>Fecha de ingreso: {reserva.fecha_ingreso}</p>
              <p>Fecha de egreso: {reserva.fecha_egreso}</p>
              <p>Hotel: {reserva.hotel_id}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>No tienes reservas realizadas.</p>
      )}

      {reservasTotales.length > 0 && (
        <div>
          <h3>Todas las reservas</h3>
          <ul>
            {reservasTotales.map((reserva) => (
              <li key={reserva.id}>
                <p>Usuario: {reserva.user_id}</p>
                <p>Fecha de ingreso: {reserva.fecha_ingreso}</p>
                <p>Fecha de egreso: {reserva.fecha_egreso}</p>
                <p>Hotel: {reserva.hotel_id}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
}

export default Reservas;
