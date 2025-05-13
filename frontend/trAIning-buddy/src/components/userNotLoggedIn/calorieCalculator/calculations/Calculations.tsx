import "./personData/PersonData";
import PersonData from "./personData/PersonData";
import "./plots/Plots";
import Plots from "./plots/Plots";

function Calculations() {
  return (
    <div style={{ width: "50%" }}>
      <PersonData />
      <Plots />
    </div>
  );
}

export default Calculations;
