import React, { useEffect, useState } from "react";

function Reservas({ reservas }) {
    const [hotelNames, setHotelNames] = useState({});

  useEffect(() => {
    // Obtener los nombres de los hoteles por ID
    const getHotelName = async (hotelId) => {
      try {
        const response = await fetch(`/hotels/${hotelId}`);
        const data = await response.json();
        if (response.ok) {
          setHotelNames((prevHotelNames) => ({
            ...prevHotelNames,
            [hotelId]: data.name,
          }));
        } else {
          // Manejar el error si la solicitud no fue exitosa
          console.error('Error al obtener el nombre del hotel:', data.error);
        }
      } catch (error) {
        // Manejar el error de la solicitud
        console.error('Error al obtener el nombre del hotel:', error);
      }
    };

    // Obtener los nombres de los hoteles para cada reserva
    reservas.forEach((reserva) => {
      if (!hotelNames[reserva.hotelId]) {
        getHotelName(reserva.hotelId);
      }
    });
  }, [reservas]);

  return (
    <div>
      <h2>Mis Reservas</h2>
      {reservas.map((reserva, index) => (
        <div key={index}>
          <p>Fecha de ingreso: {reserva.FechaIngreso}</p>
          <p>Fecha de egreso: {reserva.FechaEgreso}</p>
          <p>Hotel: {hotelNames[reserva.hotelId]}</p>
        </div>
      ))}
    </div>
  );
}

export default Reservas;