import React, { useState, useEffect } from 'react';
import { Link, useNavigate } from 'react-router-dom';

const Navbar = () => {
  const [isLoggedIn, setIsLoggedIn] = useState(false);
  const [searchTerm, setSearchTerm] = useState('');
  const redirect = useNavigate(); // Hook for programmatic navigation

  useEffect(() => {
    const checkLoginStatus = () => {
      const token = localStorage.getItem('token');
      setIsLoggedIn(!!token); 
    };

    // Initial check
    checkLoginStatus();

    // Listen for changes in localStorage (e.g., from other tabs/windows)
    window.addEventListener('storage', checkLoginStatus);

    // Cleanup listener on component unmount
    return () => {
      window.removeEventListener('storage', checkLoginStatus);
    };
  }, []);

  // Handle logout
  const handleLogout = () => {
    localStorage.removeItem('token'); // Remove the token
    setIsLoggedIn(false); // Update state
    redirect('/'); // Redirect to home page after logout
  };

  // Handle search (placeholder for now)
  const handleSearch = (e) => {
    e.preventDefault();
    console.log('Searching for:', searchTerm);
    // In a real app, you'd navigate to a search results page
    alert(`Searching for "${searchTerm}"... (Functionality not implemented)`); // Use custom modal
  };

  return (
    <header className="bg-gradient-to-r from-blue-600 to-indigo-700 text-white p-4 shadow-lg rounded-b-xl">
      <div className="container mx-auto flex flex-wrap justify-between items-center">
        {/* Blog Title/Logo */}
        <div className="flex items-center mb-4 md:mb-0">
          <h1 className="text-3xl md:text-4xl font-extrabold">
            <Link to="/" className="hover:underline">goBlog</Link>
          </h1>
          <p className="ml-4 text-md md:text-lg opacity-90 hidden md:block">Your daily dose of insights</p>
        </div>

        {/* Navigation and Actions */}
        <nav className="flex flex-grow justify-end items-center space-x-4 md:space-x-6">
          {/* Main Navigation Links */}
          <ul className="flex space-x-4">
            <li>
              <Link to="/" className="hover:underline text-base md:text-lg">Home</Link>
            </li>
            <li>
              <Link to="/posts" className="hover:underline text-base md:text-lg">Posts</Link>
            </li>
            {isLoggedIn && (
              <li>
                <Link to="/create-post" className="hover:underline text-base md:text-lg">Create Post</Link>
              </li>
            )}
          </ul>

          {/* Search Bar */}
          <form onSubmit={handleSearch} className="relative hidden md:block">
            <input
              type="text"
              placeholder="Search..."
              value={searchTerm}
              onChange={(e) => setSearchTerm(e.target.value)}
              className="py-2 px-4 pl-10 rounded-full bg-white text-gray-800 placeholder-gray-500 focus:outline-none focus:ring-2 focus:ring-blue-300 w-48"
            />
            <svg className="absolute left-3 top-1/2 transform -translate-y-1/2 text-gray-500" fill="currentColor" viewBox="0 0 20 20" width="20" height="20">
              <path fillRule="evenodd" d="M8 4a4 4 0 100 8 4 4 0 000-8zM2 8a6 6 0 1110.89 3.476l4.817 4.817a1 1 0 01-1.414 1.414l-4.816-4.816A6 6 0 012 8z" clipRule="evenodd"></path>
            </svg>
          </form>

          {/* Conditional Auth Buttons */}
          <div className="flex space-x-4 ml-4">
            {isLoggedIn ? (
              <>
                <Link
                  to="/user-details"
                  className="bg-white text-blue-700 hover:bg-gray-100 font-bold py-2 px-4 rounded-full text-base shadow-md transition-colors duration-300"
                >
                  User Details
                </Link>
                <button
                  onClick={handleLogout}
                  className="bg-red-500 hover:bg-red-600 text-white font-bold py-2 px-4 rounded-full text-base shadow-md transition-colors duration-300"
                >
                  Logout
                </button>
              </>
            ) : (
              <>
                <Link
                  to="/login"
                  className="bg-white text-blue-700 hover:bg-gray-100 font-bold py-2 px-4 rounded-full text-base shadow-md transition-colors duration-300"
                >
                  Login
                </Link>
                <Link
                  to="/signup"
                  className="bg-green-500 hover:bg-green-600 text-white font-bold py-2 px-4 rounded-full text-base shadow-md transition-colors duration-300"
                >
                  Sign Up
                </Link>
              </>
            )}
          </div>
        </nav>
      </div>
    </header>
  );
};

export default Navbar;
