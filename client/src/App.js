import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Home from './components/Home';
import About from './components/About';
import Jadwal from './components/Jadwal';
// import Reservasi from './components/Reservasi';
import CekReservasi from './components/CekReservasi';
import Login from './components/Login';

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
