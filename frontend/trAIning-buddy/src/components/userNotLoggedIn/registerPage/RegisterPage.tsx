import React, { useState } from "react";
import "./RegisterPage.css";
import { Link, useNavigate } from "react-router-dom";

function RegisterPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [role, setRole] = useState("uczen");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleRegister = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password, role }),
      });

      const data = await response.json();

      if (response.ok) {
        setMessage("Rejestracja zakończona sukcesem!");
        // Opcjonalnie automatyczne przejście na login
        setTimeout(() => navigate("/login"), 1500);
      } else {
        setMessage(data.error || "Błąd rejestracji");
      }
    } catch (err: any) {
      setMessage("Błąd sieci: " + err.message);
    }
  };

  return (
    <div className="register-site-frame">
      <div className="register-frame">
        <h1 className="register-heading">Rejestracja</h1>
        <form onSubmit={handleRegister}>
          <input
            className="register-input"
            type="text"
            placeholder="Nazwa użytkownika"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <br />
          <input
            className="register-input"
            type="password"
            placeholder="Hasło"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <br />
          <select
            className="register-input"
            value={role}
            onChange={(e) => setRole(e.target.value)}
          >
            <option value="uczen">Uczeń</option>
            <option value="trener">Trener</option>
          </select>
          <br />
          <button type="submit" className="register-button">
            Zarejestruj
          </button>
        </form>
        {message && <p>{message}</p>}
        <div className="link-to-register">
          Masz już konto? <Link to="/login">Zaloguj się!</Link>
        </div>
      </div>
    </div>
  );
}

export default RegisterPage;
