import React, { useState } from "react";
import DatePicker from "react-datepicker";
import "react-datepicker/dist/react-datepicker.css";
import "../../src/styleSheet/Hoteles.css";

function Hoteles(props) {
  const [fechaDesde, setFechaDesde] = useState(null);
  const [fechaHasta, setFechaHasta] = useState(null);

  const handleFechaDesdeChange = (date) => {
    setFechaDesde(date);
  };

  const handleFechaHastaChange = (date) => {
    setFechaHasta(date);
  };

  const handleReservaClick = () => {
    if (!props.isLoggedIn) {
      alert("Debes iniciar sesión para realizar la reserva.");
    } else if (!fechaDesde || !fechaHasta) {
      alert("Debes completar los campos de fecha desde y fecha hasta.");
    } else if (fechaDesde >= fechaHasta) {
      alert("La fecha desde debe ser anterior a la fecha hasta.");
    } else {
      // Realizar la solicitud GET para verificar la disponibilidad de habitaciones
      fetch('http://localhost:8090/reservas/rooms', {
        method: 'GET',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          fecha_ingreso: fechaDesde,
          fecha_egreso: fechaHasta,
          hotel_id: props.hotelId,
          user_id: props.userId,
        }),
      })
        .then((response) => response.json())
        .then((data) => {
          if (data.disponibilidad === "true") {
            // Realizar la solicitud POST para insertar la reserva
            fetch('http://localhost:8090/reservas', {
              method: 'POST',
              headers: {
                'Content-Type': 'application/json',
              },
              body: JSON.stringify({
                fecha_ingreso: fechaDesde,
                fecha_egreso: fechaHasta,
                hotel_id: props.hotelId,
                user_id: props.userId,
              }),
            })
              .then((response) => response.json())
              .then((data) => {
                alert("¡Reserva realizada con éxito!");
                setFechaDesde(null);
                setFechaHasta(null);
              })
              .catch((error) => {
                console.error('Error:', error);
              });
          } else {
            alert("No hay habitaciones disponibles para las fechas seleccionadas.");
          }
        })
        .catch((error) => {
          console.error('Error:', error);
        });
    }
  };

  return (
    <div className="contenedor-hoteles">
      <div className="contenedor-detalle-hoteles">
        <p className="nombre-hotel1">
          <strong>{props.nombreHotel}</strong>
        </p>
        <p className="cantidad-piezas">Habitaciones: {props.piezas}</p>
      </div>
      <div className="contenedor-inputs-fechas">
        <p>Desde: </p>
        <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
        <p>Hasta: </p>
        <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
      </div>
      <div className="boton-reserva">
        <button onClick={handleReservaClick}>Reservar</button>
      </div>
    </div>
  );
}

export default Hoteles;