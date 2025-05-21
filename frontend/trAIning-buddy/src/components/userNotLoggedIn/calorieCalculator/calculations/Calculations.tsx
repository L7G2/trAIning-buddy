import { useState, useEffect } from "react";
import PersonData from "./personData/PersonData";
import Plots from "./plots/Plots";

interface Product {
  id: number;
  name: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}

interface CalcResultType {
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}

interface PersonDataType {
  dailyCalories: number;
  dailyProtein: number;
  dailyFat: number;
  dailyCarbs: number;
  age: number;
  weight: number;
  height: number;
  activityLevel: string;
}

interface PersonBasicData {
  gender: string;
  weight: number;
  lifestyle: string;
  bodyType: string;
}

interface CalculationsProps {
  selectedProducts: Record<number, number>; // { productId: grams }
  productList: Product[];
}

function Calculations({ selectedProducts, productList }: CalculationsProps) {
  const [personData, setPersonData] = useState<PersonDataType | null>(null);

  const handlePersonDataChange = (basicData: PersonBasicData) => {
    // Przykładowe przeliczenia – tutaj możesz dodać dowolną logikę:
    const personData: PersonDataType = {
      dailyCalories: 2000, // możesz tu dodać faktyczne wyliczenia
      dailyProtein: 50,
      dailyFat: 70,
      dailyCarbs: 300,
      age: 30, // domyślna wartość (chyba że potem dodasz do formularza)
      height: 175, // jw.
      weight: basicData.weight,
      activityLevel: basicData.lifestyle,
    };

    setPersonData(personData);
  };
  // symulacja pobierania danych osoby
  useEffect(() => {
    // Tutaj możesz pobrać dane, np. z API lub localStorage
    setPersonData({
      dailyCalories: 2000,
      dailyProtein: 50,
      dailyFat: 70,
      dailyCarbs: 300,
      age: 30,
      weight: 70,
      height: 175,
      activityLevel: "moderate",
    });
  }, []);

  // funkcja licząca składniki odżywcze
  const calculateNutrition = (): CalcResultType => {
    const result = { calories: 0, protein: 0, fat: 0, carbs: 0 };
    for (const idStr in selectedProducts) {
      const id = Number(idStr);
      const grams = selectedProducts[id];
      const product = productList.find((p) => p.id === id);
      if (product && grams > 0) {
        const factor = grams / 100;
        result.calories += product.calories * factor;
        result.protein += product.protein * factor;
        result.fat += product.fat * factor;
        result.carbs += product.carbs * factor;
      }
    }
    return result;
  };

  const calcResult = calculateNutrition();

  if (!personData) return <div>Ładowanie danych...</div>;

  return (
    <div style={{ width: "50%" }}>
      <PersonData onDataChange={handlePersonDataChange} />
      <Plots calcResult={calcResult} personData={personData} />
    </div>
  );
}

export default Calculations;
