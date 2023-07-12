import React from 'react'
import '../Stylesheet/HotelesR.css'

function HotelesR(props) {
  const checkDisponibilidad = (event) => {
    fetch("http://localhost:8090/reservas/rooms", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `${props.token}`,
      },
      body: JSON.stringify({
        fecha_ingreso: props.fechaDesde,
        fecha_egreso: props.fechaHasta,
        hotel_id: parseInt(props.hotelId),
        user_id: parseInt(props.userId),
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        if (data.disponibilidad === "true") {
          Reservar();
        } else {
          alert("No hay habitaciones disponibles en esas fechas.");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  };

  const Reservar = () => {
    fetch("http://localhost:8090/reservas", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `${props.token}`,
      },
      body: JSON.stringify({
        fecha_ingreso: props.fechaDesde,
        fecha_egreso: props.fechaHasta,
        hotel_id: parseInt(props.hotelId),
        user_id: parseInt(props.userId),
      }),
    })
      .then((response) => response.json())
      .then((data) => {
        // Actualizar el estado de las reservas si es necesario
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  };

  return (
    <div className="contenedor-hoteles">
        <p className="nombre-hotel1">
          <strong>{props.nombreHotel}</strong>
        </p>
        <p className="cantidad-piezas">Habitaciones: {props.piezas}</p>
        <p className="descripcion-hotel">
          Descripción: {props.descripcion}
        </p>
      <form onSubmit={checkDisponibilidad} className="boton-reserva">
        <button className='boton-reservar' type="submit">Reservar</button>
      </form>
    </div>
  );
}

export default HotelesR;
