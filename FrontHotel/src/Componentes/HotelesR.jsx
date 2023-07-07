import React from 'react'
import '../Stylesheet/Hoteles.css'

function Hoteles(props){

    const handleReservaSubmit = () => {
        if (usuario && hotel) {
            Reservar(usuario, hotel);
        } else {
            alert("Para reservar, primero debes iniciar sesión y seleccionar un hotel.");
        }
    };

    const Reservar = (usuario, hotel) => {
        // Realizar la solicitud al backend para guardar la reserva
        fetch("/reservas", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                fecha_ingreso: fechaDesde,
                fecha_egreso: fechaHasta,
                hotel_id: props.hotelId, // Agrega el ID del hotel correspondiente
                user_id: props.userId, // Agrega el ID del usuario correspondiente
            }),
        })
            .then((response) => response.json())
            .then((data) => {
                // Mostrar un mensaje de éxito o realizar otras acciones necesarias
                alert("Reserva realizada con éxito");
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    };

    return(
        <div className="contenedor-hoteles">
            <div className="contenedor-detalle-hoteles">
                <p className="nombre-hotel1">
                    <strong>{props.nombreHotel}</strong>
                </p>
                <p className="cantidad-piezas">
                    Habitaciones: {props.piezas}
                </p>
                <p className="descripcion-hotel">
                    Descripción: {props.descripcion}
                </p>
            </div>
            <form onClick={handleReservaSubmit} className="boton-reserva">
                <button type="button">Reservar</button>
            </form>
        </div>
    );
}

export default Hoteles;