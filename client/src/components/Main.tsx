import { createBrowserRouter, RouterProvider } from "react-router";
import App from "./App";
import ErrorPage from "../pages/ErrorPage";
import HomePage from "../pages/HomePage";
import { AlertProvider } from "../context/AlertContext";
import { AuthProvider } from "../context/AuthContext";
import LoginPage from "../pages/LoginPage";
import SignUpPage from "../pages/SignUpPage";
import MoviesPage from "../pages/MoviesPage";
import MoviePage from "../pages/MoviePage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <HomePage /> },
      { path: "/login", element: <LoginPage /> },
      { path: "/sign-up", element: <SignUpPage /> },
      { path: "/movies", element: <MoviesPage /> },
      { path: "/movies/:id", element: <MoviePage /> },
    ],
  },
]);

export default function Main() {
  return (
    <AuthProvider>
      <AlertProvider>
        <div className="min-h-screen bg-[#F4F4F4]">
          <RouterProvider router={router} />
        </div>
      </AlertProvider>
    </AuthProvider>
  );
}
