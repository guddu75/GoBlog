import React, { useState } from 'react';
import { Link, useNavigate } from 'react-router-dom'; // Import Link for navigation
import axios from 'axios'; // Import axios for HTTP requests
import { API_URL } from '../util'; // Import API_URL from util

// Login component using an ES6 function expression
const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const redirect = useNavigate(); 

  const handleSubmit = (e) => {
    e.preventDefault();
    console.log('Login attempt:', { email, password });
    const userData = {
      email,
      password
    };

    axios.post(`${API_URL}/authentication/token`, userData)
      .then(response => {
        // console.log('Login successful:', response.data);
        // Handle successful login, e.g., redirect to feed or store token
        const token = response.data.data;
        localStorage.setItem('token', token); // Store token in localStorage
        redirect('/feed');
        console.log(response.data.data);
      })
      .catch(error => {
        console.error('Error during login:', error);
        alert('Login failed. Please check your credentials and try again.');
      });
    
  };

  return (
    <div className="flex items-center justify-center min-h-[calc(100vh-160px)] p-4">
      <div className="bg-white p-8 rounded-xl shadow-lg w-full max-w-md">
        <h2 className="text-3xl font-bold text-center text-gray-800 mb-6">Login</h2>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label htmlFor="email" className="block text-gray-700 text-sm font-medium mb-2">
              Email Address
            </label>
            <input
              type="email"
              id="email"
              className="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="you@example.com"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
          </div>
          <div>
            <label htmlFor="password" className="block text-gray-700 text-sm font-medium mb-2">
              Password
            </label>
            <input
              type="password"
              id="password"
              className="text-black w-full px-4 py-2 border border-gray-300 rounded-lg focus:outline-none focus:ring-2 focus:ring-blue-500"
              placeholder="********"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
          </div>
          <button
            type="submit"
            className="w-full bg-blue-600 hover:bg-blue-700 text-white font-bold py-3 px-4 rounded-lg transition-colors duration-300 shadow-md"
          >
            Login
          </button>
        </form>
        <p className="mt-6 text-center text-gray-600 text-sm">
          Don't have an account? <Link to="/signup" className="text-blue-600 hover:underline font-semibold">Sign Up</Link>
        </p>
      </div>
    </div>
  );
};

export default Login;
