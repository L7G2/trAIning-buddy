import { useState } from "react";
import PersonData from "./personData/PersonData";
import Plots from "./plots/Plots";

interface SelectedProduct {
  id: number;
  name: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
  amount: number;
}

interface CalcResultType {
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}

// Dane podstawowe z PersonData
interface PersonBasicData {
  gender: string;
  weight: number;
  lifestyle: string;
  bodyType: string;
}

interface CalculationsProps {
  selectedProducts: SelectedProduct[]; // już tylko wybrane z ilością
}

function Calculations({ selectedProducts }: CalculationsProps) {
  const [personBasicData, setPersonBasicData] =
    useState<PersonBasicData | null>(null);

  const handlePersonDataChange = (data: PersonBasicData) => {
    setPersonBasicData(data);
  };

  const calculateNutrition = (): CalcResultType => {
    return selectedProducts.reduce(
      (acc, product) => {
        const factor = product.amount / 100;
        acc.calories += product.calories * factor;
        acc.protein += product.protein * factor;
        acc.fat += product.fat * factor;
        acc.carbs += product.carbs * factor;
        return acc;
      },
      { calories: 0, protein: 0, fat: 0, carbs: 0 }
    );
  };

  const calcResult = calculateNutrition();

  return (
    <div style={{ width: "50%" }}>
      <PersonData onDataChange={handlePersonDataChange} />
      <Plots
        calcResult={calcResult}
        personData={personBasicData}
        selectedProducts={selectedProducts}
      />
    </div>
  );
}

export default Calculations;
