import { NavLink } from "react-router";
import { JSX } from "react";
import { useAuth } from "../../context/AuthContext";
import {
  FilmIcon,
  HomeIcon,
  PencilSquareIcon,
  PlusCircleIcon,
  TicketIcon,
} from "@heroicons/react/24/outline";

interface NavigationProps {
  isMobile?: boolean;
  closeMenu?: () => void;
}

export default function Navigation({ isMobile, closeMenu }: NavigationProps): JSX.Element {
  const { userDetails } = useAuth();

  const commonLinks = [
    { title: "Home", path: "/", icon: <HomeIcon className="h-5 w-5" /> },
    { title: "Movies", path: "/movies", icon: <FilmIcon className="h-5 w-5" /> },
    { title: "Genres", path: "/genres", icon: <TicketIcon className="h-5 w-5" /> },
  ];

  const loggedInLinks = [
    {
      title: "Manage Catalogue",
      path: "/manage-catalogue",
      icon: <PencilSquareIcon className="h-5 w-5" />,
    },
    { title: "Add a Movie", path: "/admin/movie/0", icon: <PlusCircleIcon className="h-5 w-5" /> },
  ];

  const links = userDetails ? [...commonLinks, ...loggedInLinks] : commonLinks;

  return (
    <nav
      className={`bg-white p-4 ${
        isMobile ? "w-full" : "min-h-screen w-48 border-r border-gray-200"
      }`}
    >
      <ul className="space-y-1">
        {links.map(({ title, path, icon }) => (
          <li key={title}>
            <NavLink
              to={path}
              onClick={closeMenu}
              className={({ isActive }) =>
                `flex items-center gap-3 px-2 py-2 font-medium transition ${
                  isActive
                    ? "border-b-2 border-blue-700 text-blue-700"
                    : "border-b-2 border-transparent text-gray-600 hover:bg-gray-100 hover:text-gray-900"
                }`
              }
            >
              {icon}
              {title}
            </NavLink>
          </li>
        ))}
      </ul>
    </nav>
  );
}
