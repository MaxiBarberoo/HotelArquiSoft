import React, { useState, useEffect } from "react"
import '../Stylesheet/Reserve.css'
import Header from '../Componentes/Header'
import DatePicker from 'react-datepicker'
import HotelesR from '../Componentes/HotelesR'
import 'react-datepicker/dist/react-datepicker.css'
import { useParams } from 'react-router-dom';

function Reserve() {
    const [fechaDesde, setFechaDesde] = useState(null);
    const [fechaHasta, setFechaHasta] = useState(null);
    const [hotelesDisponibles, setHotelesDisponibles] = useState([]);
    const [reservas, setReservas] = useState([]); // Agregar el estado de las reservas
    const { token } = useParams();
    const { user_id } = useParams();

    const clavetoken = "secreto";

    const handleFechaDesdeChange = (date) => {
        setFechaDesde(date);
    };

    const handleFechaHastaChange = (date) => {
        setFechaHasta(date);
    };

    const buscarHotelesDisponibles = () => {
        if (!fechaDesde || !fechaHasta) {
            alert("Debes completar los campos de fecha desde y fecha hasta.");
        } else if (fechaDesde >= fechaHasta) {
            alert("La fecha desde debe ser anterior a la fecha hasta.");
        } else {
            fetch("http://localhost:8090/reservas/hotelsbyfecha", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    "Authorization": `${token}`,
                },
                body: JSON.stringify({
                    fecha_ingreso: fechaDesde,
                    fecha_egreso: fechaHasta,
                }),
            })
                .then((response) => response.json())
                .then((data) => {
                    // Obtener los ID's de los hoteles disponibles
                    const hotelIds = Object.values(data).map((item) => item.hotel_id);
                    // Realizar las solicitudes GET para obtener la información completa de los hoteles
                    Promise.all(
                        hotelIds.map((hotelId) =>
                            fetch(`http://localhost:8090/hotels/${hotelId}`, {
                                method: "GET",
                                headers: {
                                    Authorization: `${token}`,
                                },
                            }).then((response) => response.json())
                        )
                    )
                        .then((hotelData) => {
                            // Actualizar el estado con la información completa de los hoteles
                            setHotelesDisponibles(hotelData);
                            // Actualizar el estado con las reservas del usuario obtenidas del backend
                        })
                        .catch((error) => console.error(error));
                })
                .catch((error) => console.error(error));
        }
    };

    useEffect(() => {
        fetch(`http://localhost:8090/reservas/reservauser/${user_id}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
            },
        })
            .then((response) => response.json())
            .then((data) => {
                setReservas(data);
            })
    }, []);


    return (
        <div className="contenedor-principal">
            <Header />
            <div className="contenedor-fechas">
                <br></br>
                <h2>INGRESE LAS FECHAS para su estadia</h2>
                <div className="fecha-desde">
                    <p>Desde: </p>
                    <DatePicker selected={fechaDesde} onChange={handleFechaDesdeChange} />
                </div>
                <div className="fecha-hasta">
                    <p>Hasta: </p>
                    <DatePicker selected={fechaHasta} onChange={handleFechaHastaChange} />
                </div>
                <button onClick={buscarHotelesDisponibles}>BUSCAR</button>
            </div>
            {hotelesDisponibles.length > 0 ? (
                hotelesDisponibles.map((hotel) => (
                    <div key={hotel.id}>
                        <HotelesR
                            nombreHotel={hotel.name}
                            piezas={hotel.cantHabitaciones}
                            descripcion={hotel.descripcion}
                            hotelId={hotel.id}
                            userId={user_id}
                            fechaDesde={fechaDesde}
                            fechaHasta={fechaHasta}
                            token={token}
                        />
                    </div>
                ))
            ) : (
                <p>No hay hoteles disponibles en esas fechas.</p>
            )}
        </div>
    );
}

export default Reserve;
