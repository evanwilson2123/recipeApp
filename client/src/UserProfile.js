import React, { useEffect, useState } from "react";
import axiosInstance from "./axiosInstance";
import { jwtDecode } from "jwt-decode";

function UserProfile() {
  const [user, setUser] = useState(null);
  const [recipes, setRecipes] = useState([]);
  const [expandedRecipeId, setExpandedRecipeId] = useState(null);

  useEffect(() => {
    const fetchUser = async () => {
      try {
        const token = localStorage.getItem("token");
        const decoded = jwtDecode(token);

        const userResponse = await axiosInstance.get(`/user/profile`);
        setUser(userResponse.data);

        const recipesResponse = await axiosInstance.get(`/user/recipes`);
        setRecipes(recipesResponse.data);
      } catch (error) {
        console.error(
          "There was an error fetching the user profile or recipes!",
          error
        );
      }
    };

    fetchUser();
  }, []);

  const toggleRecipe = (id) => {
    setExpandedRecipeId(expandedRecipeId === id ? null : id);
  };

  if (!user) {
    return <div>Loading...</div>;
  }

  return (
    <div className="min-h-screen flex flex-col items-center bg-gray-900 text-white">
      <div className="bg-gray-800 p-8 rounded-lg shadow-md w-full max-w-2xl mt-8">
        <h2 className="text-3xl font-bold mb-4 text-center">{user.username}</h2>
        <p className="mb-4 text-center">{user.bio}</p>
        <h3 className="text-2xl font-bold mb-4">Recipes</h3>
        <ul>
          {recipes.map((recipe) => (
            <li key={recipe.id} className="mb-4">
              <div
                className="bg-gray-700 p-4 rounded-lg shadow-md cursor-pointer"
                onClick={() => toggleRecipe(recipe.id)}
              >
                <h4 className="text-xl font-bold">{recipe.title}</h4>
                {expandedRecipeId === recipe.id && (
                  <div>
                    <p className="mt-2">
                      <strong>Ingredients:</strong>{" "}
                      {recipe.ingredients.join(", ")}
                    </p>
                    <p className="mt-2">
                      <strong>Instructions:</strong> {recipe.instructions}
                    </p>
                  </div>
                )}
              </div>
            </li>
          ))}
        </ul>
      </div>
    </div>
  );
}

export default UserProfile;
