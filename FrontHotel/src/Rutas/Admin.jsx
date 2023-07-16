import React, { useState, useEffect } from "react"
import '../Stylesheet/Admin.css'
import Header from '../Componentes/Header'
import Hoteles from '../Componentes/Hoteles'
import { useParams } from 'react-router-dom';

function Admin() {
    const [reservasTotales, setReservasTotales] = useState([]);
    const [nuevoHotel, setNuevoHotel] = useState({
        nombre: '',
        descripcion: '',
        cantidadHabitaciones: 0,

    });
    const [hoteles, setHoteles] = useState([]);
    const { token } = useParams();
    const { user_id } = useParams();
    const [contadorReserva, setContadorReserva] = useState(1);

    const crearNuevoHotel = () => {
        const cantHabitaciones = parseInt(nuevoHotel.cantidadHabitaciones);
        fetch('http://localhost:8090/hotels', {
            method: 'POST',
            headers: {
                "Content-Type": 'application/json',
                "Authorization": `${token}`,
            },
            body: JSON.stringify({
                name: nuevoHotel.nombre,
                cantHabitaciones: cantHabitaciones
            })
        })
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.error(data.error);
                } else {
                    setHoteles((prevHoteles) => [...prevHoteles, data]);
                    alert('Se ha creado un nuevo hotel con éxito');
                    setNuevoHotel({
                        nombre: '',
                        cantidadHabitaciones: 0
                    });
                }
            })
            .catch((error) => console.error(error));
    };

    const handleNuevoHotelChange = (event) => {
        const { name, value } = event.target;
        setNuevoHotel((prevNuevoHotel) => ({
            ...prevNuevoHotel,
            [name]: value
        }));
    };

    useEffect(() => {
        fetch('http://localhost:8090/hotels')
            .then((response) => response.json())
            .then((data) => setHoteles(data))
            .catch((error) => console.error(error));


        fetch('http://localhost:8090/reservas', {
            method: 'GET',
            headers: {
                "Content-Type": 'application/json',
                "Authorization": `${token}`
            }
        })
            .then((response) => response.json())
            .then((data) => setReservasTotales(data))
            .catch((error) => console.error(error));
    }, []);

    const formatDate = (date) => {
        const formattedDate = new Date(date);
        formattedDate.setDate(formattedDate.getDate() + 1);

        const day = formattedDate.getDate();
        const month = formattedDate.getMonth() + 1;
        const year = formattedDate.getFullYear();

        return `${day.toString().padStart(2, '0')}/${month.toString().padStart(2, '0')}/${year}`;
    };

    return (
        <div>
            <Header />
            {reservasTotales.length > 0 ? (
                reservasTotales.map((reserva, index) => (
                    <div key={reserva.id} className="reserva-item">
                        <p className="subtitulo-reserva"> Datos de la reserva {contadorReserva + index}:</p>
                        <div className="detalle-reserva">
                            <p>Usuario: {reserva.usuario_nombre}</p>
                            <p>Hotel: {reserva.hotel_nombre}</p>
                            <p>Fecha check-in: {formatDate(reserva.fecha_ingreso)}</p>
                            <p>Fecha check-out: {formatDate(reserva.fecha_egreso)}</p>
                        </div>
                    </div>
                ))
            ) : (
                <p>No tiene reservas realizadas.</p>
            )}
            <h2>Agregar nuevo hotel</h2>
            <div className="contenedor-crear-hoteles">
                <form onSubmit={handleNuevoHotelChange}>
                    <p>Nombre del hotel:</p>
                    <input
                        type="text"
                        name="nombre"
                        value={nuevoHotel.nombre}
                        onChange={handleNuevoHotelChange}
                        placeholder="Nombre del hotel"
                    />
                    <p>Descripción del hotel:</p>
                    <input
                        type="text"
                        maxLength={500}
                        name="descripcion"
                        value={nuevoHotel.descripcion}
                        onChange={handleNuevoHotelChange}
                        placeholder="Descripción"
                    />
                    <p>Cantidad de habitaciones:</p>
                    <input
                        type="number"
                        name="cantidadHabitaciones"
                        value={nuevoHotel.cantidadHabitaciones}
                        onChange={handleNuevoHotelChange}
                        placeholder="Cantidad de habitaciones"
                    />
                    <button className="boton-crear-hotel" onSubmit={crearNuevoHotel}>Crear Hotel</button>
                </form>
            </div>
            <div className="contenedor-hoteles-admin">
                <h2>Listado de Hoteles</h2>
                {hoteles.map((hotel) => (
                    <Hoteles
                        key={hotel.id}
                        imagenesURLs={hotel.imagenesURLs}
                        nombreHotel={hotel.name}
                        piezas={hotel.cantHabitaciones}
                        descripcion={hotel.descripcion}
                    />
                ))}
            </div>
        </div>
    );
}

export default Admin;