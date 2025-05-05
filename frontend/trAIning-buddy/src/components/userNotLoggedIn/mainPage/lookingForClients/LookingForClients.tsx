import "./LookingForClients.css";
import { Link } from "react-router-dom";
import personalTraining from "/training.png";

function LookingForClients() {
  return (
    <div className="client-frame">
      <img src={personalTraining} alt="Personal Trainer" />
      <div className="client-text-frame">
        <div style={{ fontSize: "6em" }}>
          Jesteś trenerem i szukasz klientów?
        </div>
        <div style={{ fontSize: "3em", marginTop: "10%" }}>
          Zarejestruj się i poznaj setki osób gotowych na współpracę z Tobą!
        </div>
        <Link to="/login" className="client-button-link">
          zarejestruj się
        </Link>
      </div>
    </div>
  );
}

export default LookingForClients;
