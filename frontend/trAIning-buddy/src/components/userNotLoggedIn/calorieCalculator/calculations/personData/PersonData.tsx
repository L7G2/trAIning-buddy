import "./PersonData.css";

function PersonData() {
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
