import React from 'react'
import '../Stylesheet/Hoteles.css'

function Hoteles(){
    
    const handleReservaSubmit = () => {
        alert("Para reservar, primero debe iniciar sesión.");
    }

    return(
    <div className="contenedor-hoteles">
      <div className="contenedor-detalle-hoteles">
        <p className="nombre-hotel1">
          <strong>Hotel</strong>
        </p>
        <p className="cantidad-piezas">Habitaciones: 2</p>
      </div>
      <forms onClick={handleReservaSubmit} className="boton-reserva">
        <button type="button">Reservar</button>
      </forms>
    </div>
    );
}

export default Hoteles;