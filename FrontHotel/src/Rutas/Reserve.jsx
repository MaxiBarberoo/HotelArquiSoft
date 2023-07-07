import React from "react"
import '../Stylesheet/Reserve.css'
import Header from '../Componentes/Header'

import dateRangePicker from "react-date-range/dist/components/DateRangePicker/index.js";
import


function Reserve(){
  return (
      <div>
          <div className="header">
              <h1>SUMMIT LUXURY Hotels</h1>
          </div>

          <div className="Reserve">
      <h1>Reserve</h1>
    </div>

          <DateRangePicker>






          </DateRangePicker>




          </div>
  );
}

export default Reserve;