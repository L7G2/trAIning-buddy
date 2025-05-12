import { useState } from "react";
import reactLogo from "./assets/react.svg";
import viteLogo from "/vite.svg";
import NotLoggedNav from "./components/userNotLoggedIn/navigation/NotLoggedNav";
import AboutPage from "./components/userNotLoggedIn/aboutPage/AboutPage";
import CoachesPage from "./components/userNotLoggedIn/coachesPage/CoachesPage";
import LoginPage from "./components/userNotLoggedIn/loginPage/LoginPage";
import MainPage from "./components/userNotLoggedIn/mainPage/MainPage";
import CalorieCalculator from "./components/userNotLoggedIn/calorieCalculator/calorieCalculator";
import RegisterPage from "./components/userNotLoggedIn/registerPage/RegisterPage";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import "./App.css";

const router = createBrowserRouter([
  {
    path: "/",
    element: <NotLoggedNav />,
    id: "main",
    children: [
      { index: true, element: <MainPage /> },
      { path: "about", element: <AboutPage /> },
      { path: "coaches", element: <CoachesPage /> },
      { path: "caloriecalculator", element: <CalorieCalculator /> },
      { path: "login", element: <LoginPage /> },
      { path: "register", element: <RegisterPage /> },
    ],
  },
]);

function App() {
  return <RouterProvider router={router} />;
}
export default App;
