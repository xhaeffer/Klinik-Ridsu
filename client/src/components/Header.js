import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';

import './css/Header.css';

function Header () {
  const [isMenuOpen, setMenuOpen] = useState(false);
  const [isLoggedIn, setLoggedIn] = useState(false);

  useEffect(() => {
    const handleDocumentClick = (event) => {
      const menuList = document.getElementById("menu-list");
      const userInfo = document.getElementById("user-info");
      const menuIcon = document.getElementById("menu-icon");
    
      const isMenu = menuList.contains(event.target);
      const isUserInfo = userInfo.contains(event.target);
      const isIcon = menuIcon.contains(event.target);
    
      console.log('isMenu:', isMenu);
      console.log('isUserInfo:', isUserInfo);
      console.log('isIcon:', isIcon);
    
      // if (isMenuOpen && !isMenu && !isUserInfo && !isIcon) {
      //   setMenuOpen(false);
      // } else {
      //   menuList.classList.add('show'); // log ini untuk memastikan class 'show' ditambahkan saat mengklik menu
      // }
    };
    document.addEventListener("click", handleDocumentClick);

    return () => {
      document.removeEventListener("click", handleDocumentClick);
    };
  }, []);

  useEffect(() => {
    const loggedIn = checkUserLoginStatus();
    setLoggedIn(loggedIn);
  }, []);

  const toggleMenu = () => {
    console.log('Toggle Menu Clicked');
    setMenuOpen(!isMenuOpen);
  };

  console.log('isMenuOpen:', isMenuOpen);

  return (
    <div className={`navbarr ${isMenuOpen ? 'open' : ''}`}>
      <div className="menu-icon" id="menu-icon" onClick={toggleMenu}>
        <div className="bar"></div>
        <div className="bar"></div>
        <div className="bar"></div>
      </div>

    <ul className="menu-list" id="menu-list">
      <div id="user-info" style={{ display: isLoggedIn ? 'block' : 'none' }}>
        {/* USER INFO GOES HERE */}
        <img src= {require('./img/kosong.jpeg')} id="profile-image" alt="" />
        <div className="user-text">
          <p id="greeting">{getGreeting()}</p>
          <p id="user-name">{isLoggedIn ? 'Nama Pengguna' : 'Anda belum login'}</p>
          <p id="user-no-rs">{isLoggedIn ? 'No RS: RS12345' : 'Silahkan login terlebih dahulu'}</p>
        </div>
      </div>

      <li><Link to="/">HOME</Link></li>
      <li><Link to="/about">ABOUT US</Link></li>
      <li><Link to="/jadwal">FIND DOCTOR</Link></li>
      <li>
        <span>APPOINTMENT</span>
        <ul className="submenu">
          <li><Link to="/reservasi">BOOK APPOINTMENT</Link></li>
          <li><Link to="/cekreservasi">CHECK APPOINTMENT</Link></li>
        </ul>
      </li>
        {/* Uncomment the following lines when you have the corresponding routes */}
        {/* <li id="login-menu"><Link to="/login">LOGIN</Link></li>
        <li id="logout-menu" style={{ display: isLoggedIn ? 'block' : 'none' }}><Link to="/logout">LOGOUT</Link></li> */}
      </ul>
    </div>
  );
}

function getGreeting() {
  const currentTime = new Date().getHours();
  let greetingMessage = "Selamat ";

  if (currentTime < 12) {
    greetingMessage += "Pagi";
  } else if (currentTime < 18) {
    greetingMessage += "Siang";
  } else {
    greetingMessage += "Malam";
  }

  return greetingMessage;
}

function checkUserLoginStatus() {
  const token = getCookie("token"); // Change "token" to the appropriate cookie name
  const userSession = getSession("user"); // Change "user" to the appropriate session name

  // Adjust the logic based on your needs
  return token !== null && token !== undefined && token !== "" && userSession !== null && userSession !== undefined;
}

function getSession(key) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${key}=`);

  if (parts.length === 2) {
    return JSON.parse(parts.pop().split(';').shift());
  }

  return null;
}

function getCookie(name) {
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);

  if (parts.length === 2) {
    return parts.pop().split(';').shift();
  }

  return null;
}

export default Header;
