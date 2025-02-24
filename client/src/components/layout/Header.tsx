import { JSX } from "react";
import { Link, useNavigate } from "react-router";
import UserDetails from "../../models/UserDetails";

interface HeaderProps {
  userDetails: UserDetails | null;
  setUserDetails: (userDetails: UserDetails | null) => void;
}

export default function Header({ userDetails, setUserDetails }: HeaderProps): JSX.Element {
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      await fetch("/api/logout", {
        credentials: "include",
      });
      setUserDetails(null);
      navigate("/");
    } catch (err) {
      console.error("Error during logout:", err);
      alert("Logout failed. Please try again.");
    }
  };

  return (
    <div className="flex flex-wrap items-center">
      <div className="max-w-full flex-1 flex-grow px-4">
        <Link to="/">
          <h1 className="text-3xl font-bold tracking-tight">MovieSearch</h1>
        </Link>
      </div>
      <div className="max-w-full flex-grow px-4 text-end">
        {userDetails ? (
          <a href="#!" onClick={handleLogout}>
            <span className="mb-2 rounded-lg border border-gray-300 bg-white px-5 py-2.5 text-sm font-medium text-gray-900 hover:bg-gray-200">
              Log Out
            </span>
          </a>
        ) : (
          <Link to="/login">
            <span className="mb-2 rounded-lg bg-blue-600 px-5 py-2.5 text-sm font-medium text-white hover:bg-blue-800">
              Log In
            </span>
          </Link>
        )}
      </div>
    </div>
  );
}
