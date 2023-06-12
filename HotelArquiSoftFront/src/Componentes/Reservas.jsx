import React, { useEffect, useState } from "react";

function Reservas({ userId }) {
    const [reservas, setReservas] = useState([]);
    const [hoteles, setHoteles] = useState([]);
  
    useEffect(() => {
      // Obtener las reservas del usuario
      const fetchReservas = async () => {
        try {
          const response = await fetch(`http://localhost:8090/reservas/reservauser/${userId}`);
          const data = await response.json();
          if (response.ok) {
            setReservas(data);
          } else {
            // Manejar el error si la solicitud no fue exitosa
            console.error('Error al obtener las reservas:', data.error);
          }
        } catch (error) {
          // Manejar el error de la solicitud
          console.error('Error al obtener las reservas:', error);
        }
      };
  
      // Obtener todos los hoteles
      const fetchHoteles = async () => {
        try {
          const response = await fetch('http://localhost:8090/hotels');
          const data = await response.json();
          if (response.ok) {
            setHoteles(data);
          } else {
            // Manejar el error si la solicitud no fue exitosa
            console.error('Error al obtener los hoteles:', data.error);
          }
        } catch (error) {
          // Manejar el error de la solicitud
          console.error('Error al obtener los hoteles:', error);
        }
      };
  
      // Llamar a las funciones de obtención de reservas y hoteles
      fetchReservas();
      fetchHoteles();
    }, [userId]);
  
    return (
      <div>
        <h2>Reservas</h2>
        {reservas.map((reserva) => (
          <div key={reserva.id}>
            <p>Fecha de ingreso: {reserva.fechaIngreso}</p>
            <p>Fecha de egreso: {reserva.fechaEgreso}</p>
            <p>
              Hotel: {hoteles.find((hotel) => hotel.id === reserva.hotelId)?.nombre}
            </p>
          </div>
        ))}
      </div>
    );
  }

  export default Reservas;