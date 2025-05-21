import "./Plots.css";
import { Pie } from "react-chartjs-2";
import { Doughnut } from "react-chartjs-2";
import { Chart as ChartJS, ArcElement, Tooltip, Legend } from "chart.js";
ChartJS.register(ArcElement, Tooltip, Legend);

// Rejestracja elementów wykresu
ChartJS.register(ArcElement, Tooltip, Legend);

interface CalcResult {
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
}

// Dane użytkownika (normy)
interface PersonData {
  gender: string;
  weight: number;
  lifestyle: string;
  bodyType: string;
}

// Dane wybranych produktów z ilością
interface SelectedProduct {
  id: number;
  name: string;
  calories: number;
  protein: number;
  fat: number;
  carbs: number;
  amount: number;
}

interface PlotsProps {
  calcResult: CalcResult;
  personData: PersonData | null;
  selectedProducts: SelectedProduct[];
}

function calculateDailyCalories(personData: PersonData): number {
  const { gender, weight, lifestyle, bodyType } = personData;

  // Podstawowa przemiana materii (BMR) — uproszczona
  let bmr = gender === "man" ? 24 * weight : 22 * weight;

  // Współczynniki aktywności
  const lifestyleFactor: Record<string, number> = {
    lazy: 1.2,
    mix: 1.4,
    active: 1.6,
    vactive: 1.8,
  };

  // Korekta na sylwetkę
  const bodyTypeModifier: Record<string, number> = {
    slim: 0.95,
    athlete: 1.0,
    overweight: 1.05,
    obesity: 1.1,
  };

  const lifestyleMult = lifestyleFactor[lifestyle] ?? 1.2;
  const bodyTypeMult = bodyTypeModifier[bodyType] ?? 1.0;

  return Math.round(bmr * lifestyleMult * bodyTypeMult);
}

function Plots({ calcResult, personData, selectedProducts }: PlotsProps) {
  const data = [
    { name: "B", value: calcResult.protein },
    { name: "T", value: calcResult.fat },
    { name: "W", value: calcResult.carbs },
  ];

  const pieData = {
    labels: data.map((d) => d.name),
    datasets: [
      {
        data: data.map((d) => d.value),
        backgroundColor: ["#82ca9d", "#ffc658", "#8884d8"],
      },
    ],
  };

  const pieOptions = {
    plugins: {
      tooltip: {
        callbacks: {
          label: function (context: any) {
            const value = context.parsed;
            const total =
              calcResult.protein + calcResult.fat + calcResult.carbs;
            const percentage = ((value / total) * 100).toFixed(1);
            return `${context.label}: ${percentage}% (${value} g)`;
          },
        },
      },
      legend: {
        display: true,
      },
    },
  };

  const calorieChart = personData
    ? (() => {
        const daily = calculateDailyCalories(personData);
        const consumed = selectedProducts.reduce(
          (sum, p) => sum + p.calories * p.amount,
          0
        );
        const remaining = Math.max(daily - consumed, 0);

        return {
          data: {
            labels: ["Zjedzone", "Pozostało"],
            datasets: [
              {
                data: [consumed, remaining],
                backgroundColor: ["#ff6384", "#cccccc"],
              },
            ],
          },
          options: {
            plugins: {
              tooltip: {
                callbacks: {
                  label: function (context: any) {
                    const value = context.parsed;
                    const percentage = ((value / daily) * 100).toFixed(1);
                    return `${context.label}: ${percentage}% (${value} kcal)`;
                  },
                },
              },
              legend: {
                display: true,
              },
            },
          },
        };
      })()
    : null;

  return (
    <div className="plots-frame">
      <Heading />
      <div className="plots-panel">
        {calcResult &&
          (calcResult.protein || calcResult.fat || calcResult.carbs) > 0 && (
            <div style={{ maxWidth: "50%", margin: "2rem auto" }}>
              <h3>Makroskładniki</h3>
              <Pie data={pieData} options={pieOptions} />
            </div>
          )}

        {personData && calorieChart && (
          <div style={{ maxWidth: "50%", margin: "2rem auto" }}>
            <h3>Pokrycie zapotrzebowania kalorycznego</h3>
            <Doughnut data={calorieChart.data} options={calorieChart.options} />
          </div>
        )}
      </div>
    </div>
  );
}

function Heading() {
  return <div className="plots-heading">Wartości odżywcze</div>;
}

export default Plots;
