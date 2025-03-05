import { createContext, ReactNode, useContext, useEffect, useRef, useState } from "react";

import { API_URL } from "../api";

interface UserDetails {
  id: number;
  firstName: string;
  lastName: string;
  email: string;
  pictureUrl?: string;
}

interface AuthContextType {
  userDetails: UserDetails | null | undefined;
  setUserDetails: (user: UserDetails | null) => void;
  login: () => Promise<void>;
  logout: () => Promise<void>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [userDetails, setUserDetails] = useState<UserDetails | null | undefined>(undefined);
  const hasFetched = useRef(false);

  const fetchUserDetails = async () => {
    try {
      const res = await fetch(`${API_URL}/api/users/me`, { credentials: "include" });

      if (!res.ok) {
        throw new Error("Not authenticated");
      }

      const userData: UserDetails = await res.json();
      setUserDetails(userData);
    } catch (err) {
      setUserDetails(null);
    }
  };

  useEffect(() => {
    if (!hasFetched.current) {
      hasFetched.current = true;
      fetchUserDetails();
    }
  }, []);

  const login = async () => {
    await fetchUserDetails();
  };

  const logout = async () => {
    try {
      const response = await fetch(`${API_URL}/auth/logout`, {
        method: "POST",
        credentials: "include",
      });

      if (response.ok) {
        setUserDetails(null);
      } else {
        console.error("Failed to logout:", response.status);
      }
    } catch (error) {
      console.error("Logout error:", error);
    }
  };

  return (
    <AuthContext.Provider value={{ userDetails, setUserDetails, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (!context) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
