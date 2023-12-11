import React, { useState } from 'react';
import Header from '../../components/Header/Header'; // Gantilah dengan path yang sesuai
import Footer from '../../components/Footer/Footer';
import './CekReservasi.css';

function CekReservasi() {
  const [noRS, setNoRS] = useState('');
  const [resultData, setResultData] = useState([]);
  const [showResult, setShowResult] = useState(false);
  const [updatePopupData, setUpdatePopupData] = useState(null);

  const handleChange = (e) => {
    setNoRS(e.target.value);
  };

  const cekReservasi = () => {
    const noRSValue = noRS;

    fetch(`http://localhost:8080/reservasi/api/byNoRS/${noRSValue}`)
      .then(response => {
        console.log('Response status:', response.status);
        return response.json();
      })
      .then(data => {
        console.log('Response data:', data);
        setResultData(data);
        setShowResult(true);
      })
      .catch(error => {
        console.error("Error fetching data:", error);
      });
  };

  const openUpdatePopup = (idReservasi) => {
    fetch(`http://localhost:8080/reservasi/api/byID/${idReservasi}`)
      .then(response => response.json())
      .then(data => {
        setUpdatePopupData(data[0]);
      })
      .catch(error => {
        console.error("Error fetching data:", error);
      });
  };

  const submitUpdate = (updatedData) => {
    fetch(`http://localhost:8080/reservasi/api/byID/${updatePopupData.id_reservasi}`, {
      method: 'PUT',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(updatedData),
    })
      .then(response => response.json())
      .then(data => {
        closePopup("updatePopup");
        cekReservasi(updatePopupData.id_reservasi);
        alert("Data reservasi berhasil diupdate!");
      })
      .catch(error => {
        console.error("Error updating data:", error);
        alert("Terjadi kesalahan saat melakukan update data. Silakan coba lagi.");
      });
  };

  const closePopup = (popupId) => {
    if (popupId === 'updatePopup') {
      setUpdatePopupData(null);
    }
  };

  const deleteReservasi = (reservasiId) => {
    const isConfirmed = window.confirm("Apakah Anda yakin ingin menghapus data?");

    if (!isConfirmed) {
      return;
    }

    fetch(`http://localhost:8080/reservasi/api/byID/${reservasiId}`, {
      method: 'DELETE',
    })
      .then(response => response.json())
      .then(data => {
        console.log(data.message);
        cekReservasi(reservasiId);
      })
      .catch(error => {
        console.error("Error deleting data:", error);
      });
  };

  const UpdatePopup = ({ data, onUpdate, onClose }) => {
    return (
      <div id="updatePopup" className="popup" style={{ display: 'flex' }}>
        <div className="popup-content">
          <span className="close" onClick={() => onClose('updatePopup')}>&times;</span>
          <h2>Update Reservasi</h2>
          <form id="updateForm">
            {/* Implementasikan form update reservasi di sini */}
            {/* Pastikan untuk menggunakan onChange untuk mengubah nilai state saat input berubah */}
          </form>
        </div>
      </div>
    );
  };

  return (
    <div>
      <Header />

      <div className="isi">
        <h1>Cek Reservasi</h1>
      </div>

      <div className="cek-reservasi-container">
        <form id="cekReservasiForm" className="cek-reservasi-form" action="/reservasi/api/byNoRS">
          <label>Masukkan Nomor RS</label>
          <input type="text" id="NoRS" name="no_rs" onChange={handleChange} required />
          <br />
          <button type="button" onClick={cekReservasi}>Cek Reservasi</button>
        </form>

        <div id="resultContainer" style={{ display: showResult ? 'block' : 'none' }}>
          <p id="noDataMessage" style={{ display: resultData.length === 0 ? 'block' : 'none' }}>Data tidak ditemukan!</p>
          <table id="resultTable" border="1">
            <thead>
              <tr>
                <th>Nomor RS</th>
                <th>NIK</th>
                <th>Tanggal Kunjungan</th>
                <th>Nama</th>
                <th>Jenis Kelamin</th>
                <th>Tanggal Lahir</th>
                <th>Email</th>
                <th>Nomor Telepon</th>
                <th>Pembayaran</th>
                <th>Aksi</th>
              </tr>
            </thead>
            <tbody id="resultTableBody">
              {resultData.map(reservasi => (
                <tr key={reservasi.id_reservasi}>
                  <td>{reservasi.no_rs}</td>
                  <td>{reservasi.nik}</td>
                  <td>{reservasi.tgl_kunjungan}</td>
                  <td>{reservasi.nama}</td>
                  <td>{reservasi.jenis_kelamin}</td>
                  <td>{reservasi.tgl_lahir}</td>
                  <td>{reservasi.email}</td>
                  <td>{reservasi.no_telp}</td>
                  <td>{reservasi.pembayaran}</td>
                  <td>
                    <button onClick={() => openUpdatePopup(reservasi.id_reservasi)}>Update</button>
                    <p></p>
                    <button onClick={() => deleteReservasi(reservasi.id_reservasi)}>Delete</button>
                  </td>
                </tr>
              ))}
            </tbody>
          </table>
        </div>

        {updatePopupData && (
          <UpdatePopup
            data={updatePopupData}
            onUpdate={submitUpdate}
            onClose={(popupId) => closePopup(popupId)}
          />
        )}
      </div>

      <Footer />
    </div>
  );
}

export default CekReservasi;
