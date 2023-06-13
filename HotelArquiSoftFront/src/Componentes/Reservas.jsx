import React, { useEffect, useState } from "react";

function Reservas(props) {
  const { reservas } = props;

  return (
    <div>
      <h2>Reservas</h2>
      {reservas.length > 0 ? (
        <ul>
          {reservas.map((reserva) => (
            <li key={reserva.id}>
              <p>Fecha de ingreso: {reserva.fecha_ingreso}</p>
              <p>Fecha de egreso: {reserva.fecha_egreso}</p>
              <p>Hotel: {reserva.hotel_nombre}</p>
            </li>
          ))}
        </ul>
      ) : (
        <p>No tienes reservas realizadas.</p>
      )}
    </div>
  );
}

export default Reservas;