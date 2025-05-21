import { useState } from "react";
import Products from "./products/Products";
import Calculations from "./calculations/Calculations";
import "./CalorieCalculator.css";

interface Product {
  id: number;
  name: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}

function CalorieCalculator() {
  const [selectedProducts, setSelectedProducts] = useState<
    Record<number, number>
  >({}); // <-- tutaj obiekt!
  const [productList, setProductList] = useState<Product[]>([]);

  return (
    <div className="calorie-frame">
      <Products
        onSelectionChange={setSelectedProducts}
        onProductList={setProductList}
      />
      <Calculations
        selectedProducts={selectedProducts}
        productList={productList}
      />
    </div>
  );
}

export default CalorieCalculator;
