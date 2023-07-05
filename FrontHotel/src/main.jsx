import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './Rutas/Home';
import LoginRegister from './Rutas/LoginRegister';
import Reserve from './Rutas/Reserve';

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/loginandregister" element={<LoginRegister />} />
        <Route path="/reserve" element={<Reserve />} />
      </Routes>
    </Router>
  </React.StrictMode>
);