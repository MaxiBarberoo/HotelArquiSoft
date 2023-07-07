import React, { useState } from "react"
import '../Stylesheet/Reserve.css'
import Header from '../Componentes/Header'
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

function Reserve(){

    const [fechaDesde, setFechaDesde] = useState(null);
    const [fechaHasta, setFechaHasta] = useState(null);

    const handleFechaDesdeChange = (date) => {
        setFechaDesde(date);
    };

    const handleFechaHastaChange = (date) => {
        setFechaHasta(date);
    };


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
              <button>Reservar</button>
          </div>

          </div>
  );
}

export default Reserve;