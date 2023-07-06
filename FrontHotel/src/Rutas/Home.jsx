import React, { useState, useEffect } from "react";
import '../Stylesheet/Home.css';
import Header from '../Componentes/Header';
import { useNavigate } from "react-router-dom";
import Hoteles from '../Componentes/Hoteles';

function Home() {
    const [hoteles, setHoteles] = useState([]);
    const navigate = useNavigate();

    const handleButtonSubmit = () => {
        navigate('/loginandregister');
    }

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await fetch('http://localhost:8090/hotels', {
                    method: 'GET',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });
                if (response.ok) {
                    const data = await response.json();
                    setHoteles(data);
                } else {
                    throw new Error('Error en la petición GET');
                }
            } catch (error) {
                console.error(error);
            }
        };

        fetchData();
    }, []);

    return (
        <div className="header">
            <h1>SUMMIT LUXURY Hotels</h1>
            <form onSubmit={handleButtonSubmit} className="contenedor-boton">
                <button type="submit" className="boton-uno">
                    <span>
                        <strong>INGRESA PARA HACER TUS RESERVAS!</strong>
                    </span>
                </button>
            </form>
            <div className="contenedor-de-hoteles">
                {hoteles.map((hotel) => (
                    <Hoteles
                        key={hotel.id}
                        nombreHotel={hotel.name}
                        piezas={hotel.cantHabitaciones}
                        descripcion={hotel.descripcion}
                    />
                ))}
            </div>
        </div>
    );
}

export default Home;
