import React from "react";
import { Link } from "react-router-dom";

function Navbar() {
  const token = localStorage.getItem("token");

  return (
    <nav className="bg-gray-800 p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-white text-lg font-bold">
          Home
        </Link>
        {token && (
          <Link to="/profile" className="text-white text-lg font-bold">
            Profile
          </Link>
        )}
      </div>
    </nav>
  );
}

export default Navbar;
