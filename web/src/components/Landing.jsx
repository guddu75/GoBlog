import React from 'react';
import { Link } from 'react-router-dom';


const Landing = () => {
  const handleLoginClick = () => {
    alert('Login button clicked! (Functionality not implemented)');
  };

  const handleSignupClick = () => {
    alert('Sign Up button clicked! (Functionality not implemented)');
  };

  return (
    <div className="flex flex-col items-center justify-center text-center p-4">
      {/* Hero Section of the landing page */}
      <section className="bg-gradient-to-br from-blue-500 to-purple-600 text-white py-16 px-6 rounded-3xl shadow-2xl max-w-4xl w-full mb-12 transform hover:scale-105 transition-transform duration-500 ease-in-out">
        {/* Main Heading for the Landing Page */}
        <h2 className="text-4xl md:text-6xl font-extrabold mb-6 leading-tight">
          Discover, Learn, Grow with goBlog
        </h2>
        <p className="text-xl md:text-2xl mb-8 opacity-90">
          Dive into a world of engaging articles and fresh perspectives.
        </p>
        {/* Login/Signup Buttons within the landing page's hero section */}
        <div className="flex justify-center space-x-4 mt-8">
          <Link to="/login" >
          <button
            className="inline-block bg-white text-blue-700 hover:bg-gray-100 font-bold py-3 px-6 rounded-full text-lg shadow-md transition-colors duration-300 transform hover:translate-y-1"
          >
            Login
          </button>
          </Link>
          <Link to="/signup" >
          <button
            className="inline-block bg-green-500 hover:bg-green-600 text-white font-bold py-3 px-6 rounded-full text-lg shadow-md transition-colors duration-300 transform hover:translate-y-1"
          >
            Sign Up
          </button>
          </Link>
        </div>
      </section>

      {/* About Section */}
      <section className="bg-white p-10 rounded-xl shadow-lg max-w-3xl w-full">
        <h2 className="text-4xl font-bold text-gray-800 mb-6">What We Offer</h2>
        <div className="grid grid-cols-1 md:grid-cols-2 gap-8 text-gray-700">
          <div className="text-left">
            <h3 className="text-2xl font-semibold text-blue-600 mb-3">Diverse Topics</h3>
            <p>
              From cutting-edge technology and programming tutorials to lifestyle tips and personal development, goBlog covers a wide array of subjects to keep you informed and inspired.
            </p>
          </div>
          <div className="text-left">
            <h3 className="text-2xl font-semibold text-blue-600 mb-3">Engaging Content</h3>
            <p>
              Our articles are crafted to be insightful, easy to understand, and thought-provoking. We believe in sharing knowledge in a way that resonates with our readers.
            </p>
          </div>
        </div>
      </section>
    </div>
  );
};

export default Landing;
