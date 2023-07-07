import React, { useState } from "react"
import '../Stylesheet/Reserve.css'
import Header from '../Componentes/Header'
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

function Reserve() {
    const [fechaDesde, setFechaDesde] = useState(null);
    const [fechaHasta, setFechaHasta] = useState(null);
    const [hotelesDisponibles, setHotelesDisponibles] = useState([]);
    const clavetoken = "secreto";

    const handleFechaDesdeChange = (date) => {
        setFechaDesde(date);
    };

    const handleFechaHastaChange = (date) => {
        setFechaHasta(date);
    };

    const buscarHotelesDisponibles = () => {
        if (fechaDesde && fechaHasta) {
            // Realizar la solicitud al backend para obtener los hoteles disponibles
            fetch("/reservas/hotelsbyfecha", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `Bearer ${clavetoken}`
                },
                body: JSON.stringify({
                    fechaDesde,
                    fechaHasta,
                }),
            })
                .then((response) => response.json())
                .then((data) => {
                    // Actualizar el estado con los hoteles disponibles obtenidos del backend
                    setHotelesDisponibles(data);
                })
                .catch((error) => console.error(error));
        } else {
            if (fechaDesde > fechaHasta) {
                alert("La fecha de checkin no puede ser mayor a la de checkout.");
            } else {
                alert("Por favor, seleccione las fechas desde y hasta.");
            }
        }
    };

<<<<<<< HEAD



  return (
      <div>
          <div className="header">
              <h1>SUMMIT LUXURY Hotels</h1>
          </div>

          <div>
              <br></br>
              <h2>INGRESE LAS FECHAS para su estadia</h2>
              <div>
                  <p>Desde: </p>
                  <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
              </div>
              <div>
                  <p>Hasta: </p>
                  <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
              </div>
              <button onClick={buscarHotelesDisponibles}>BUSCAR</button>
          </div>

    {hotelesDisponibles.length > 0 && (
=======
    return (
>>>>>>> 97b7dce058677dd1620a026c7aaa725455482576
        <div>
            <Header />
            <div>
                <br></br>
                <h2>INGRESE LAS FECHAS para su estadia</h2>
                <div>
                    <p>Desde: </p>
                    <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
                </div>
                <div>
                    <p>Hasta: </p>
                    <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
                </div>
                <button onClick={buscarHotelesDisponibles}>BUSCAR</button>
            </div>
            {hotelesDisponibles.length > 0 && (
                <div>
                    <h2>Hoteles Disponibles:</h2>
                    <ul>
                        {hotelesDisponibles.map((hotel) => (
                            <li key={hotel.id}>
                                <p>Nombre: {hotel.nombre}</p>
                                <p>Cantidad de Habitaciones: {hotel.cantidadHabitaciones}</p>
                            </li>
                        ))}
                    </ul>
                </div>
            )}
        </div>
    );
}

export default Reserve;