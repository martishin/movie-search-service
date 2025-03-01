import { JSX } from "react";
import { Outlet } from "react-router";
import { useAuth } from "../context/AuthContext";
import Navigation from "./layout/Navigation";
import Header from "./layout/Header";

export default function App(): JSX.Element | null {
  const { userDetails } = useAuth();

  if (userDetails === undefined) {
    return null;
  }

  return (
    <div className="container mx-auto mt-8 max-w-screen-lg">
      <Header />
      <div className="mt-4 flex">
        <div className="w-48">
          <Navigation />
        </div>
        <div className="mr-4 ml-4 w-min flex-grow">
          <Outlet />
        </div>
      </div>
    </div>
  );
}
