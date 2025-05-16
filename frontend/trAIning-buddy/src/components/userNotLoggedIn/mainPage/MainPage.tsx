import LookingForCoach from "./lookingForCoach/LookingForCoach";
import LookingForClients from "./lookingForClients/LookingForClients";
import CalorieComponent from "./calorieComponent/CalorieComponent";
import "./MainPage.css";
function MainPage() {
  return (
    <div className="main-page-container">
      <LookingForCoach />
      <LookingForClients />
      <CalorieComponent />
    </div>
  );
}

export default MainPage;
