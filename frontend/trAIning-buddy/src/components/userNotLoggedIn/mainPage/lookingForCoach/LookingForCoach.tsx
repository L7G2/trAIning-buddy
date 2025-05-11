import { useEffect, useState } from "react";
import "./LookingForCoach.css";
import { Link } from "react-router-dom";
import personalTrainer from "/personal trainer.png";

function LookingForCoach() {
  const [scrollY, setScrollY] = useState(0);

  useEffect(() => {
    const handleScroll = () => {
      setScrollY(window.scrollY);
    };

    window.addEventListener("scroll", handleScroll);

    return () => window.removeEventListener("scroll", handleScroll);
  }, []);

  const scrollPercent = Math.min(scrollY / 300, 1);

  const blackStop = scrollPercent * 100;

  const dynamicBackground = `linear-gradient(to right, green 0%, black ${blackStop}%)`;

  return (
    <div className="coach-frame" style={{ background: dynamicBackground }}>
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
