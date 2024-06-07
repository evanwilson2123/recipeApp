import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import axiosInstance from "./axiosInstance"; // Import the axios instance

function Login() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const navigate = useNavigate(); // Initialize useNavigate

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axiosInstance.post("/login", {
        email,
        password,
      });
      console.log(response.data); // Handle the JWT token
      localStorage.setItem("token", response.data.token); // Store the token in local storage
      navigate("/recipe"); // Redirect to the recipe form page after successful login
    } catch (error) {
      console.error("There was an error logging in!", error);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900">
      <div className="bg-gray-800 p-8 rounded-lg shadow-md w-full max-w-sm border-t-4 border-red-500">
        <h2 className="text-3xl font-bold mb-6 text-center text-red-500">
          Login
        </h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-300 mb-2">Email:</label>
            <input
              type="email"
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="w-full px-3 py-2 border border-gray-700 rounded-md bg-gray-700 text-white focus:outline-none focus:ring focus:ring-red-500"
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-300 mb-2">Password:</label>
            <input
              type="password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="w-full px-3 py-2 border border-gray-700 rounded-md bg-gray-700 text-white focus:outline-none focus:ring focus:ring-red-500"
            />
          </div>
          <button
            type="submit"
            className="w-full bg-red-500 text-white py-2 rounded-md hover:bg-red-600 transition duration-300"
          >
            Login
          </button>
        </form>
      </div>
    </div>
  );
}

export default Login;
