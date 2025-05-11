import { useEffect, useState } from "react";
import "./LookingForClients.css";
import { Link } from "react-router-dom";
import personalTraining from "/training.png";

function LookingForClients() {
  const [scrollY, setScrollY] = useState(0);

  useEffect(() => {
    const handleScroll = () => {
      setScrollY(window.scrollY);
    };

    window.addEventListener("scroll", handleScroll);

    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  const scrollOffset = Math.max(scrollY - 350, 0);
  const scrollPercent = Math.min(scrollOffset / 600, 1);

  const blackStop = scrollPercent * 100;

  const dynamicBackground = `linear-gradient(to left, yellow 0%, black ${blackStop}%)`;
  return (
    <div className="client-frame" style={{ background: dynamicBackground }}>
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
