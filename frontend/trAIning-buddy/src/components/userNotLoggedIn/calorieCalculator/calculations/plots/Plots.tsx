import "./Plots.css";
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

function Plots({ calcResult, personData, selectedProducts }: PlotsProps) {
  return (
    <div className="plots-panel">
      <Heading />
      <h3>Przetestuj przesyłanie danych:</h3>

      <div>
        <h4>Obliczone wartości odżywcze:</h4>
        <pre>{JSON.stringify(calcResult, null, 2)}</pre>
      </div>

      <div>
        <h4>Dane użytkownika:</h4>
        <pre>
          {personData
            ? JSON.stringify(personData, null, 2)
            : "Brak danych użytkownika"}
        </pre>
      </div>

      <div>
        <h4>Wybrane produkty:</h4>
        <pre>{JSON.stringify(selectedProducts, null, 2)}</pre>
      </div>
    </div>
  );
}

function Heading() {
  return <div className="plots-heading">Wartości odżywcze</div>;
}

export default Plots;
