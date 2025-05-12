import React, { useState } from "react";
import "./LoginPage.css";
import { Link } from "react-router-dom";
import { useNavigate } from "react-router-dom";

function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [message, setMessage] = useState("");
  const navigate = useNavigate();

  const handleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username, password }),
      });

      const data = await response.json();

      if (response.ok) {
        localStorage.setItem("token", data.token);
        setMessage(`Zalogowano jako ${data.user.username}`);
        navigate("/dashboard");
      } else {
        setMessage(data.error || "Błąd logowania");
      }
    } catch (err: any) {
      setMessage("Błąd sieci: " + err.message);
    }
  };

  return (
    <div className="login-site-frame">
      <div className="login-frame">
        <h1 className="login-heading">Logowanie</h1>
        <form onSubmit={handleLogin}>
          <input
            className="login-input"
            type="text"
            placeholder="Nazwa użytkownika"
            value={username}
            onChange={(e) => setUsername(e.target.value)}
            required
          />
          <br />
          <input
            className="login-input"
            type="password"
            placeholder="Hasło"
            value={password}
            onChange={(e) => setPassword(e.target.value)}
            required
          />
          <br />
          <button type="submit" className="login-button">
            Zaloguj
          </button>
        </form>
        {message && <p>{message}</p>}
        <div className="link-to-register">
          Nie masz jeszcze konta? <Link to="/register">Zarejestruj się!</Link>
        </div>
      </div>
    </div>
  );
}

export default LoginPage;
