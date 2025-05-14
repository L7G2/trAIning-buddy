import React, { useState } from 'react';
import { loginUser } from './api';

interface LoginProps {
    onLoginSuccess: (token: string) => void;
}

const Login: React.FC<LoginProps> = ({ onLoginSuccess }) => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState('');

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            const { token } = await loginUser({ username, password });
            onLoginSuccess(token);
            setError('');
        } catch (err) {
            setError('Logowanie nie powiodło się.');
        }
    };

    return (
        <div>
            <h2>Logowanie</h2>
            {error && <p style={{ color: 'red' }}>{error}</p>}
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
                <button type="submit">Zaloguj</button>
            </form>
        </div>
    );
};

export default Login;