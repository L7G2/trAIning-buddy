import "./Plots.css";
interface CalcResult {
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}

interface PersonData {
  dailyCalories: number;
  dailyProtein: number;
  dailyFat: number;
  dailyCarbs: number;
}

interface PlotsProps {
  calcResult: CalcResult | null;
  personData: PersonData | null;
}

function Plots({ calcResult, personData }: PlotsProps) {
  if (!calcResult || !personData) return <div>Brak danych</div>;

  const { calories, protein, fat, carbs } = calcResult;
  const { dailyCalories, dailyProtein, dailyFat, dailyCarbs } = personData;

  const percent = {
    calories: (calories / dailyCalories) * 100,
    protein: (protein / dailyProtein) * 100,
    fat: (fat / dailyFat) * 100,
    carbs: (carbs / dailyCarbs) * 100,
  };

  return (
    <div className="plots-panel">
      <Heading />
      {/* tu będą wykresy */}
    </div>
  );
}

function PlotsPanel() {
  return (
    <div className="plots-panel">
      <Heading />
    </div>
  );
}

function Heading() {
  return <div className="plots-heading">Wartości odżywcze</div>;
}

export default Plots;
