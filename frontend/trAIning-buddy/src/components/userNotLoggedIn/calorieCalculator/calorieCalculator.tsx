import "./CalorieCalculator.css";
import "./products/Products";
import Products from "./products/Products";
import "./calculations/Calculations";
import Calculations from "./calculations/Calculations";

function CalorieCalculator() {
  return (
    <div className="calorie-frame">
      <Products />
      <Calculations />
    </div>
  );
}

export default CalorieCalculator;
