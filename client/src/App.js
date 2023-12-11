import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './pages/Home/Home';
import About from './pages/About/About';
import Jadwal from './pages/Jadwal/Jadwal';
// import Reservasi from './pages/Reservasi/Reservasi';
import CekReservasi from './pages/CekReservasi/CekReservasi';
import Login from './pages/Login/Login';

function App() {
  return (
    <Router>
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/about" element={<About />} />
        <Route path="/jadwal" element={<Jadwal />} />
        {/* <Route path="/reservasi" element={<Reservasi />} /> */}
        <Route path="/cekreservasi" element={<CekReservasi />} />
        <Route path="/Login" element={<Login />} />
      </Routes>
    </Router>
  );
}

export default App;
