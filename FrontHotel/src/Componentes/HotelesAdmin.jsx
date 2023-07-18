import React, { useEffect, useRef, useState } from 'react';
import '../Stylesheet/HotelesAdmin.css';

function HotelesAdmin(props) {
    const [hotels, setHotels] = useState([]);
    const [amenities, setAmenities] = useState([]);
    const [todosAmenities, setTodosAmenities] = useState([]);
    const hotelIdRef = useRef(props.hotelId);
    const [imagenSeleccionada, setImagenSeleccionada] = useState(null);
    const [eleccionDeAmenitie, setEleccionDeAmenitie] = useState(null);

    const handleImagenesChange = (event) => {
        event.preventDefault();
        const { files } = event.target;
        if (files && files.length > 0) {
            const imagenFile = files[0];
            const reader = new FileReader();
            reader.onloadend = () => {
                const buffer = new Uint8Array(reader.result); // Obtenemos el Uint8Array
                const byteArray = Array.from(buffer); // Convertimos el Uint8Array a []byte
                setImagenSeleccionada(byteArray);
            };
            reader.readAsArrayBuffer(imagenFile);
        }
    };

    const HandleOpcionAmenitieChange = (event) => {
        event.preventDefault();
        setEleccionDeAmenitie(event.target.value);
    }

    const agregarImagen = (event) => {
        event.preventDefault();

        if (!imagenSeleccionada) {
            alert("Por favor, seleccione una imagen.");
            return;
        }

        const imagenData = {
            hotel_id: props.hotelId,
            nombre: props.nombreHotel,
            contenido: imagenSeleccionada,
        };

        fetch("http://localhost:8090/imagenes", {
            method: "POST",
            headers: {
                Authorization: `${props.token}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(imagenData),
        })
            .then((response) => response.json())
            .then((data) => {
                if (data.error) {
                    console.error(data.error);
                } else {
                    alert("Imagen agregada con éxito");
                    setImagenSeleccionada(null);
                    window.location.reload();
                }
            })
            .catch((error) => console.error(error));

    };

    const agregarAmenitie = (event) => {
        event.preventDefault();
        const amenitieElegidaId = parseInt(eleccionDeAmenitie);
        const amenitieExistente = amenities.find(amenitie => amenitie.amenitie_id === amenitieElegidaId);
    
        if (amenitieExistente) {
            alert("El amenitie seleccionado ya existe en el hotel.");
            return;
        }
    
        const amenitieData = {
            hotel_id: props.hotelId,
            amenitie_id: amenitieElegidaId,
        };
    
        fetch("http://localhost:8090/amenitiehotel/assign", {
            method: "POST",
            headers: {
                Authorization: `${props.token}`,
                "Content-Type": "application/json",
            },
            body: JSON.stringify(amenitieData),
        })
        .then((response) => response.json())
        .then((data) => {
            if (data.error) {
                console.error(data.error);
            } else {
                alert("Amenitie agregado con éxito");
                setEleccionDeAmenitie(null);
                window.location.reload();
            }
        })
        .catch((error) => console.error(error));
    }
    

    useEffect(() => {
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
        const fetchTodosAmenities = async () => {
            try {
                const response = await fetch("http://localhost:8090/amenities");
                if (response.ok) {
                    const data = await response.json();
                    setTodosAmenities(data);
                } else {
                    throw new Error(`Error en la petición GET de las amenidades para el hotel ${props.hotelId}`);
                }
            } catch (error) {
                console.error('Error:', error);
            }
        };

        if (hotelIdRef.current) {
            fetchTodosAmenities();
        }
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
            <form>
                <h5>Cargar imagen:</h5>
                <input
                    type="file"
                    name="imagen"
                    onChange={handleImagenesChange}
                />
                <button className="boton-subir-campos" onClick={agregarImagen}>Agregar imagen</button>
            </form>
            <form>
                <h5>Amenities disponibles para cargar:</h5>
                <ol>
                    {todosAmenities.sort((a, b) => a.id - b.id).map(amenitie => (
                        <li key={amenitie.id}>{amenitie.tipo}</li>
                    ))}
                </ol>
                <h5>Elija el número de amenitie a cargar para el hotel:</h5>
                <input type="number" onChange={HandleOpcionAmenitieChange}></input>
                <button className="boton-subir-campos" onClick={agregarAmenitie}>Agregar amenitie</button>
            </form>

        </div>
    );
}

export default HotelesAdmin;