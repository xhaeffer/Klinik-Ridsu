import React, { useState, useEffect } from 'react';
import Header from './Header';
import Footer from './Footer';
import './css/Jadwal.css';

const JadwalDokter = () => {
  const [data, setData] = useState([]);

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await fetch('http://localhost:8080/jadwal');
        if (!response.ok) {
          throw new Error('Gagal mengambil data dari server');
        }

        const jsonData = await response.json();
        setData(jsonData);
      } catch (error) {
        console.error(error);
      }
    };

    fetchData();
  }, []); //

  return (
    <div>
      <Header />
      <div className="isi">
        <h1>Jadwal Dokter</h1>
      </div>

      <div className="schedule">
        {data.map((dokter) => (
          <div key={dokter.nama_dokter} className="card">
            <h2>{dokter.nama_dokter}</h2>
            <h3>Poli: {dokter.poli}</h3>
            <div className="event">
              <img src={`data:image/png;base64,${dokter.gambar}`} width="300" alt={dokter.nama_dokter} />
              <table>
                <thead>
                  <tr>
                    <th>Hari</th>
                    <th>Jam Mulai</th>
                    <th>Jam Selesai</th>
                  </tr>
                </thead>
                <tbody>
                  {dokter.JadwalDokter.map((jadwal) => (
                    <tr key={jadwal.hari}>
                      <td>{jadwal.hari}</td>
                      <td>{jadwal.jam_mulai || '-'}</td>
                      <td>{jadwal.jam_selesai || '-'}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
              <a href={`reservasi?nama=${dokter.nama_dokter}&poli=${dokter.poli}`} className="btn">
                Reservasi
              </a>
            </div>
          </div>
        ))}
      </div>

      <Footer />
    </div>
  );
};

export default JadwalDokter;
