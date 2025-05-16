import "./Plots.css";

function Plots() {
  return (
    <div className="plots-frame">
      <PlotsPanel />
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
