import { JSX } from "react";
import { isRouteErrorResponse, useRouteError, useNavigate } from "react-router";
import Header from "../components/layout/Header";

export default function ErrorPage(): JSX.Element {
  const error = useRouteError();
  const navigate = useNavigate();

  return (
    <div className="container mx-auto mt-8 h-screen max-w-screen-lg">
      <Header />

      <div className="flex h-4/6 items-center justify-center">
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
              className="mb-2 rounded-lg bg-blue-600 px-5 py-2.5 text-sm font-medium text-white hover:bg-blue-800"
            >
              Go Home
            </button>
          </div>
        </div>
      </div>
    </div>
  );
}
