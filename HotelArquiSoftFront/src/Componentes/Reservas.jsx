import React, { useState, useEffect } from 'react';

function Reservas(props) {
  const [reservas, setReservas] = useState([]);

  useEffect(() => {
    obtenerReservas();
  }, []);

  const obtenerReservas = () => {
    fetch(`http://localhost:8090/reservas/reservauser/${props.userId}`)
      .then((response) => response.json())
      .then((data) => {
        setReservas(data);
      })
      .catch((error) => {
        console.error('Error:', error);
      });
  };

  const cerrarReservas = () => {
    setReservas([]);
  };

  return (
    <div>
      <h2>Mis Reservas</h2>
      {reservas.length === 0 ? (
        <p>No tienes reservas realizadas.</p>
      ) : (
        <ul>
          {reservas.map((reserva) => (
            <li key={reserva.id}>
              Hotel ID: {reserva.hotelId}, Fecha de Ingreso: {reserva.fechaIngreso}, Fecha de Egreso: {reserva.fechaEgreso}
            </li>
          ))}
        </ul>
      )}
      <button onClick={cerrarReservas}>Cerrar</button>
    </div>
  );
}

export default Reservas;