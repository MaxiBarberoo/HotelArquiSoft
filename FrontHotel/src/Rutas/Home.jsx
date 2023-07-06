import React from "react"
import '../Stylesheet/Home.css'
import Header from '../Componentes/Header'

function Home() {
    return(
        <div className="header">
            <h1>SUMMIT LUXURY Hotels</h1>
            <div className="contenedor-boton">
                <button className="boton-uno"><span><strong> INGRESA PARA HACER TUS RESERVAS! </strong></span></button>
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