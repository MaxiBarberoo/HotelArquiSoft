import React from "react"
import '../Stylesheet/Home.css'
import Header from '../Componentes/Header'
import { useNavigate } from "react-router-dom"

function Home() {
    const navigate = useNavigate();
    
    const handleButtonSubmit = () =>{
        navigate('/loginandregister');
    }

    return(

        <div>
        <div className="header">
            <h1>SUMMIT LUXURY Hotels</h1>
            <form onSubmit={handleButtonSubmit} className="contenedor-boton">
                <button type="submit" className="boton-uno"><span><strong>INGRESA PARA HACER TUS RESERVAS!</strong></span></button>
            </form>

        </div>
            <div className="contenedor-hoteles">
                <div className="contenedor-detalle-hoteles">
                    <p className="nombre-hotel1"></p>
                </div>
            </div>

        </div>
    );
}

export default Home;