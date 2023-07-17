import React, { useEffect, useRef, useState } from 'react';
import '../Stylesheet/Hoteles.css';

function Hoteles(props) {
    const [hotels, setHotels] = useState([]);
    const [amenities, setAmenities] = useState([]);
    const hotelIdRef = useRef(props.hotelId);

    useEffect(() => {
        // Obtener los IDs de los hoteles mediante una solicitud fetch
        fetch('http://localhost:8090/hotels')
            .then(response => response.json())
            .then(data => {
                setHotels(data);
            })
            .catch(error => {
                console.error('Error:', error);
            });
    }, []);

    useEffect(() => {
        hotelIdRef.current = props.hotelId;
    }, [props.hotelId]);

    useEffect(() => {
        const fetchAmenities = async () => {
            try {
                const response = await fetch(`http://localhost:8090/amenitiehotel/${props.hotelId}`);
                if (response.ok) {
                    const data = await response.json();
                    setAmenities(data);

                    const amenityIds = data.map(amenity => amenity.amenitie_id); // Obtener los IDs de las amenidades
                    const amenityTypes = await fetchAmenityTypes(amenityIds); // Obtener los tipos de amenidades
                    const amenitiesWithTypes = data.map((amenity, index) => ({
                        ...amenity,
                        tipo: amenityTypes[index] // Agregar el tipo de amenidad al objeto de amenidad
                    }));
                    setAmenities(amenitiesWithTypes);
                } else {
                    throw new Error(`Error en la petición GET de las amenidades para el hotel ${props.hotelId}`);
                }
            } catch (error) {
                console.error('Error:', error);
            }
        };

        if (hotelIdRef.current) {
            fetchAmenities();
        }
    }, []);

    const fetchAmenityTypes = async (amenityIds) => {
        const amenityTypesPromises = amenityIds.map(async (amenityId) => {
            const response = await fetch(`http://localhost:8090/amenities/${amenityId}`);
            if (response.ok) {
                const data = await response.json();
                return data.tipo; // Suponiendo que el tipo se encuentra en la propiedad "tipo"
            } else {
                throw new Error(`Error en la petición GET de la amenidad ${amenityId}`);
            }
        });

        try {
            const amenityTypes = await Promise.all(amenityTypesPromises);
            return amenityTypes;
        } catch (error) {
            console.error('Error:', error);
            return [];
        }
    };

    return (
        <div className="contenedor-hoteles">
            <p className="nombre-hotel1">
                <strong>{props.nombreHotel}</strong>
            </p>
            <p className="cantidad-piezas">
                Habitaciones: {props.piezas}
            </p>
            <p className="descripcion-hotel">
                Descripción: {props.descripcion}
            </p>

            <h2>Amenities del hotel:</h2>
            <ul>
                {amenities.map(amenitie => (
                    <li key={amenitie.id}>{amenitie.tipo}</li>
                ))}
            </ul>
        </div>
    );
}

export default Hoteles;
