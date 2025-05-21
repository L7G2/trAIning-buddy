import { useState } from "react";
import Products from "./products/Products";
import Calculations from "./calculations/Calculations";
import "./CalorieCalculator.css";

interface SelectedProduct {
  id: number;
  name: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
  amount: number; // ilość wybrana przez użytkownika
}

function CalorieCalculator() {
  const [selectedProducts, setSelectedProducts] = useState<SelectedProduct[]>(
    []
  );

  return (
    <div className="calorie-frame">
      <Products onSelectionChange={setSelectedProducts} />
      <Calculations selectedProducts={selectedProducts} />
    </div>
  );
}

export default CalorieCalculator;
