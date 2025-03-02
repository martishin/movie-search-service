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
    <div className="flex h-screen justify-center bg-[#F4F4F4]">
      <div className="flex h-full w-full max-w-screen-lg flex-col bg-white shadow-md">
        <Header />
        <div className="mx-auto w-full border-t border-gray-300" />
        <div className="flex flex-1 overflow-hidden">
          <div className="hidden h-full border-r border-gray-300 md:block md:w-48">
            <Navigation />
          </div>

          <div className="flex-grow overflow-auto p-4">
            <Outlet />
          </div>
        </div>
      </div>
    </div>
  );
}
