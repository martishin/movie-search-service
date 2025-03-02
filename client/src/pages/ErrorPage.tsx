import { JSX } from "react";
import { isRouteErrorResponse, useRouteError, useNavigate } from "react-router";
import Header from "../components/layout/Header";
import Navigation from "../components/layout/Navigation";

export default function ErrorPage(): JSX.Element {
  const error = useRouteError();
  const navigate = useNavigate();

  return (
    <div className="flex h-screen justify-center bg-[#F4F4F4]">
      <div className="flex h-full w-full max-w-screen-lg flex-col bg-white shadow-md">
        <Header />
        <div className="mx-auto w-full border-t border-gray-300" />
        <div className="flex flex-1 overflow-hidden">
          <div className="hidden h-full border-r border-gray-300 md:block md:w-48">
            <Navigation />
          </div>

          <div className="flex flex-grow items-center justify-center p-4">
            <div className="text-center">
              <h1 className="text-3xl font-bold tracking-tight">Oops!</h1>
              <p className="mt-2">Sorry, an unexpected error has occurred.</p>
              <p className="mt-2">
                {isRouteErrorResponse(error) ? (
                  <em>{error.statusText}</em>
                ) : error instanceof Error ? (
                  <em>{error.message}</em>
                ) : null}
              </p>
              <div className="mt-6">
                <button
                  onClick={() => navigate("/", { replace: true })}
                  className="rounded-lg bg-blue-600 px-5 py-2.5 text-sm font-medium text-white hover:bg-blue-800"
                >
                  Go Home
                </button>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
