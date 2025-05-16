import "./CalorieComponent.css";
import { Link } from "react-router-dom";

function CalorieComponent() {
  return (
    <div className="calorie-component">
      <div className="calorie-text">
        A TY? Ile kalorii tak naprawdę potrzebujesz?
      </div>
      <Link to="/caloriecalculator" className="calorie-button-link">
        Kalkulator kalorii
      </Link>
    </div>
  );
}

export default CalorieComponent;
