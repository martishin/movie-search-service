import React, { useState } from "react";
import { Link, useNavigate } from "react-router";

import { API_URL } from "../api";
import GoogleAuthButton from "../components/GoogleAuthButton";
import { useAlert } from "../context/AlertContext";
import { useAuth } from "../context/AuthContext";
import LoginInput from "../form/LoginInput";

export default function LoginPage() {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [errors, setErrors] = useState<{ email?: string; password?: string }>({});
  const [loading, setLoading] = useState(false);

  const { showAlert } = useAlert();
  const { login } = useAuth();
  const navigate = useNavigate();

  const handleEmailLogin = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    const newErrors: { email?: string; password?: string } = {};
    if (!email) newErrors.email = "Email is required";
    if (!password) newErrors.password = "Password is required";

    setErrors(newErrors);
    if (Object.keys(newErrors).length > 0) return;

    setLoading(true);

    try {
      const res = await fetch(`${API_URL}/auth/login`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({ email, password }),
      });

      const data = await res.json();
      if (!res.ok) throw new Error(data.error || "Invalid credentials");

      await login();
      navigate("/");
    } catch (error: unknown) {
      if (error instanceof Error) {
        showAlert(error.message);
      } else {
        showAlert("An unknown error occurred.");
      }
    } finally {
      setLoading(false);
    }
  };

  return (
    <div>
      <form className="mt-3 mr-auto ml-auto w-3/5 max-w-xs" onSubmit={handleEmailLogin}>
        <LoginInput
          title="Email Address"
          type="email"
          name="email"
          autoComplete="email"
          onChange={(event) => setEmail(event.target.value)}
          hasError={!!errors.email}
          errorMsg={errors.email}
        />

        <LoginInput
          title="Password"
          type="password"
          name="password"
          autoComplete="current-password"
          onChange={(event) => setPassword(event.target.value)}
          hasError={!!errors.password}
          errorMsg={errors.password}
        />

        <div className="mt-6">
          <input
            type="submit"
            disabled={loading}
            className="flex w-full justify-center rounded-md bg-blue-600 px-3 py-1.5 text-sm leading-6 font-semibold text-white hover:bg-blue-800 focus-visible:bg-blue-600 focus-visible:outline-2 focus-visible:outline-offset-2"
            value={loading ? "Logging in..." : "Log In"}
          />
        </div>
      </form>

      <div className="mt-4">
        <GoogleAuthButton
          href={`${API_URL}/auth/start?provider=google`}
          text="Log in with Google"
        />
      </div>

      <p className="mt-4 text-center text-sm text-gray-500">
        Not a member?{" "}
        <Link to="/sign-up" className="leading-6 font-semibold text-blue-700 hover:underline">
          Sign up
        </Link>
      </p>
    </div>
  );
}
