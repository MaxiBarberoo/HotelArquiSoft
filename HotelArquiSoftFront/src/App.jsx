<<<<<<< HEAD
// eslint-disable-next-line no-unused-vars
import React from "react";
import './App.css'
import Hoteles from './Componentes/Hoteles.jsx'

    function App() {
        <div>
             <div>
               <h1>ENCONTRA LA MEJOR OPCION</h1>

                 <Hoteles
                    nombreHotel='HOTEL LUXURY'
                    image='1'
                    piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='HOTEL PALACE'
                     image='2'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='HOTEL CENTRAL PARK'
                     image='3'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='HOTEL PARAÍSO'
                     image='4'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='CASA DE LAS MONTAÑAS'
                     image='5'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='HOTEL EXOTIC'
                     image='6'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='HOTEL MANSION DEL SUR'
                     image='7'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='HOTEL JARDIN'
                     image='8'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='GRAN HOTEL'
                     image='9'
                     piezas= '5'
                 />
                 <Hoteles
                     nombreHotel='CASA NATURAL'
                     image='10'
                     piezas= '5'
                 />

             </div>
           </div>

        }
    export default App;
=======
import React, { useState } from 'react';
import './App.css'
import BotonLogin from './Componentes/BotonLogin';
import BotonRegister from './Componentes/BotonRegister';
import Hoteles from './Componentes/Hoteles.jsx'

function App() {
    return (
    <div className = 'App'>
        <div>
           <h1>ENCONTRA LA MEJOR OPCION</h1>
           <BotonLogin />
           <BotonRegister />
           <Hoteles
                nombreHotel='HOTEL LUXURY'
                image='1'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='HOTEL PALACE'
                image='2'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='HOTEL CENTRAL PARK'
                image='3'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='HOTEL PARAÍSO'
                image='4'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='CASA DE LAS MONTAÑAS'
                image='5'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='HOTEL EXOTIC'
                image='6'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='HOTEL MANSION DEL SUR'
                image='7'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='HOTEL JARDIN'
                image='8'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='GRAN HOTEL'
                image='9'
                piezas= '5'
            />
            <Hoteles
                nombreHotel='CASA NATURAL'
                image='10'
                piezas= '5'
            />
        </div>
    </div>
    )
}
export default App;
>>>>>>> 6e49f87269bdeb9c23bfb518fb72bf6788cce12f
