import React, { useState } from 'react';
import Header from './Header'; // Sesuaikan dengan path yang sesuai
import Footer from './Footer'; // Sesuaikan dengan path yang sesuai
import './css/Login.css'; // Sesuaikan dengan path yang sesuai

function Login() {
  const [identifier, setIdentifier] = useState('');
  const [password, setPassword] = useState('');
  const [error, setError] = useState('');

  const handleLogin = async () => {
    try {
        const formData = new URLSearchParams();
        formData.append('identifier', identifier);
        formData.append('password', password);
        
        const response = await fetch('http://localhost:8080/login', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/x-www-form-urlencoded',
          },
          body: formData,
        });

      if (!response.ok) {
        throw new Error('Login failed');
      }
      console.log('Login successful');
      setError('');
    } catch (error) {
      console.error('Login error:', error);
      setError('Login failed. Please check your credentials.');
    }
  };

  return (
    <div>
      <Header />

      <div className="login-container">
        <h2>Login</h2>
        <form className="login-form" action="/login" method="post">
          <label>No RS / NIK:</label>
          <input
            type="text"
            name="identifier"
            value={identifier}
            onChange={(e) => setIdentifier(e.target.value)}
            placeholder="No RS / NIK"
            required
          />
          <br />
          <label>Password:</label>
          <input
            type="password"
            name="password"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            placeholder="Password"
            required
          />
          <br />
          <button type="button" onClick={handleLogin}>
            Login
          </button>
          {error && <p style={{ color: 'red' }}>{error}</p>}
          {/* Uncomment the lines below if you want to include a registration link */}
          {/* <p>Don't have an account? </p>
            <a href="/register" className="btn">Register</a> */}
        </form>
      </div>

      <Footer />
    </div>
  );
}

export default Login;
