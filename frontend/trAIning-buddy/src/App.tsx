import { useUser } from "./UserContext";
import { useMemo } from "react";
import NotLoggedNav from "./components/userNotLoggedIn/navigation/NotLoggedNav";
import CoachNav from "./components/userLoggedInAsCoach/navigation/CoachNav";
import StudentNav from "./components/userLoggedInAsClient/navigation/StudentNav";
import AboutPage from "./components/userNotLoggedIn/aboutPage/AboutPage";
import CoachesPage from "./components/userNotLoggedIn/coachesPage/CoachesPage";
import LoginPage from "./components/userNotLoggedIn/loginPage/LoginPage";
import MainPage from "./components/userNotLoggedIn/mainPage/MainPage";
import CalorieCalculator from "./components/userNotLoggedIn/calorieCalculator/CalorieCalculator";
import RegisterPage from "./components/userNotLoggedIn/registerPage/RegisterPage";
import StudentDashboard from "./components/userLoggedInAsClient/dashboard/StudentDashboard";
import CoachDashboard from "./components/userLoggedInAsCoach/dashboard/CoachDashboard";
import {
  createBrowserRouter,
  RouterProvider,
  BrowserRouter,
  Routes,
  Route,
} from "react-router-dom";
import "./App.css";

function App() {
  const { role } = useUser();

  console.log("Aktualna rola:", role);

  let Navigation;
  if (role === "uczen") Navigation = <StudentNav />;
  else if (role === "trener") Navigation = <CoachNav />;
  else Navigation = <NotLoggedNav />;

  return (
    <BrowserRouter>
      {Navigation}
      <Routes>
        {role === "uczen" && <Route path="/" element={<StudentDashboard />} />}
        {role === "trener" && <Route path="/" element={<CoachDashboard />} />}
        {!role && (
          <>
            <Route path="/" element={<MainPage />} />
            <Route path="/about" element={<AboutPage />} />
            <Route path="/coaches" element={<CoachesPage />} />
            <Route path="/caloriecalculator" element={<CalorieCalculator />} />
            <Route path="/login" element={<LoginPage />} />
            <Route path="/register" element={<RegisterPage />} />
          </>
        )}
      </Routes>
    </BrowserRouter>
  );
}

export default App;
