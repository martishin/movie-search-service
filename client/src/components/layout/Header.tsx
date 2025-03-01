import { JSX } from "react";
import { Link, useNavigate } from "react-router";
import { useAuth } from "../../context/AuthContext";

export default function Header(): JSX.Element {
  const { userDetails, logout } = useAuth();
  const navigate = useNavigate();

  const handleLogout = async () => {
    await logout();
    navigate("/");
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
          <button
            onClick={handleLogout}
            className="min-w-[90px] rounded-md border border-gray-300 bg-white px-5 py-1.5 text-sm leading-6 font-semibold text-gray-900 hover:bg-gray-200"
          >
            Log Out
          </button>
        ) : (
          <Link to="/login">
            <button className="min-w-[90px] rounded-md bg-blue-600 px-5 py-1.5 text-sm leading-6 font-semibold text-white hover:bg-blue-800">
              Log In
            </button>
          </Link>
        )}
      </div>
    </div>
  );
}
