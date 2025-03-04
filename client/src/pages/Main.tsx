import { createBrowserRouter, RouterProvider } from "react-router";
import App from "./App";
import ErrorPage from "./ErrorPage";
import WatchOnlinePage from "./WatchOnlinePage";
import { AlertProvider } from "../context/AlertContext";
import { AuthProvider } from "../context/AuthContext";
import LoginPage from "./LoginPage";
import SignUpPage from "./SignUpPage";
import MoviesPage from "./MoviesPage";
import MoviePage from "./MoviePage";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    errorElement: <ErrorPage />,
    children: [
      { index: true, element: <WatchOnlinePage /> },
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
