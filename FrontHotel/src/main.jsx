import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Home from './Rutas/Home';
import LoginRegister from './Rutas/LoginRegister';
import Reserve from './Rutas/Reserve';
import Admin from './Rutas/Admin.jsx';

ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/loginandregister" element={<LoginRegister />} />
        <Route path="/reserve" element={<Reserve />} />
        <Route path= "/Admin" element={<Admin />} />
      </Routes>
    </Router>
  </React.StrictMode>
);
