import React from "react";
import { BrowserRouter as Router, Route, Routes } from "react-router-dom";
import Login from "./Login";
import Register from "./Register";
import RecipeForm from "./RecipeForm";
import UserProfile from "./UserProfile";
import Navbar from "./components/Navbar"; // Correct import path

function App() {
  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/register" element={<Register />} />
        <Route path="/recipe" element={<RecipeForm />} />
        <Route path="/profile" element={<UserProfile />} />{" "}
        {/* Changed path to profile */}
      </Routes>
    </Router>
  );
}

export default App;
