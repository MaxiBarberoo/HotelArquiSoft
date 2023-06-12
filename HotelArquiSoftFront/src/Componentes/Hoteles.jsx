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
      // Lógica de reserva
      alert("¡Reserva realizada con éxito!");
      setFechaDesde(null);
      setFechaHasta(null);
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