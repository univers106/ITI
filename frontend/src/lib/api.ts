/**
 * API service for interacting with the backend.
 */

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || "";

export interface LoginCredentials {
  login: string;
  password: string;
}

export interface ApiError {
  message: string;
  status?: number;
}

/**
 * Perform a login request.
 * @param credentials - login and password
 * @returns Promise resolving to response data
 * @throws {ApiError} on failure
 */
export async function login(
  credentials: LoginCredentials,
): Promise<{ message: string }> {
  const formData = new FormData();
  formData.append("login", credentials.login);
  formData.append("password", credentials.password);

  const response = await fetch(`${API_BASE_URL}/api/public/login`, {
    method: "POST",
    body: formData,
    credentials: "include", // include cookies for session
  });

  if (!response.ok) {
    let errorMessage = "Login failed";
    try {
      const errorData = await response.json();
      errorMessage = errorData.message || errorMessage;
    } catch {
      // ignore
    }
    throw { message: errorMessage, status: response.status };
  }

  return response.json();
}

/**
 * Perform a logout request.
 * @returns Promise
 * @throws {ApiError} on failure
 */
export async function logout(): Promise<void> {
  const response = await fetch(`${API_BASE_URL}/api/private/logout`, {
    method: "GET",
    credentials: "include",
  });

  if (!response.ok) {
    let errorMessage = "Logout failed";
    try {
      const errorData = await response.json();
      errorMessage = errorData.message || errorMessage;
    } catch {
      // ignore
    }
    throw { message: errorMessage, status: response.status };
  }
}

/**
 * Check if user is authenticated by calling a protected endpoint.
 * @returns Promise<boolean>
 */
export async function checkAuth(): Promise<boolean> {
  try {
    const response = await fetch(`${API_BASE_URL}/api/private/hello`, {
      method: "GET",
      credentials: "include",
    });
    return response.ok;
  } catch {
    return false;
  }
}

/**
 * Get user greeting from the protected hello endpoint.
 * Returns the raw greeting string.
 * @returns Promise<string>
 * @throws {ApiError} on failure
 */
export async function getUserGreeting(): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/api/private/hello`, {
    method: "GET",
    credentials: "include",
  });

  if (!response.ok) {
    let errorMessage = "Failed to get user info";
    try {
      const errorData = await response.json();
      errorMessage = errorData.message || errorMessage;
    } catch {
      // ignore
    }
    throw { message: errorMessage, status: response.status };
  }

  const greeting = await response.text();
  // Remove JSON quotes if present
  return greeting.replace(/^"|"$/g, '');
}
