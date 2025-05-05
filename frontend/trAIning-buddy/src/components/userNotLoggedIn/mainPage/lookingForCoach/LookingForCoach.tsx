import "./LookingForCoach.css";
import { Link } from "react-router-dom";
import personalTrainer from "/personal trainer.png";

function LookingForCoach() {
  return (
    <div className="coach-frame">
      <div className="coach-text-frame">
        <div style={{ fontSize: "6em" }}>Szukasz trenera personalnego?</div>
        <div style={{ fontSize: "3em", marginTop: "10%" }}>
          Zarejestruj się i znajdź idealnego trenera jeszcze dziś!
        </div>
        <Link to="/login" className="coach-button-link">
          zarejestruj się
        </Link>
      </div>
      <img src={personalTrainer} alt="Personal Trainer" />
    </div>
  );
}

export default LookingForCoach;
