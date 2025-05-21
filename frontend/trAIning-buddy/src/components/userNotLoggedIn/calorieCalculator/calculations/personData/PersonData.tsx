import { useState } from "react";
import "./PersonData.css";

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

interface PersonDataProps {
  onDataChange: (data: PersonBasicData) => void;
}

function PersonData({ onDataChange }: PersonDataProps) {
  const [gender, setGender] = useState("woman");
  const [weight, setWeight] = useState<number | "">("");
  const [lifestyle, setLifestyle] = useState("lazy");
  const [bodyType, setBodyType] = useState("athlete");
  const handleSubmit = () => {
    if (weight === "" || weight <= 0) {
      alert("Podaj poprawną wagę!");
      return;
    }
    onDataChange({
      gender,
      weight: Number(weight),
      lifestyle,
      bodyType,
    });
  };
  return (
    <div className="person-data-frame">
      <PersonDataPanel />
    </div>
  );
}

function PersonDataPanel() {
  return (
    <div className="person-data-panel">
      <Heading />
      <DataInput />
      <button className="person-data-button">Zatwierdź</button>
    </div>
  );
}

function Heading() {
  return <div className="person-data-heading">Informacje o Tobie</div>;
}

function DataInput() {
  return (
    <div className="data-input">
      <div className="data-column">
        <div className="person-data-label">
          Płeć:
          <select className="person-data-input">
            <option value="woman">Kobieta</option>
            <option value="man">Mężczyzna</option>
          </select>
        </div>
        <div className="person-data-label">
          Waga (kg):
          <input className="person-data-input" placeholder="waga" />
        </div>
      </div>
      <div className="data-column">
        <div className="person-data-label">
          Tryb życia:
          <select className="person-data-input">
            <option value="lazy">Siedzący</option>
            <option value="mix">Mieszany</option>
            <option value="active">Aktywny</option>
            <option value="vactive">Bardzo aktywny</option>
          </select>
        </div>

        <div className="person-data-label">
          Sylwetka:
          <select className="person-data-input">
            <option value="athlete">Wysportowana</option>
            <option value="slim">Szczupła</option>
            <option value="overweight">Nadwaga</option>
            <option value="obesity">Otyłość</option>
          </select>
        </div>
      </div>
    </div>
  );
}

export default PersonData;
