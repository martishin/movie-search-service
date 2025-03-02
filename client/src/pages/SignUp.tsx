import React, { useState } from "react";
import { Link, useNavigate } from "react-router";
import LoginInput from "../form/LoginInput";
import { useAuth } from "../context/AuthContext";
import { useAlert } from "../context/AlertContext";
import GoogleAuthButton from "../components/GoogleAuthButton";

export default function SignUp() {
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password1, setPassword1] = useState("");
  const [password2, setPassword2] = useState("");
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [loading, setLoading] = useState(false);

  const { login } = useAuth();
  const { showAlert } = useAlert();
  const navigate = useNavigate();

  const validateForm = (): boolean => {
    const newErrors: Record<string, string> = {};

    if (!firstName.trim()) newErrors.firstName = "First name is required.";
    if (!lastName.trim()) newErrors.lastName = "Last name is required.";
    if (!email.trim()) newErrors.email = "Email is required.";
    else if (!/\S+@\S+\.\S+/.test(email)) newErrors.email = "Enter a valid email address.";
    if (!password1) newErrors.password1 = "Password is required.";
    else if (password1.length < 6) newErrors.password1 = "Password must be at least 6 characters.";
    if (!password2) newErrors.password2 = "Confirm password.";
    else if (password1 !== password2) newErrors.password2 = "Passwords do not match.";

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();

    if (!validateForm()) {
      return;
    }

    setLoading(true);

    try {
      const res = await fetch("/api/signup", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        credentials: "include",
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          email: email,
          password: password1,
        }),
      });

      const data = await res.json();

      if (!res.ok) {
        throw new Error(data.error || "Sign up failed. Please try again.");
      }

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
      <form className="mt-3 mr-auto ml-auto w-3/5 max-w-xs" onSubmit={handleSubmit}>
        <LoginInput
          title="First Name"
          type="text"
          name="first-name"
          autoComplete="given-name"
          onChange={(event) => setFirstName(event.target.value)}
          hasError={!!errors.firstName}
          errorMsg={errors.firstName}
        />

        <LoginInput
          title="Last Name"
          type="text"
          name="last-name"
          autoComplete="family-name"
          onChange={(event) => setLastName(event.target.value)}
          hasError={!!errors.lastName}
          errorMsg={errors.lastName}
        />

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
          autoComplete="new-password"
          onChange={(event) => setPassword1(event.target.value)}
          hasError={!!errors.password1}
          errorMsg={errors.password1}
        />

        <LoginInput
          title="Confirm Password"
          type="password"
          name="password-confirm"
          autoComplete="new-password"
          onChange={(event) => setPassword2(event.target.value)}
          hasError={!!errors.password2}
          errorMsg={errors.password2}
        />

        <div className="mt-6">
          <input
            type="submit"
            disabled={loading}
            className={`flex w-full justify-center rounded-md px-3 py-1.5 text-sm leading-6 font-semibold text-white shadow-sm ${
              loading ? "bg-gray-400" : "bg-blue-600 hover:bg-blue-800"
            } focus-visible:outline-2 focus-visible:outline-offset-2`}
            value={loading ? "Signing up..." : "Sign Up"}
          />
        </div>
      </form>

      <div className="mt-4">
        <GoogleAuthButton href="/auth?provider=google" text="Sign up with Google" />
      </div>

      <p className="mt-4 text-center text-sm text-gray-500">
        Already a member?{" "}
        <Link to="/login" className="leading-6 font-semibold text-blue-700 hover:underline">
          Log in
        </Link>
      </p>
    </div>
  );
}
