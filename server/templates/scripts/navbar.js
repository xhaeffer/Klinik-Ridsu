document.addEventListener("DOMContentLoaded", function () {
  const menuIcon = document.getElementById("menu-icon");
  const menuList = document.getElementById("menu-list");
  const userInfo = document.getElementById("user-info");

  menuIcon.addEventListener("click", function () {
    menuList.classList.toggle("show");
  });

  document.addEventListener("click", function (event) {
    const isMenu = menuList.contains(event.target);
    const isUserInfo = userInfo.contains(event.target);
    const isIcon = menuIcon.contains(event.target);

    if (!isMenu && !isUserInfo && !isIcon) {
      menuList.classList.remove("show");
    }
  });

  const isLoggedIn = checkUserLoginStatus();

  if (isLoggedIn) {
    showUserInfo();
    showLogoutOption();
  } else {
    showLoginPrompt();
  }

  function showUserInfo() {
    document.getElementById("profile-image").src = "server/templates/img/kosong.jpeg";
    document.getElementById("greeting").textContent = getGreeting();
    document.getElementById("user-name").textContent = "Nama Pengguna";
    document.getElementById("user-no-rs").textContent = "No RS: RS12345";
    userInfo.style.display = "block";
  }

  function showLoginPrompt() {
    document.getElementById("greeting").textContent = getGreeting();
    document.getElementById("user-name").textContent = "Anda belum login";
    document.getElementById("user-no-rs").textContent =
      "Silahkan login terlebih dahulu";
    userInfo.style.display = "block";

    const loginMenuItem = document.createElement("li");
    const loginLink = document.createElement("a");
    loginLink.href = "/login";
    loginLink.textContent = "LOGIN";
    loginMenuItem.appendChild(loginLink);
    menuList.appendChild(loginMenuItem);
  }

  function showLogoutOption() {
    const loginMenuItem = document.querySelector("li a[href='/login']");
    if (loginMenuItem) {
      loginMenuItem.parentNode.remove();
    }

    const logoutMenuItem = document.createElement("li");
    const logoutLink = document.createElement("a");
    logoutLink.href = "/logout";
    logoutLink.textContent = "LOGOUT";
    logoutMenuItem.appendChild(logoutLink);
    menuList.appendChild(logoutMenuItem);
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
    const token = getCookie("token"); // Ganti "token" dengan nama cookie yang sesuai
    const userSession = getSession("user"); // Ganti "user" dengan nama sesi yang sesuai
  
    // Ganti logika berikut sesuai dengan kebutuhan kamu
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
  
});
