import React, { useState } from "react";
import axiosInstance from "./axiosInstance"; // Import the axios instance

function RecipeForm() {
  const [title, setTitle] = useState("");
  const [ingredients, setIngredients] = useState("");
  const [instructions, setInstructions] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axiosInstance.post("/recipes", {
        title,
        ingredients: ingredients.split(","),
        instructions,
      });
      console.log(response.data); // Handle the response
    } catch (error) {
      console.error("There was an error creating the recipe!", error);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-900">
      <div className="bg-gray-800 p-8 rounded-lg shadow-md w-full max-w-sm border-t-4 border-orange-500">
        <h2 className="text-3xl font-bold mb-6 text-center text-orange-500">
          Create Recipe
        </h2>
        <form onSubmit={handleSubmit}>
          <div className="mb-4">
            <label className="block text-gray-300 mb-2">Title:</label>
            <input
              type="text"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              className="w-full px-3 py-2 border border-gray-700 rounded-md bg-gray-700 text-white focus:outline-none focus:ring focus:ring-orange-500"
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-300 mb-2">
              Ingredients (comma separated):
            </label>
            <input
              type="text"
              value={ingredients}
              onChange={(e) => setIngredients(e.target.value)}
              className="w-full px-3 py-2 border border-gray-700 rounded-md bg-gray-700 text-white focus:outline-none focus:ring focus:ring-orange-500"
            />
          </div>
          <div className="mb-4">
            <label className="block text-gray-300 mb-2">Instructions:</label>
            <textarea
              value={instructions}
              onChange={(e) => setInstructions(e.target.value)}
              className="w-full px-3 py-2 border border-gray-700 rounded-md bg-gray-700 text-white focus:outline-none focus:ring focus:ring-orange-500"
            />
          </div>
          <button
            type="submit"
            className="w-full bg-orange-500 text-white py-2 rounded-md hover:bg-orange-600 transition duration-300"
          >
            Create Recipe
          </button>
        </form>
      </div>
    </div>
  );
}

export default RecipeForm;
