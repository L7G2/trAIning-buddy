export interface AuthResponse {
    token: string;
}

export interface RegisterData {
    username: string;
    password: string;
    role: string;
}

export interface LoginData {
    username: string;
    password: string;
}

// Rejestracja
export async function registerUser(data: RegisterData): Promise<AuthResponse> {
    const response = await fetch('http://localhost:8080/register', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });
    if (!response.ok) {
        throw new Error('Błąd podczas rejestracji');
    }
    return response.json();
}

// Logowanie
export async function loginUser(data: LoginData): Promise<AuthResponse> {
    const response = await fetch('http://localhost:8080/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
    });
    if (!response.ok) {
        throw new Error('Błąd podczas logowania');
    }
    return response.json();
}

// Przykład wywołania zasobu chronionego
export async function getProtectedResource(token: string): Promise<any> {
    const response = await fetch('http://localhost:8080/protected', {
        headers: {
            Authorization: `Bearer ${token}`,
        },
    });
    if (!response.ok) {
        throw new Error('Brak dostępu (niepoprawny token)');
    }
    return response.json();
}