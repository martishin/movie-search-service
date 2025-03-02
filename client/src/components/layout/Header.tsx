import { JSX, useState, useEffect } from "react";
import { Link, useNavigate } from "react-router";
import { useAuth } from "../../context/AuthContext";
import { Bars3Icon, XMarkIcon } from "@heroicons/react/24/outline";
import Navigation from "./Navigation";

export default function Header(): JSX.Element {
  const { userDetails, logout } = useAuth();
  const navigate = useNavigate();
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  // Prevent scrolling when mobile menu is open
  useEffect(() => {
    document.body.style.overflow = isMenuOpen ? "hidden" : "auto";
  }, [isMenuOpen]);

  const handleLogout = async () => {
    await logout();
    navigate("/");
  };

  return (
    <header className="bg-white">
      <div className="mx-auto flex items-center justify-between p-4">
        <div className="flex items-center space-x-3">
          <button
            className="p-2 md:hidden"
            onClick={() => setIsMenuOpen(true)}
            aria-label="Open navigation"
          >
            <Bars3Icon className="h-6 w-6 text-gray-900" />
          </button>

          <Link to="/" className="text-2xl font-bold tracking-normal">
            MovieSearch
          </Link>
        </div>

        <div className="flex items-center space-x-4">
          {userDetails ? (
            <>
              <span className="hidden font-medium text-gray-900 md:inline">
                {`${userDetails.firstName} ${userDetails.lastName}`}
              </span>
              <button
                onClick={handleLogout}
                className="rounded-md border border-gray-300 bg-white px-4 py-2 text-sm font-semibold text-gray-900 hover:bg-gray-200"
              >
                Log Out
              </button>
            </>
          ) : (
            <Link to="/login">
              <button className="rounded-md bg-blue-600 px-4 py-2 text-sm font-semibold text-white hover:bg-blue-800">
                Log In
              </button>
            </Link>
          )}
        </div>
      </div>

      <div
        className={`fixed inset-y-0 left-0 z-50 w-64 transform bg-white shadow-lg transition-transform duration-300 ease-in-out md:hidden ${
          isMenuOpen ? "translate-x-0" : "-translate-x-full"
        }`}
      >
        <div className="flex justify-end p-4">
          <button onClick={() => setIsMenuOpen(false)} aria-label="Close menu">
            <XMarkIcon className="h-6 w-6 text-gray-900" />
          </button>
        </div>

        <Navigation isMobile closeMenu={() => setIsMenuOpen(false)} />
      </div>
    </header>
  );
}
