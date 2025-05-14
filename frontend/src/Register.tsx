import React, { useState } from 'react';
import { registerUser } from './api';

const Register: React.FC = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [role, setRole] = useState('uczen'); // przykładowe role: trener, uczen
    const [error, setError] = useState('');
    const [success, setSuccess] = useState(false);

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await registerUser({ username, password, role });
            setSuccess(true);
            setError('');
        } catch (err) {
            setError('Rejestracja nie powiodła się.');
            setSuccess(false);
        }
    };

    return (
        <div>
            <h2>Rejestracja</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
            {success && <p style={{ color: 'green' }}>Rejestracja przebiegła pomyślnie!</p>}
            <form onSubmit={handleSubmit}>
                <div>
                    <label>Nazwa użytkownika:</label>
                    <input
                        type="text"
                        value={username}
                        onChange={(e) => setUsername(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Hasło:</label>
                    <input
                        type="password"
                        value={password}
                        onChange={(e) => setPassword(e.target.value)}
                        required
                    />
                </div>
                <div>
                    <label>Rola:</label>
                    <select value={role} onChange={(e) => setRole(e.target.value)}>
                        <option value="uczen">Uczeń</option>
                        <option value="trener">Trener</option>
                    </select>
                </div>
                <button type="submit">Zarejestruj</button>
            </form>
        </div>
    );
};

export default Register;