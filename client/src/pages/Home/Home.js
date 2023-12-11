// Home.js
import React, { useEffect } from 'react';
import Header from '../../components/Header/Header';
import Footer from '../../components/Footer/Footer';
import './Home.css';

const Home = () => {
  useEffect(() => {
    const animatedDiv = document.getElementById('animated-div');

    function fadeInOnScroll() {
      var rect = animatedDiv.getBoundingClientRect();
      var isVisible = rect.top < window.innerHeight && rect.bottom >= 0;

      if (isVisible) {
        animatedDiv.classList.add('show');
        window.removeEventListener('scroll', fadeInOnScroll);
      }
    }

    window.addEventListener('scroll', fadeInOnScroll);
    fadeInOnScroll();
  }, []);

  return (
    <div>
      <Header />
      
      <div className="head">
        <h1>Selamat Datang di Klinik Ridsu</h1>
        <p>Tempat Terbaik untuk Pelayanan Kesehatan</p>
        <img style={{ width: '500px' }} src={require('../../assets/images/klinik.jpg')} alt=""/>
      </div>

      <section id="Informasi">
        <h1 style={{ textAlign: 'center' }}>INFORMASI</h1>
        <div className="isi" id="animated-div">
          <div className="dalam">
            <h2>Klinik Terbaik</h2>
            <img style={{ width: '500px' }} src={require('../../assets/images/tempat.jpg')} alt="" />
            <p>
              Klinik Ridsu adalah pilihan terbaik untuk layanan kesehatan yang berkualitas dan komprehensif. Dengan
              komitmen untuk memberikan perawatan yang canggih dan perhatian yang personal, klinik kami menjadi destinasi
              utama bagi pasien yang mencari pengalaman medis yang unggul.
            </p>
          </div>

          <div className="dalam">
            <h2>Dokter Kompeten</h2>
            <img style={{ width: '500px' }} src={require('../../assets/images/bagus.jpeg')} alt="" />
            <p>
              Tim profesional kami terdiri dari dokter dan tenaga kesehatan berpengalaman yang didedikasikan untuk
              memberikan perawatan holistik kepada setiap pasien. Dengan fasilitas modern dan peralatan medis terkini,
              Klinik Ridsu menawarkan lingkungan yang nyaman dan aman bagi setiap kunjungan.
            </p>
          </div>
        </div>
      </section>

      <Footer />
    </div>
  );
};

export default Home;
